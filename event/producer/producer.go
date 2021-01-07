package eventproducer

import (
	"context"
	"encoding/json"

	"github.com/DoonaaldArmando/go-library/init/message"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Producer //
type (
	Producer struct{}

	// IProducer //
	IProducer interface {
		InitProducer(topicName string, client pulsar.Client) (pulsar.Producer, error)
		SendMessage(message message.Message, producer pulsar.Producer) (pulsar.MessageID, error)
	}
)

// InitProducer //
func (*Producer) InitProducer(
	topicName string,
	client pulsar.Client,
) (pulsar.Producer, error) {
	return client.
		CreateProducer(
			pulsar.ProducerOptions{
				Topic: topicName,
			},
		)
}

// SendMessage //
func (*Producer) SendMessage(
	message message.Message,
	producer pulsar.Producer,
) (pulsar.MessageID, error) {
	mBytes, serializationError := json.Marshal(message)

	if serializationError != nil {
		return nil, serializationError
	}

	messageID, producerError := producer.Send(
		context.Background(),
		&pulsar.ProducerMessage{
			Payload: []byte(mBytes),
		},
	)

	if producerError != nil {
		return nil, producerError
	}

	return messageID, nil
}
