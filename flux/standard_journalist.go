package flux

import (
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/relation_type"
)

type StandardJournalist struct {
	crud CRUD
}

func (journalist *StandardJournalist) CreateRelation(relationName string) error {
	recordmap := NewRecordMap()

	recordmap.Put("x_type", "relation")
	recordmap.Put("x_table", relationName)
	recordmap.Put("x_field", "x_id")
	recordmap.Put("x_field_type", string(relation_type.String))

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteRelation(relationName string) error {
	record := NewRecord("x_schema", journalist.crud)

	_ = record.Add("x_table", query.Equals, relationName)

	_ = record.Query()

	for {
		has, _ := record.Next()

		if !has {
			break
		}

		id, _ := record.Get("x_id")
		_ = journalist.crud.Delete("x_schema", id)
	}

	return nil
}

func (journalist *StandardJournalist) CreateField(relationName string, field *relation_type.Field) error {
	recordmap := NewRecordMap()

	recordmap.Put("x_type", "field")
	recordmap.Put("x_table", relationName)
	recordmap.Put("x_field", field.Name)
	recordmap.Put("x_field_type", string(field.Type))

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteField(relationName string, fieldname string) error {
	record := NewRecord("x_schema", journalist.crud)

	_ = record.Add("x_table", query.Equals, relationName)
	_ = record.Add("x_field", query.Equals, fieldname)

	_ = record.Query()

	for {
		has, _ := record.Next()

		if !has {
			break
		}

		id, _ := record.Get("x_id")
		_ = journalist.crud.Delete("x_schema", id)
	}

	return nil
}