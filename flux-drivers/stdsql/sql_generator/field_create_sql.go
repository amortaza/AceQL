package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux/relations"
)

type FieldCreate_SqlGenerator struct {}

func NewFieldCreate_SqlGenerator() *FieldCreate_SqlGenerator {
	return &FieldCreate_SqlGenerator{}
}

func (generator *FieldCreate_SqlGenerator) GenerateCreateFieldSQL(table string, field * relations.Field) (string, error) {
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

func (generator *FieldCreate_SqlGenerator) fieldTypeToDefaultValue(fieldType relations.FieldType) (string, error) {
	if fieldType == relations.String {
		return "NULL", nil
	}

	if fieldType == relations.Bool {
		return "0", nil
	}

	if fieldType == relations.Number {
		return "0", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype `%s`", fieldType)
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToSQLType(fieldType relations.FieldType) (string, error) {
	if fieldType == relations.String {
		return "VARCHAR(255)", nil
	}

	if fieldType == relations.Bool {
		return "TINYINT", nil
	}

	if fieldType == relations.Number {
		return "FLOAT", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype %s", fieldType)
}
