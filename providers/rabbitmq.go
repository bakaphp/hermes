package providers

import (
	"encoding/json"
	"log"

	"src/feeds/models"

	"github.com/streadway/amqp"
)

// IncomingData : Message Data
type IncomingData struct {
	Action          string `json:"action"`
	EntityID        int    `json:"entity_id"`
	UsersID         int    `json:"users_id"`
	MessageID       int    `json:"message_id"`
	NumMessages     int    `json:"num_messages"`
	IsDeleted       int    `json:"is_deleted"`
	EntityNamespace string `json:"entity_namespace"`
	DeleteMessage   int    `json:"delete_message"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// RunQueue connects and runs RabbitMQ
func RunQueue(channelName string) {

	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		channelName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			processMessage(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// processMessage processes the incoming messages
func processMessage(msg []byte) {

	var incomingData IncomingData
	json.Unmarshal([]byte(msg), &incomingData)

	//Incoming data declares which action must happen
	switch action := incomingData.Action; action {
	case "new_follower":
		distributeXMessages(incomingData)
	default:
		distributeMessagesToFollowers(incomingData)
	}

}

// Distribute all messages of an entity to all its followers
func distributeMessagesToFollowers(incomingData IncomingData) {

	db, dbError := MysqlConnect()

	if dbError != nil {
		panic(dbError.Error())
	}

	userFollows := models.UsersFollows{EntityID: incomingData.EntityID, EntityNamespace: incomingData.EntityNamespace, IsDeleted: incomingData.IsDeleted}
	userMessages := models.UserMessages{}
	userFollowsArray := []models.UsersFollows{}

	//Convert json to struct data

	//Find all the users followers by users_id and entity_namespace
	db.Debug().Where(&userFollows).Find(&userFollowsArray)

	// Traverse array of user follows
	for _, userFollow := range userFollowsArray {
		//Batch create users messages or group messages

		if incomingData.DeleteMessage == 0 {
			userMessages.MessageID = incomingData.MessageID
			userMessages.UsersID = userFollow.UsersID
			userMessages.IsDeleted = 0

			db.Debug().Create(&userMessages)
		} else {
			db.Debug().Model(&userMessages).Where("users_id = ? and message_id = ?", userFollow.EntityID, incomingData.MessageID).Update("is_deleted", 1)
		}
	}

	db.Close()
	log.Printf("Process Completed")

}

// Distributes X messages of an entity to a new follower
func distributeXMessages(incomingData IncomingData) {

	db, dbError := MysqlConnect()

	if dbError != nil {
		panic(dbError.Error())
	}

	userFollows := models.UsersFollows{EntityID: incomingData.EntityID, UsersID: incomingData.UsersID, EntityNamespace: incomingData.EntityNamespace, IsDeleted: incomingData.IsDeleted}
	userMessages := models.UserMessages{}
	recentFollower := models.UsersFollows{}

	//Convert json to struct data

	//Find all the users followers by users_id and entity_namespace
	db.Debug().Where(&userFollows).Find(&recentFollower)

	// Find the last X messages from entity
	messageQuery := models.Messages{}
	messages := []models.Messages{}

	db.Debug().Where(&messageQuery).Order("id desc").Limit(incomingData.NumMessages).Find(&messages)

	// For each message assign it to the recent follower
	for _, message := range messages {
		//Batch create users messages or group messages
		userMessages.MessageID = int(message.ID)
		userMessages.UsersID = recentFollower.UsersID
		userMessages.IsDeleted = 0
		db.Debug().Create(&userMessages)
	}

	db.Close()
	log.Printf("Process Completed")
}
