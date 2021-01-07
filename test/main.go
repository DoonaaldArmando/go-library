package main

import (
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://ec2-18-220-29-225.us-east-2.compute.amazonaws.com:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	go Producer(client)
	go Consumer(client, "persistent://public/default/query-project-service1")
	Consumer(client, "persistent://public/default/query-project-service2")
}
