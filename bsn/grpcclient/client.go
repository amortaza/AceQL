package grpcclient

import (
	"github.com/amortaza/aceql/bsn/hook"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var gGRPC_Connection *grpc.ClientConn

func init() {
	var err error
	gGRPC_Connection, err = grpc.Dial(":50051", grpc.WithInsecure()) // :9000
	if err != nil {
		log.Fatalf("could not connect to NodeJS gRPC Server: %s", err)
	}

	// TODO: use signals to close(), see https://gobyexample.com/signals
	//defer conn.Close()
}

func GRPC_CallScript(scriptName string) {
	c := hook.NewHookServiceClient(gGRPC_Connection)

	params := map[string]string{"": ""}

	message := hook.ScriptCall{
		Name:   scriptName,
		Params: params,
	}

	response, err := c.OnRecordUpdate(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling : %s", err)
	}

	log.Printf("Response from Server: %s", response.Result)
}

func GRPC_OnRecordUpdate(scriptName string, grpcMap map[string]string) {
	c := hook.NewHookServiceClient(gGRPC_Connection)

	params := grpcMap

	message := hook.ScriptCall{
		Name:   scriptName,
		Params: params,
	}

	response, err := c.OnRecordUpdate(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling : %s", err)
	}

	log.Printf("Response from Server: %s", response.Result)
}
