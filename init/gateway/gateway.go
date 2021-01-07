package gateway

import (
	"fmt"
	eventproducer "pulsar-go/event/producer"
	"pulsar-go/init/message"

	"github.com/apache/pulsar-client-go/pulsar"
)

// QuerysSend //
var QuerysSend map[string]func(message.Message)

func init() {
	QuerysSend = map[string]func(message.Message){}
}

// Gateway //
type Gateway struct {
	client      pulsar.Client
	serviceName string
	projectName string
	producer    eventproducer.IProducer
}

// CreateGateway //
func CreateGateway(client pulsar.Client,
	serviceName string,
	projectName string,
	producer eventproducer.IProducer,
) *Gateway {
	return &Gateway{
		client:      client,
		serviceName: serviceName,
		projectName: projectName,
		producer:    producer,
	}
}

// IGateway //
type IGateway interface {
	SendQuery(
		project string,
		service string,
		userAgent string,
		user string,
		eventCode string,
		correlationID string,
		handleQuery func(message.Message),
		data interface{},
	) (pulsar.MessageID, error)
	SendEvent(
		project string,
		service string,
		userAgent string,
		user string,
		eventCode string,
		correlationID string,
		data interface{},
	) (pulsar.MessageID, error)
}

// SendEvent //
func (gateway Gateway) SendEvent(
	project string,
	service string,
	userAgent string,
	user string,
	eventCode string,
	correlationID string,
	data interface{},
) (pulsar.MessageID, error) {

	producer, errorProducer := gateway.
		producer.
		InitProducer(
			fmt.Sprintf(
				"%s-%s-%s",
				message.EVENT,
				project,
				service,
			),
			gateway.client,
		)

	if errorProducer != nil {
		return nil, errorProducer
	}

	return gateway.producer.SendMessage(
		message.Message{}.NewMessage(
			userAgent,
			user,
			message.JSON,
			eventCode,
			correlationID,
			fmt.Sprintf("%s-%s", gateway.projectName, gateway.serviceName),
			message.EVENT,
			data,
		),
		producer,
	)

}

// SendQuery //
func (gateway Gateway) SendQuery(
	project string,
	service string,
	userAgent string,
	user string,
	eventCode string,
	correlationID string,
	handleQuery func(message.Message),
	data interface{},
) (pulsar.MessageID, error) {

	producer, errorProducer := gateway.
		producer.
		InitProducer(
			fmt.Sprintf("%s-%s-%s", message.QUERY, project, service),
			gateway.client,
		)

	if errorProducer != nil {
		return nil, errorProducer
	}

	QuerysSend[fmt.Sprintf("%s-%s", eventCode, message.RESPONSE)] = handleQuery

	return gateway.producer.SendMessage(
		message.Message{}.NewMessage(
			userAgent,
			user,
			message.JSON,
			eventCode,
			correlationID,
			fmt.Sprintf("%s-%s", gateway.projectName, gateway.serviceName),
			message.QUERY,
			data,
		),
		producer,
	)

}
