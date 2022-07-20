package grpc_client

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_hook"
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

func GRPC_CallScript(directory, scriptName string, params map[string]string) error {
	c := grpc_hook.NewHookServiceClient(gGRPC_Connection)

	if params == nil {
		params = map[string]string{"testParam1": "value1"}
	}

	scriptRequest := grpc_hook.ScriptRequest{
		ScriptPath: directory + "/" + scriptName + ".js",
		Params:     params,
	}

	scriptResponse, err := c.OnScriptCall(context.Background(), &scriptRequest)
	if err != nil {
		return logger.Err(err, "GRPC")
	}

	if scriptResponse != nil {
		logger.Info(fmt.Sprintf("Answer: %s", scriptResponse.Answer), "GRPC")
	}

	return nil
}

func GRPC_ImportSet(adapter string, params, ctx map[string]string) error {
	c := grpc_hook.NewHookServiceClient(gGRPC_Connection)

	if params == nil {
		params = map[string]string{"testParam1": "value1"}
	}

	importsetRequest := grpc_hook.ImportSetRequest{
		Page:     12,
		Pagesize: 50,
		Adapter:  adapter,
	}

	importsetResponse, err := c.OnImportSet(context.Background(), &importsetRequest)
	if err != nil {
		return logger.Err(err, "GRPC")
	}

	if importsetResponse != nil {
		schema := importsetResponse.Schema
		rows := importsetResponse.Rows

		logger.Info(fmt.Sprintf("ImportSet Schema: %d", len(schema)), "GRPC")
		logger.Info(fmt.Sprintf("ImportSet Rows: %d", len(rows)), "GRPC")

		for i, field := range schema {
			logger.Info(fmt.Sprintf("Schema: %d %s", i, field), "GRPC")
		}

		for i, row := range rows {
			logger.Info(fmt.Sprintf("ImportSet Rows: %d %s", i, row.Values), "GRPC")
			//logger.Info(fmt.Sprintf("ImportSet cool"), "GRPC")
		}
	}

	return nil
}
