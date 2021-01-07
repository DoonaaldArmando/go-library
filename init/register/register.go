package register

import (
	"pulsar-go/init/configuration"
	"pulsar-go/init/gateway"
	"pulsar-go/init/message"
)

// Register //
type Register interface {
	InitDataConsumers() *configuration.InitData
	InitDataServiceName(serviceName string) *configuration.InitData
	InitDataProject(project string) *configuration.InitData
	InitDataBusConfig(pulsar *configuration.BusConfig) *configuration.InitData
	ListenEvent(
		serviceName string,
		project string,
		event string,
		handle func(message.Message),
	) *configuration.InitData
	HandleCommand(
		serviceName string,
		project string,
		command string,
		handle func(message.Message),
	) *configuration.InitData
	HandleQuery(
		query string,
		handle func(message.Message) (interface{}, string),
	) *configuration.InitData
	Build() (gateway.IGateway, error)
	Get() *configuration.InitData
}

// InitRegister //
func InitRegister() Register {
	return &configuration.InitData{}
}
