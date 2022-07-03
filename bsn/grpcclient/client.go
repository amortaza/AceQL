package grpcclient

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/hook"
	"github.com/amortaza/aceql/logger"
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

func GRPC_CallScript(scriptName string) error {
	c := hook.NewHookServiceClient(gGRPC_Connection)

	params := map[string]string{"": ""}

	message := hook.ScriptCall{
		Name:   scriptName,
		Params: params,
	}

	response, err := c.OnRecordUpdate(context.Background(), &message)
	if err != nil {
		return logger.Err(err, logger.GRPC)
	}

	if response != nil {
		return logger.Error(fmt.Sprintf("Response from Server: %s", response.Result), logger.GRPC)
	}

	return nil
}

func GRPC_OnRecordUpdate(scriptName string, grpcMap map[string]string) error {
	c := hook.NewHookServiceClient(gGRPC_Connection)

	params := grpcMap

	message := hook.ScriptCall{
		Name:   scriptName,
		Params: params,
	}

	response, err := c.OnRecordUpdate(context.Background(), &message)
	if err != nil {
		return logger.Err(err, logger.GRPC)
	}

	if response != nil {
		return logger.Error(fmt.Sprintf("Response from Server: %s", response.Result), logger.GRPC)
	}

	return nil
}
