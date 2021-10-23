package flux

import (
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/relations"
)

type StandardJournalist struct {
	crud CRUD
}

func (journalist *StandardJournalist) CreateRelation(relationName string) error {
	recordmap := NewRecordMap()

	recordmap.PutString("x_type", "relation" )
	recordmap.PutString("x_table", relationName)
	recordmap.PutString("x_field", "x_id")
	recordmap.PutString("x_field_type", string(relations.String))

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteRelation(relationName string) error {
	record := NewRecord(GetRelation("x_schema", journalist.crud), journalist.crud)

	_ = record.Add("x_table", query.Equals, relationName)

	_, _ = record.Query()

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

func (journalist *StandardJournalist) CreateField(relationName string, field *relations.Field) error {
	recordmap := NewRecordMap()

	recordmap.PutString("x_type", "field")
	recordmap.PutString("x_table", relationName)
	recordmap.PutString("x_field", field.Name)
	recordmap.PutString("x_field_type", string(field.Type))

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteField(relationName string, fieldname string) error {
	record := NewRecord(GetRelation("x_schema", journalist.crud), journalist.crud)

	_ = record.Add("x_table", query.Equals, relationName)
	_ = record.Add("x_field", query.Equals, fieldname)

	_, _ = record.Query()

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