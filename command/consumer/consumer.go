package commandconsumer

import (
	"context"
	"encoding/json"

	"github.com/DoonaaldArmando/go-library/init/message"

	"github.com/apache/pulsar-client-go/pulsar"
)

var consumer pulsar.Consumer
var cerr error
var ccerr error
var client pulsar.Client

// InitCommandConsumer //
func InitCommandConsumer(topicName string, subscriptionName string, c pulsar.Client) {

	consumer, ccerr = c.Subscribe(pulsar.ConsumerOptions{
		Topic:            topicName,
		SubscriptionName: subscriptionName,
		Type:             pulsar.Shared,
	})

}

// Receive //
func Receive() (message.Message, error) {

	msg, err := consumer.Receive(context.Background())

	if err != nil {
		return message.Message{}, err
	}

	if ccerr != nil {
		return message.Message{}, ccerr
	}

	var m message.Message

	derr := json.Unmarshal(msg.Payload(), &m)

	if derr != nil {
		return message.Message{}, derr
	}

	consumer.Ack(msg)

	return m, nil
}
