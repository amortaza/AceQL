package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux/relation_type"
)

type FieldCreate_SqlGenerator struct {}

func NewFieldCreate_SqlGenerator() *FieldCreate_SqlGenerator {
	return &FieldCreate_SqlGenerator{}
}

func (generator *FieldCreate_SqlGenerator) GenerateCreateFieldSQL(table string, field * relation_type.Field) (string, error) {
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

func (generator *FieldCreate_SqlGenerator) fieldTypeToDefaultValue(fieldType relation_type.FieldType) (string, error) {
	if fieldType == relation_type.String {
		return "NULL", nil
	}

	if fieldType == relation_type.Bool {
		return "0", nil
	}

	if fieldType == relation_type.Number {
		return "0", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype `%s`", fieldType)
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToSQLType(fieldType relation_type.FieldType) (string, error) {
	if fieldType == relation_type.String {
		return "VARCHAR(255)", nil
	}

	if fieldType == relation_type.Bool {
		return "TINYINT", nil
	}

	if fieldType == relation_type.Number {
		return "DECIMAL(10,5)", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype %s", fieldType)
}
