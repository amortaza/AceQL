package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
)

type RowInsert_SqlGenerator struct {}

func NewRowInsert_SqlGenerator() *RowInsert_SqlGenerator {
	return &RowInsert_SqlGenerator{}
}

func (generator *RowInsert_SqlGenerator) GenerateInsertSQL(table string, newId string, values *flux.RecordMap) string {
	columnsSQL := "`x_id`"
	valuesSQL := fmt.Sprintf("'%s'", newId)

	for column, typedValue := range values.Data {
		if column == "x_id" {
			continue
		}

		sqlValue := generator.typedValueToSQL(typedValue)

		columnsSQL = fmt.Sprintf("%s, `%s`", columnsSQL, column)
		valuesSQL = fmt.Sprintf("%s, %s", valuesSQL, sqlValue)
	}

	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES(%s);", table, columnsSQL, valuesSQL)
}

func (generator *RowInsert_SqlGenerator) typedValueToSQL(typedValue *flux.TypedValue) string {
	sql := ""

	if typedValue.IsString() {
		sql = fmt.Sprintf("'%s'", typedValue.GetString())

	} else if typedValue.IsNumber() {
		sql = fmt.Sprintf("%1.2f", typedValue.GetNumber())

	} else if typedValue.IsBool() {
		sql = fmt.Sprintf("%t", typedValue.GetBool())

	} else {
		panic("TypedValue type is unrecognized in typedValueToSQL()")
	}

	return sql
}
