package main

import (
	"fmt"
	"os"
	"src/feeds/services"
)

func main() {

	if len(os.Args) >= 1 {
		var queueName string

		queueName = os.Args[1]

		fmt.Println("Working with queue: ", queueName)

		fmt.Println("Entity Namespace: ", os.Getenv("ENTITY_NAMESPACE"))

		services.RunQueue(queueName)

		// storage.MigrateDB()
		// storage.SeedDB()

	} else {
		fmt.Println("Name of the queue not passed as argument")
	}
}
