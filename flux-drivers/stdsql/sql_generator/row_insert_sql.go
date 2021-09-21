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

	for column, value := range values.Data {
		sqlValue := generator.valueToSQL(value)

		columnsSQL = fmt.Sprintf("%s, `%s`", columnsSQL, column)
		valuesSQL = fmt.Sprintf("%s, %s", valuesSQL, sqlValue)
	}

	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES(%s);", table, columnsSQL, valuesSQL)
}

func (generator *RowInsert_SqlGenerator) valueToSQL(value interface{}) string {
	sql := ""

	if stringValue, ok := value.(string); ok {
		sql = fmt.Sprintf("'%s'", stringValue)
	} else {
		sql = fmt.Sprintf("%v", value)
	}

	return sql
}
