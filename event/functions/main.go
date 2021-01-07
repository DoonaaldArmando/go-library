package main

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar/pulsar-function-go/pf"
)

// HandleResponse //
func HandleResponse(ctx context.Context, in []byte) ([]byte, error) {
	res := append(in, 110)
	pf.FromContext(ctx)
	if fc, ok := pf.FromContext(ctx); ok {
		fc.NewOutputMessage("persistent://public/default/query-project-service1").
			Send(context.Background(), &pulsar.ProducerMessage{
				Payload: append(in, 111),
			})
	}
	return res, nil
}

func main() {
	pf.Start(HandleResponse)
}
