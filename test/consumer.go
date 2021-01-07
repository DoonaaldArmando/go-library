package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Consumer function
func Consumer(client pulsar.Client, topic string) {

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: topic,
		Type:             pulsar.Exclusive,
	})

	if err == nil {
		for {
			msg, err := consumer.Receive(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			/*var person Person
			err1 := json.Unmarshal(msg.Payload(), &person)
			if err1 != nil {
				fmt.Println("error:", err)
			}*/

			fmt.Printf("content: '%+v'\n", string(msg.Payload()))

			consumer.Ack(msg)

		}

	}

}
