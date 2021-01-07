package commandproducer

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

// CommandProducer //
type CommandProducer struct {
	Producers map[string]pulsar.Producer
	cerr      error
	ccerr     error
	client    pulsar.Client
}

func (commandProducer *CommandProducer) init() {
	commandProducer.Producers = make(map[string]pulsar.Producer)
}

// InitCommandproducer //
func (commandProducer *CommandProducer) InitCommandproducer(topicName string, c pulsar.Client) {

	commandProducer.Producers[topicName], commandProducer.ccerr = c.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})

}
