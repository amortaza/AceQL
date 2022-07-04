package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/logger"
)

type RowInsert_SqlGenerator struct{}

func NewRowInsert_SqlGenerator() *RowInsert_SqlGenerator {
	return &RowInsert_SqlGenerator{}
}

func (generator *RowInsert_SqlGenerator) GenerateInsertSQL(table string, newId string, values *flux.RecordMap) (string, error) {
	columnsSQL := "`x_id`"
	valuesSQL := fmt.Sprintf("'%s'", newId)

	for column, typedValue := range values.Data {
		if column == "x_id" {
			continue
		}

		var sqlValue string
		var err error
		if sqlValue, err = generator.typedValueToSQL(typedValue); err != nil {
			return "invalid sql", err
		}

		columnsSQL = fmt.Sprintf("%s, `%s`", columnsSQL, column)
		valuesSQL = fmt.Sprintf("%s, %s", valuesSQL, sqlValue)
	}

	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES(%s);", table, columnsSQL, valuesSQL), nil
}

func (generator *RowInsert_SqlGenerator) typedValueToSQL(typedValue *flux.TypedValue) (string, error) {
	sql := ""

	value := typedValue.GetValue()

	if typedValue.IsString() {
		// todo value should be escaped
		sql = fmt.Sprintf("'%s'", value)

	} else if typedValue.IsNumber() {
		sql = fmt.Sprintf("%s", value)

	} else if typedValue.IsBool() {
		sql = fmt.Sprintf("%s", value)

	} else {
		return "invalid sql", logger.Error("unrecognized type, cant even see what type it is it", "???")
	}

	return sql, nil
}
