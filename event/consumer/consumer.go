package consumer

import (
	"context"
	"encoding/json"

	"github.com/DoonaaldArmando/go-library/init/message"

	"github.com/apache/pulsar-client-go/pulsar"
)

type (
	// Consumer //
	Consumer struct{}

	// IConsumer //
	IConsumer interface {
		InitEventConsumer(
			TopicName string,
			SubscriptionName string,
			client pulsar.Client,
		) (pulsar.Consumer, error)
		Receive(consumer pulsar.Consumer) (message.Message, error)
	}
)

// InitEventConsumer //
func (*Consumer) InitEventConsumer(
	TopicName string,
	SubscriptionName string,
	client pulsar.Client,
) (pulsar.Consumer, error) {

	return client.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            TopicName,
			SubscriptionName: SubscriptionName,
			Type:             pulsar.Shared,
		},
	)

}

// Receive //
func (*Consumer) Receive(consumer pulsar.Consumer) (message.Message, error) {

	msg, err := consumer.Receive(context.Background())

	if err != nil {
		return message.Message{}, err
	}

	var m message.Message

	derr := json.Unmarshal(msg.Payload(), &m)

	if derr != nil {
		return message.Message{}, derr
	}

	consumer.Ack(msg)

	return m, nil
}
