package sql_generator

import (
	"fmt"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
)

type FieldCreate_SqlGenerator struct{}

func NewFieldCreate_SqlGenerator() *FieldCreate_SqlGenerator {
	return &FieldCreate_SqlGenerator{}
}

func (generator *FieldCreate_SqlGenerator) GenerateCreateFieldSQL(table string, field *dbschema.Field) (string, error) {
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

func (generator *FieldCreate_SqlGenerator) fieldTypeToDefaultValue(fieldType dbschema.FieldType) (string, error) {
	if fieldType == dbschema.String {
		return "''", nil
	}

	if fieldType == dbschema.Bool {
		return "0", nil
	}

	if fieldType == dbschema.Number {
		return "0", nil
	}

	err := logger.Error(fmt.Sprintf("unrecognized fieldtype `%s`", fieldType), "FieldCreate_SqlGenerator.fieldTypeToDefaultValue")
	return "", err
}

func (generator *FieldCreate_SqlGenerator) fieldTypeToSQLType(fieldType dbschema.FieldType) (string, error) {
	if fieldType == dbschema.String {
		return "VARCHAR(255)", nil
	}

	if fieldType == dbschema.Bool {
		return "VARCHAR(15)", nil
	}

	if fieldType == dbschema.Number {
		return "VARCHAR(31)", nil
	}

	return "", logger.Error(fmt.Sprintf("unrecognized fieldtype %s", fieldType), "FieldCreate_SqlGenerator.fieldTypeToSQLType")
}
