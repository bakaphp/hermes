package providers

import (
	"encoding/json"
	"log"

	"src/feeds/models"

	"github.com/streadway/amqp"
)

// IncomingData : Message Data
type IncomingData struct {
	UsersID         int    `json:"users_id"`
	MessageID       int    `json:"message_id"`
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

	db, dbError := MysqlConnect()

	if dbError != nil {
		panic(dbError.Error())
	}

	var incomingData IncomingData
	json.Unmarshal([]byte(msg), &incomingData)

	userFollows := models.UsersFollows{EntityID: incomingData.UsersID, EntityNamespace: incomingData.EntityNamespace, IsDeleted: incomingData.IsDeleted}
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
