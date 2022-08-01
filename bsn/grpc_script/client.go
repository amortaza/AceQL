package grpc_script

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strings"
)

var gGRPC_Connection *grpc.ClientConn

const (
	BUSINESS_RULE_ON_CREATE = "on_create"
	BUSINESS_RULE_ON_READ   = "on_read"
	BUSINESS_RULE_ON_UPDATE = "on_update"
	BUSINESS_RULE_ON_DELETE = "on_delete"
)

func init() {
	var err error
	gGRPC_Connection, err = grpc.Dial(":50051", grpc.WithInsecure()) // :9000
	if err != nil {
		log.Fatalf("could not connect to NodeJS gRPC Server: %s", err)
	}

	// TODO: use signals to close(), see https://gobyexample.com/signals
	//defer conn.Close()
}

func GRPC_CallBusinessRule(directory, scriptName, table, recordId string, originals, current map[string]string) error {
	LOG_SOURCE := "GRPC_CallBusinessRule"

	c := NewScriptServiceClient(gGRPC_Connection)

	logger.Info("GRPC_CallBusinessRule: current , see below", LOG_SOURCE)
	for k, v := range (current) {
		logger.Info("GRPC_CallBusinessRule: current "+k+" "+v, LOG_SOURCE)
	}

	businessRuleRequest := BusinessRuleRequest{
		ScriptPath: directory + "/" + scriptName + ".js",
		Action:     BUSINESS_RULE_ON_UPDATE,
		Table:      table,
		RecordId:   recordId,
		Originals:  originals,
		Current:    current,
	}

	businessRuleResponse, err := c.OnBusinessRule(context.Background(), &businessRuleRequest)
	if err != nil {
		return logger.Err(err, "GRPC")
	}

	if businessRuleResponse != nil {
		if businessRuleResponse.Fault != "" {
			logger.Info("BusinessRule Error: "+err.Error(), LOG_SOURCE)

		} else if businessRuleResponse.CancelAction != "" {
			logger.Info("BusinessRule Cancel Action: "+businessRuleResponse.CancelAction, LOG_SOURCE)

		} else if businessRuleResponse.MessageToDisplay != "" {
			logger.Info("BusinessRule Message: "+businessRuleResponse.MessageToDisplay, LOG_SOURCE)

		} else if businessRuleResponse.ErrorToDisplay != "" {
			logger.Info("BusinessRule Error Message: "+businessRuleResponse.ErrorToDisplay, LOG_SOURCE)

		}
	}

	return nil
}

func GRPC_CallScript(directory, scriptName string, params map[string]string) error {
	c := NewScriptServiceClient(gGRPC_Connection)

	if params == nil {
		params = map[string]string{"testParam1": "value1"}
	}

	scriptRequest := ScriptRequest{
		ScriptPath: directory + "/" + scriptName + ".js",
		Params:     params,
	}

	scriptResponse, err := c.OnScriptCall(context.Background(), &scriptRequest)
	if err != nil {
		return logger.Err(err, "GRPC")
	}

	if scriptResponse != nil {
		m := scriptResponse.Answer
		for k, v := range m {
			logger.Info("Answer "+k+" : "+v, "GRPC: Script Response")
		}
		logger.Info(fmt.Sprintf("Answer: %s", scriptResponse.Answer), "GRPC: Script Response")
	}

	return nil
}

func GRPC_ImportSet(importsetName string) error {
	c := NewScriptServiceClient(gGRPC_Connection)

	importset, err := getImportSetRecord(importsetName)

	adapter, err := importset.Get("x_adapter")
	if err != nil {
		return err
	}

	target, err := importset.Get("x_target_table")
	if err != nil {
		return err
	}

	mappings_str, err := importset.Get("x_mappings")
	if err != nil {
		return err
	}

	importsetRequest := ImportSetRequest{
		Page:     12,
		Pagesize: 50,
		Adapter:  adapter,
	}

	importsetResponse, err := c.OnImportSet(context.Background(), &importsetRequest)
	if err != nil {
		return logger.Err(err, "GRPC")
	}

	if importsetResponse != nil {
		fields := importsetResponse.Fields
		rows := importsetResponse.Rows

		//createImportTable(table, fields)

		if err := insertRows(target, mappings_str, fields, rows); err != nil {
			return err
		}
	}

	return nil
}

func getImportSetRecord(importsetName string) (*flux.Record, error) {
	record, err := stdsql.NewRecord("x_importset")
	if err != nil {
		return nil, err
	}

	record.AddEq("x_name", importsetName)

	count, err2 := record.Query()
	if err2 != nil {
		return nil, err2
	}

	if count > 1 {
		return nil, logger.Error(fmt.Sprintf("found duplicatte %d importsets \"%s\"", count, importsetName), "importset")
	}

	if count == 0 {
		return nil, logger.Error(fmt.Sprintf("found no importsets named \"%s\"", importsetName), "importset")
	}

	hasNext, err3 := record.Next()
	if err3 != nil {
		return nil, err3
	}

	if !hasNext {
		return nil, logger.Error(fmt.Sprintf("found importset named \"%s\", but record.Next() returned false", importsetName), "importset")
	}

	return record, nil
}

func insertRows(table, mappings_str string, fields []string, rows []*Row) error {
	mappings := parseMappings(mappings_str)

	for _, row := range rows {
		if err := insertRow(table, mappings, fields, row); err != nil {
			return err
		}
	}

	return nil
}

func parseMappings(str string) map[string]string {
	str = strings.ReplaceAll(str, " ", "")
	parts := strings.Split(str, ",")

	mapping := make(map[string]string)

	for _, part := range parts {
		left_right := strings.Split(part, ">>")
		mapping[left_right[0]] = left_right[1]
	}

	return mapping
}

func insertRow(table string, mapping map[string]string, fields []string, row *Row) error {
	gr, err := stdsql.NewRecord(table)
	if err != nil {
		return err
	}

	//logger.Info(fmt.Sprintf("Hi field count %d", len(fields)), "GRPC")
	//logger.Info(fmt.Sprintf("Hi value count %d", len(row.Values)), "GRPC")

	for i, field := range fields {
		logger.Info(fmt.Sprintf("value %s", row.Values[i]), "GRPC")

		mappedField, ok := mapping[field]
		if !ok {
			return logger.Error(fmt.Sprintf("mapping not found \"%s\"", field), "importset")
		}

		err := gr.Set(mappedField, row.Values[i])
		if err != nil {
			return err
		}
	}

	_, err = gr.Insert()

	return err
}

/*
func createImportTable(tablename string, fields []*Field) error {
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
*/
