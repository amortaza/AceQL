package flux

import (
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
)

type Schema struct {
	journalist dbschema.Journalist
	crud       CRUD
}

func NewSchema(journalist dbschema.Journalist, crud CRUD) *Schema {
	return &Schema{
		journalist: journalist,
		crud:       crud,
	}
}

func (schema *Schema) Close() error {
	return schema.crud.Close()
}

func (schema *Schema) CreateRelation_withFields(relation *dbschema.Table, journal bool) error {
	if err := schema.CreateTable_withName(relation.Name(), relation.Label(), journal); err != nil {
		return err
	}

	for _, field := range relation.Fields() {
		if err := schema.CreateField(relation.Name(), field, journal); err != nil {
			return err
		}
	}

	return nil
}

func (schema *Schema) CreateTable_withName(name string, label string, journal bool) error {

	// the order here is importan because if we are creating 'x_schema'
	// we want to create the table first THEN journal it

	err := schema.crud.CreateTable(name)
	if err != nil {
		return err
	}

	if !journal {
		return nil
	}

	return schema.journalist.CreateTable(name, label)
}

func (schema *Schema) DeleteRelation(name string) error {
	if err := schema.journalist.DeleteTable(name); err != nil {
		return err
	}

	return schema.crud.DeleteTable(name)
}

func (schema *Schema) CreateField(tableName string, field *dbschema.Field, journal bool) error {
	if journal {
		if err := schema.journalist.CreateField(tableName, field); err != nil {
			return err
		}
	}

	if field.Name == "x_id" {
		return nil
	}

	return schema.crud.CreateField(tableName, field)
}

func (schema *Schema) DeleteField(tableName string, fieldname string) error {
	if fieldname == "x_id" {
		return logger.Error("field x_id cannot be deleted", "Schema.DeleteField()")
	}

	if err := schema.journalist.DeleteField(tableName, fieldname); err != nil {
		return err
	}

	return schema.crud.DeleteField(tableName, fieldname)
}
