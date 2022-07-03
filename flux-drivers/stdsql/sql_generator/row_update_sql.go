package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
)

type RowUpdate_SqlGenerator struct{}

func NewRowUpdate_SqlGenerator() *RowUpdate_SqlGenerator {
	return &RowUpdate_SqlGenerator{}
}

func (generator *RowUpdate_SqlGenerator) GenerateSQL(table string, pk string, recordmap *flux.RecordMap) (string, error) {
	sql := fmt.Sprintf("UPDATE `%s` SET ", table)
	first := true

	for fieldname, _ := range recordmap.Data {

		// skip primary key
		if fieldname == "x_id" {
			continue
		}

		// add commas (,)
		if first {
			first = false
		} else {
			sql = fmt.Sprintf("%s,", sql)
		}

		fieldSQL, err := writeFieldUpdateSQL(fieldname, recordmap)
		if err != nil {
			return "", err
		}

		sql = fmt.Sprintf("%s %s", sql, fieldSQL)
	}

	return fmt.Sprintf("%s WHERE %s ='%s';", sql, "x_id", pk), nil
}

func writeFieldUpdateSQL(fieldname string, recordmap *flux.RecordMap) (string, error) {
	var sql, valueAsString string
	var err error

	isString, err := recordmap.IsFieldString(fieldname)
	if err != nil {
		return "", err
	}

	valueAsString, err = recordmap.GetFieldValue(fieldname)
	if err != nil {
		return "", err
	}

	if isString {
		// todo valueAsString should be encoded
		sql = fmt.Sprintf("`%s` = '%s'", fieldname, valueAsString)

	} else {
		sql = fmt.Sprintf("`%s` = %s", fieldname, valueAsString)
	}

	return sql, nil
}
