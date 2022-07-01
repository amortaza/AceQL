package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux/table"
)

type FieldCreate_SqlGenerator struct{}

func NewFieldCreate_SqlGenerator() *FieldCreate_SqlGenerator {
	return &FieldCreate_SqlGenerator{}
}

func (generator *FieldCreate_SqlGenerator) GenerateCreateFieldSQL(table string, field *table.Field) (string, error) {
	sqlType, err := generator.fieldTypeToSQLType(field.Type)
	if err != nil {
		return "", err
	}

	var defaultValue string
	defaultValue, err = generator.fieldTypeToDefaultValue(field.Type)
	if err != nil {
		return "", err
	}

	sql := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s NULL DEFAULT %s;", table, field.Name, sqlType, defaultValue)

	return sql, nil
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToDefaultValue(fieldType table.FieldType) (string, error) {
	if fieldType == table.String {
		return "''", nil
	}

	if fieldType == table.Bool {
		return "0", nil
	}

	if fieldType == table.Number {
		return "0", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype `%s`", fieldType)
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToSQLType(fieldType table.FieldType) (string, error) {
	if fieldType == table.String {
		return "VARCHAR(255)", nil
	}

	if fieldType == table.Bool {
		return "TINYINT", nil
	}

	if fieldType == table.Number {
		return "FLOAT", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype %s", fieldType)
}
