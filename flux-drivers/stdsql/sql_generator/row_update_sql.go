package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
)

type RowUpdate_SqlGenerator struct {}

func NewRowUpdate_SqlGenerator() *RowUpdate_SqlGenerator {
	return &RowUpdate_SqlGenerator{}
}

func (generator *RowUpdate_SqlGenerator) GenerateSQL(table string, pk string, data *flux.RecordMap) string {
	sql := fmt.Sprintf("UPDATE `%s` SET ", table)
	first := true

	for key, _ := range data.Data {

		// skip primary key
		if key == "x_id" {
			continue
		}

		// add commas (,)
		if first {
			first = false
		} else {
			sql = fmt.Sprintf("%s, ", sql)
		}

		// debug
		valueAsString, _ := data.Get(key)
		fmt.Println( "valueAsString " , valueAsString )

		sql = fmt.Sprintf("%s `%s` = %s", sql, key, generator.valueToSQL(valueAsString))
	}

	return fmt.Sprintf("%s WHERE %s ='%s';", sql, "x_id", pk)
}

func (generator *RowUpdate_SqlGenerator) valueToSQL(value interface{}) string {
	sql := ""

	if stringValue, ok := value.(string); ok {
		sql = fmt.Sprintf("'%s'", stringValue)
	} else {
		sql = fmt.Sprintf("%v", value)
	}

	return sql
}
