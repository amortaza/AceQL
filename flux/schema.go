package flux

import (
	"github.com/amortaza/aceql/flux/relation_type"
	"github.com/amortaza/aceql/flux/schema_journalist"
)

type Schema struct {
	journalist schema_journalist.Journalist
	crud       CRUD
}

func NewSchema(journalist schema_journalist.Journalist, crud CRUD) *Schema {
	return &Schema{
		journalist:   journalist,
		crud: crud,
	}
}

func (schema *Schema) CreateRelation_withFields(relation *relation_type.Relation, journal bool) error {
	if err := schema.CreateRelation_withName(relation.Name(), journal); err != nil {
		return err
	}

	for _, field := range relation.Fields() {
		if err := schema.CreateField(relation.Name(), field, journal); err != nil {
			return err
		}
	}

	return nil
}

func (schema *Schema) CreateRelation_withName(name string, journal bool) error {

	// the order here is importan because if we are creating 'x_schema'
	// we want to create the relation first THEN journal it

	err := schema.crud.CreateRelation(name)
	if err != nil {
		return err
	}

	if !journal {
		return nil
	}

	return schema.journalist.CreateRelation(name)
}

func (schema *Schema) DeleteRelation(name string) error {

	_ = schema.journalist.DeleteRelation(name)

	return schema.crud.DeleteRelation(name)
}

func (schema *Schema) CreateField(relationName string, field *relation_type.Field, journal bool) error {

	if journal {
		_ = schema.journalist.CreateField(relationName, field)
	}

	return schema.crud.CreateField(relationName, field)
}

func (schema *Schema) DeleteField(relationName string, fieldname string) error {

	_ = schema.journalist.DeleteField(relationName, fieldname)

	return schema.crud.DeleteField(relationName, fieldname)
}
