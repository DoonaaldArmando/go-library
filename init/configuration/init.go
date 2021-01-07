package configuration

import (
	"fmt"
	"pulsar-go/init/gateway"
	"pulsar-go/init/message"
	"strings"

	EventConsumer "pulsar-go/event/consumer"
	eventproducer "pulsar-go/event/producer"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Init //
type Init struct {
	consumer EventConsumer.IConsumer
	producer eventproducer.IProducer
}

// InitPulsar //
func InitPulsar(initData *InitData, init *Init) (gateway.IGateway, error) {

	clientPulsar, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               initData.Pulsar.Host,
		OperationTimeout:  initData.Pulsar.ConnectionTimeout,
		ConnectionTimeout: initData.Pulsar.ConnectionTimeout,
	})

	if err != nil {
		return nil, err
	}

	init.initQueryReceive(
		fmt.Sprintf("%s-%s-%s", message.QUERY, initData.Project, initData.ServiceName),
		fmt.Sprintf("%s-%s", initData.Project, initData.ServiceName),
		initData.Querys,
		initData.Project,
		initData.ServiceName,
		clientPulsar,
	)

	init.initEventReceive(
		initData.Consume,
		clientPulsar,
	)

	return gateway.CreateGateway(
			clientPulsar,
			initData.ServiceName,
			initData.Project,
			init.producer,
		),
		nil
}

func (init *Init) initEventReceive(
	consumers []Consume,
	client pulsar.Client,
) error {

	for _, consumer := range consumers {
		go func(consumer Consume) {
			for {

				initializateConsumer, _ := init.consumer.
					InitEventConsumer(
						fmt.Sprintf(
							"%s-%s-%s",
							message.EVENT,
							consumer.Project,
							consumer.ServiceName,
						),
						fmt.Sprintf(
							"%s-%s",
							consumer.Project,
							consumer.ServiceName,
						),
						client,
					)
				messageReceived, _ := init.consumer.Receive(initializateConsumer)
				proccess, ok := consumer.Event[messageReceived.Header.EventCode]

				if ok {
					proccess(messageReceived)

				}

			}
		}(consumer)
	}

	return nil
}

func (init *Init) initQueryReceive(
	topicName string,
	subscriptionName string,
	queries map[string]func(message.Message) (interface{}, string),
	project string,
	service string,
	client pulsar.Client,
) error {

	initializateConsumer, errorConsumer := init.consumer.
		InitEventConsumer(
			topicName,
			subscriptionName,
			client,
		)

	if errorConsumer != nil {
		return errorConsumer
	}

	go func(initializateConsumer pulsar.Consumer) {
		for {
			messageReceived, _ := init.consumer.Receive(initializateConsumer)

			proccessReceive, okReceive := gateway.QuerysSend[messageReceived.Header.EventCode]
			if okReceive {
				proccessReceive(messageReceived)
				delete(gateway.QuerysSend, messageReceived.Header.EventCode)
			}

			proccess, ok := queries[messageReceived.Header.EventCode]
			if ok {

				producer, _ := init.producer.InitProducer(fmt.Sprintf(
					"%s-%s-%s",
					message.QUERY,
					strings.Split(messageReceived.Audit.Source, "-")[0],
					strings.Split(messageReceived.Audit.Source, "-")[1],
				), client)

				dataMessage, user := proccess(messageReceived)

				init.producer.
					SendMessage(message.Message{}.
						NewMessage(
							messageReceived.Audit.UserAgent,
							user,
							message.JSON,
							fmt.Sprintf(
								"%s-%s",
								messageReceived.Header.EventCode,
								message.RESPONSE,
							),
							messageReceived.Header.CorrelationID,
							fmt.Sprintf("%s-%s", project, service),
							message.QUERY,
							dataMessage,
						),
						producer,
					)

			}

		}

	}(initializateConsumer)

	return nil
}
