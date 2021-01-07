package configuration

import (
	"time"

	eventconsumer "github.com/DoonaaldArmando/go-library/event/consumer"
	eventproducer "github.com/DoonaaldArmando/go-library/event/producer"
	"github.com/DoonaaldArmando/go-library/init/gateway"
	"github.com/DoonaaldArmando/go-library/init/message"
)

// InitData //
type InitData struct {
	ServiceName string
	Project     string
	Pulsar      BusConfig
	Querys      map[string]func(message.Message) (interface{}, string)
	Consume     []Consume
}

// Consume //
type Consume struct {
	ServiceName string
	Project     string
	Commands    map[string]func(message.Message)
	Event       map[string]func(message.Message)
}

// BusConfig //
type BusConfig struct {
	Host              string
	OperationTimeout  time.Duration
	ConnectionTimeout time.Duration
}

// InitDataConsumers //
func (initData *InitData) InitDataConsumers() *InitData {
	initData = &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}
	return initData
}

// InitDataServiceName //
func (initData *InitData) InitDataServiceName(serviceName string) *InitData {
	initData.ServiceName = serviceName
	return initData
}

// InitDataProject //
func (initData *InitData) InitDataProject(project string) *InitData {
	initData.Project = project
	return initData
}

// InitDataBusConfig //
func (initData *InitData) InitDataBusConfig(pulsar *BusConfig) *InitData {

	if pulsar == nil {
		return initData
	}

	if pulsar.Host != "" {
		initData.Pulsar.Host = pulsar.Host
	}

	if pulsar.ConnectionTimeout != 0 {
		initData.Pulsar.ConnectionTimeout = pulsar.ConnectionTimeout
	}

	if pulsar.OperationTimeout != 0 {
		initData.Pulsar.OperationTimeout = pulsar.OperationTimeout
	}

	return initData
}

func validateExits(
	initConsume []Consume,
	project string,
	serviceName string,
) (bool, int) {

	hasServiceNameAndProject := false
	position := 0

	for i, consume := range initConsume {
		if consume.Project == project && consume.ServiceName == serviceName {
			hasServiceNameAndProject = true
			position = i
		}
	}

	return hasServiceNameAndProject, position
}

// ListenEvent //
func (initData *InitData) ListenEvent(
	serviceName string,
	project string,
	event string,
	handle func(message.Message),
) *InitData {

	hasServiceNameAndProject, position := validateExits(initData.Consume, project, serviceName)

	if hasServiceNameAndProject == false {
		initData.Consume = append(
			initData.Consume,
			Consume{
				ServiceName: serviceName,
				Project:     project,
				Event:       map[string]func(message.Message){event: handle},
				Commands:    map[string]func(message.Message){},
			},
		)

		return initData
	}

	consume := initData.Consume[position]
	consume.Event[event] = handle
	initData.Consume[position] = consume

	return initData

}

// HandleCommand //
func (initData *InitData) HandleCommand(
	serviceName string,
	project string,
	command string,
	handle func(message.Message),
) *InitData {

	hasServiceNameAndProject, position := validateExits(initData.Consume, project, serviceName)

	if hasServiceNameAndProject == false {
		initData.Consume = append(
			initData.Consume,
			Consume{
				ServiceName: serviceName,
				Project:     project,
				Event:       map[string]func(message.Message){},
				Commands:    map[string]func(message.Message){command: handle},
			},
		)

		return initData
	}

	consume := initData.Consume[position]
	consume.Commands[command] = handle
	initData.Consume[position] = consume

	return initData

}

// HandleQuery //
func (initData *InitData) HandleQuery(
	query string,
	handle func(message.Message) (interface{}, string),
) *InitData {

	initData.Querys[query] = handle
	return initData

}

// Get //
func (initData *InitData) Get() *InitData {
	return initData
}

// Build //
func (initData *InitData) Build() (gateway.IGateway, error) {
	return InitPulsar(initData, &Init{
		&eventconsumer.Consumer{},
		&eventproducer.Producer{},
	})
}
