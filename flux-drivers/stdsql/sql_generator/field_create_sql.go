package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux/tableschema"
)

type FieldCreate_SqlGenerator struct{}

func NewFieldCreate_SqlGenerator() *FieldCreate_SqlGenerator {
	return &FieldCreate_SqlGenerator{}
}

func (generator *FieldCreate_SqlGenerator) GenerateCreateFieldSQL(table string, field *tableschema.Field) (string, error) {
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

func (generator *FieldCreate_SqlGenerator) fieldTypeToDefaultValue(fieldType tableschema.FieldType) (string, error) {
	if fieldType == tableschema.String {
		return "''", nil
	}

	if fieldType == tableschema.Bool {
		return "0", nil
	}

	if fieldType == tableschema.Number {
		return "0", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype `%s`", fieldType)
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToSQLType(fieldType tableschema.FieldType) (string, error) {
	if fieldType == tableschema.String {
		return "VARCHAR(255)", nil
	}

	if fieldType == tableschema.Bool {
		return "VARCHAR(15)", nil
		//return "TINYINT", nil
	}

	if fieldType == tableschema.Number {
		return "VARCHAR(31)", nil
		//return "FLOAT", nil
	}

	return "", fmt.Errorf("unrecognized fieldtype %s", fieldType)
}
