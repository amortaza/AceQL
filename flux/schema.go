package flux

import (
	"fmt"
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/schema_journalist"
	"github.com/amortaza/aceql/flux/table"
)

type Schema struct {
	journalist schema_journalist.Journalist
	crud       CRUD
}

func NewSchema(journalist schema_journalist.Journalist, crud CRUD) *Schema {
	return &Schema{
		journalist: journalist,
		crud:       crud,
	}
}

func (schema *Schema) Close() {
	schema.crud.Close()
}

func (schema *Schema) CreateRelation_withFields(relation *table.Relation, journal bool) error {
	if err := schema.CreateRelation_withName(relation.Name(), relation.Label(), journal); err != nil {
		return err
	}

	for _, field := range relation.Fields() {
		if err := schema.CreateField(relation.Name(), field, journal); err != nil {
			return err
		}
	}

	return nil
}

func (schema *Schema) CreateRelation_withName(name string, label string, journal bool) error {

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

	_ = schema.journalist.DeleteTable(name)

	return schema.crud.DeleteTable(name)
}

func (schema *Schema) CreateField(relationName string, field *table.Field, journal bool) error {

	if journal {
		_ = schema.journalist.CreateField(relationName, field)
	}

	if field.Name == "x_id" {
		return nil
	}

	return schema.crud.CreateField(relationName, field)
}

func (schema *Schema) DeleteField(relationName string, fieldname string) error {
	if fieldname == "x_id" {
		err := fmt.Errorf("field x_id cannot be deleted")
		logger.Error(err, "Schema.DeleteField()")
		return err
	}

	_ = schema.journalist.DeleteField(relationName, fieldname)

	return schema.crud.DeleteField(relationName, fieldname)
}
