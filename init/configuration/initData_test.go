package configuration

import (
	"fmt"
	"pulsar-go/init/message"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type person struct{}

func handleQuery(message message.Message) (interface{}, string) {
	fmt.Println("print!")
	return person{}, "test!"
}

func handleFunction(message message.Message) {
	fmt.Println("print!")
}

// TestListenEvent //
func TestListenEvent(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()
	registerConsumers.ListenEvent("ServiceName", "Project", "serviceName.project.request", handleFunction)
	initDataResult := &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys: map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{
			{
				ServiceName: "ServiceName",
				Project:     "Project",
				Commands:    map[string]func(message.Message){},
				Event: map[string]func(message.Message){
					"serviceName.project.request": handleFunction,
				},
			},
		},
	}
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "New configuration")

	registerConsumers.ListenEvent("ServiceName", "Project", "serviceName.project.save", handleFunction)
	initDataResult.Consume[0].Event["serviceName.project.save"] = handleFunction
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Diferent configuration")

	registerConsumers.ListenEvent("OtherServiceName", "OtherProject", "otherserviceName.otherproject.save", handleFunction)
	initDataResult.Consume = append(initDataResult.Consume, Consume{
		ServiceName: "OtherServiceName",
		Project:     "OtherProject",
		Commands:    map[string]func(message.Message){},
		Event: map[string]func(message.Message){
			"otherserviceName.otherproject.save": handleFunction,
		},
	},
	)
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Add configuration in previus configuration")
}

// TestRegisterServiceName //
func TestRegisterServiceName(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()

	initDataResult := &InitData{
		ServiceName: "ServiceName",
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}

	registerConsumers.InitDataServiceName("ServiceName")

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Add configuration")

}

func TestRegisterProject(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()

	initDataResult := &InitData{
		Project: "Project",
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}

	registerConsumers.InitDataProject("Project")

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Add configuration")

}

// TestHandleCommand //
func TestHandleCommand(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()

	registerConsumers.HandleCommand("ServiceName", "Project", "serviceName.project.request", handleFunction)
	initDataResult := &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys: map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{
			{
				ServiceName: "ServiceName",
				Project:     "Project",
				Commands: map[string]func(message.Message){
					"serviceName.project.request": handleFunction,
				},
				Event: map[string]func(message.Message){},
			},
		},
	}
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "New configuration")

	registerConsumers.HandleCommand("ServiceName", "Project", "serviceName.project.save", handleFunction)
	initDataResult.Consume[0].Commands["serviceName.project.save"] = handleFunction
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Diferent configuration")

	registerConsumers.HandleCommand("OtherServiceName", "OtherProject", "otherserviceName.otherproject.save", handleFunction)
	initDataResult.Consume = append(initDataResult.Consume, Consume{
		ServiceName: "OtherServiceName",
		Project:     "OtherProject",
		Commands: map[string]func(message.Message){
			"otherserviceName.otherproject.save": handleFunction,
		},
		Event: map[string]func(message.Message){},
	},
	)
	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Add configuration in previus configuration")

}

// TestHandleCommand //
func TestHandleQuery(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()
	registerConsumers.HandleQuery("serviceName.project.request", handleQuery)
	initDataResult := &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys: map[string]func(message.Message) (interface{}, string){
			"serviceName.project.request": handleQuery,
		},
		Consume: []Consume{},
	}

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Add new configuration")

}

func TestRegisterBusConfig(t *testing.T) {

	var initData *InitData = &InitData{}
	registerConsumers := initData.InitDataConsumers()

	busConfig := &BusConfig{
		Host:              "anyHost",
		OperationTimeout:  10 * time.Second,
		ConnectionTimeout: 10 * time.Second,
	}

	registerConsumers.InitDataBusConfig(busConfig)

	initDataResult := &InitData{
		Pulsar: BusConfig{
			Host:              "anyHost",
			OperationTimeout:  10 * time.Second,
			ConnectionTimeout: 10 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "Change all values")

	registerConsumers = initData.InitDataConsumers()
	busConfig = &BusConfig{}

	registerConsumers.InitDataBusConfig(busConfig)

	initDataResult = &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "No change values")

	registerConsumers = initData.InitDataConsumers()
	registerConsumers.InitDataBusConfig(nil)

	initDataResult = &InitData{
		Pulsar: BusConfig{
			Host:              "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		},
		Querys:  map[string]func(message.Message) (interface{}, string){},
		Consume: []Consume{},
	}

	assert.Equal(t, fmt.Sprintf("%#v", initDataResult), fmt.Sprintf("%#v", registerConsumers), "No change values")

}

func TestGet(t *testing.T) {
	var initData *InitData = &InitData{}
	assert.Equal(t, &InitData{}, initData.Get())
}
