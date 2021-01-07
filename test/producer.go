package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// Producer function
func Producer(client pulsar.Client) {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "persistent://public/default/query-project-service",
	})

	i := 0
	for {
		person := Person{Name: "Alice " + strconv.Itoa(i), Age: i}

		_, err1 := json.Marshal(person)

		if err1 != nil {
			fmt.Println("error:", err)
		}

		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte("Hola"),
		})

		//time.Sleep(2 * time.Second)

		if err != nil {
			fmt.Println("Failed to publish message", err)
			break
		}
		i++
	}
	defer producer.Close()

}
