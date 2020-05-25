package main

import (
	"fmt"
	"os"
	"src/feeds/providers"
)

func main() {

	if len(os.Args) >= 1 {
		var queueName string

		queueName = os.Args[1]

		fmt.Println("Working with queue: ", queueName)

		providers.RunQueue(queueName)

	} else {
		fmt.Println("Name of the queue not passed as argument")
	}
}
