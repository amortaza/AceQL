package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
)

type RowUpdate_SqlGenerator struct {}

func NewRowUpdate_SqlGenerator() *RowUpdate_SqlGenerator {
	return &RowUpdate_SqlGenerator{}
}

func (generator *RowUpdate_SqlGenerator) GenerateSQL(table string, pk string, data *flux.RecordMap) (string, error) {
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

		var valueAsString string
		var err error

		isString, err := data.IsString(key)
		if err != nil {
			return "", err
		}

		isNumber, err := data.IsNumber(key)
		if err != nil {
			return "", err
		}

		isBool, err := data.IsBool(key)
		if err != nil {
			return "", err
		}

		if isString {
			valueAsString, err = data.Get(key)
			if err != nil {
				return "", err
			}

			sql = fmt.Sprintf("%s `%s` = '%s'", sql, key, valueAsString)

		} else if isNumber {
			valueAsFloat, err := data.GetNumber(key)
			if err != nil {
				return "", err
			}

			valueAsString = fmt.Sprintf("%f", valueAsFloat)

			sql = fmt.Sprintf("%s `%s` = %s", sql, key, valueAsString)

		} else if isBool {
			valueAsBool, err := data.GetBool(key)
			if err != nil {
				return "", err
			}

			if valueAsBool {
				valueAsString = "true"
			} else {
				valueAsString = "false"
			}

			sql = fmt.Sprintf("%s `%s` = %s", sql, key, valueAsString)
		}
	}

	return fmt.Sprintf("%s WHERE %s ='%s';", sql, "x_id", pk), nil
}

