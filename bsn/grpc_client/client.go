package grpc_client

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_hook"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
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

func GRPC_ImportSet(adapter string) error {
	c := grpc_hook.NewHookServiceClient(gGRPC_Connection)

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
		table := importsetResponse.Table
		fields := importsetResponse.Fields
		rows := importsetResponse.Rows

		//createImportTable(table, fields)

		if err := insertRows(table, fields, rows); err != nil {
			return err
		}
	}

	return nil
}

func insertRows(table string, fields []*grpc_hook.Field, rows []*grpc_hook.Row) error {
	for _, row := range rows {
		if err := insertRow(table, fields, row); err != nil {
			return err
		}
	}

	return nil
}

func insertRow(table string, fields []*grpc_hook.Field, row *grpc_hook.Row) error {
	gr, err := stdsql.NewRecord(table)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Hi field count %d", len(fields)), "GRPC")
	logger.Info(fmt.Sprintf("Hi value count %d", len(row.Values)), "GRPC")

	for i, field := range fields {
		logger.Info(fmt.Sprintf("value %s", row.Values[i]), "GRPC")
		err := gr.Set(field.Fieldname, row.Values[i])
		if err != nil {
			return err
		}
	}

	_, err = gr.Insert()

	return err
}

func createImportTable(tablename string, fields []*grpc_hook.Field) error {
	table := dbschema.NewTable(tablename)

	for _, field := range fields {
		fieldtype, err := dbschema.GetFieldTypeByName(field.Fieldtype)
		if err != nil {
			return err
		}

		table.AddField(field.Fieldname, field.Fieldname, fieldtype)
	}

	return createStdSqlTable(table)
}

func createStdSqlTable(tableschema *dbschema.Table) error {
	schema := stdsql.NewSchema()

	if err := schema.CreateTable_withName(tableschema.Name(), tableschema.Label(), tableschema.Name() != "x_schema"); err != nil {
		return err
	}

	for _, field := range tableschema.Fields() {
		if err := schema.CreateField(tableschema.Name(), field, tableschema.Name() != "x_schema"); err != nil {
			return err
		}
	}

	return nil
}
