package flux

import (
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/tableschema"
)

type StandardJournalist struct {
	crud CRUD
}

func (journalist *StandardJournalist) CreateTable(tableName string, tableLabel string) error {
	recordmap := NewRecordMap()

	recordmap.SetFieldValue("x_type", "relation", tableschema.String)
	recordmap.SetFieldValue("x_table", tableName, tableschema.String)
	recordmap.SetFieldValue("x_label", tableLabel, tableschema.String)
	recordmap.SetFieldValue("x_field", "x_id", tableschema.String)
	recordmap.SetFieldValue("x_field_type", string(tableschema.String), tableschema.String)

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteTable(tableName string) error {
	x_schema := GetTableSchema("x_schema", journalist.crud)

	record := NewRecord(x_schema, journalist.crud)

	_ = record.Add("x_table", query.Equals, tableName)

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

func (journalist *StandardJournalist) CreateField(tableName string, field *tableschema.Field) error {
	recordmap := NewRecordMap()

	recordmap.SetFieldValue("x_type", "field", tableschema.String)
	recordmap.SetFieldValue("x_table", tableName, tableschema.String)
	recordmap.SetFieldValue("x_field", field.Name, tableschema.String)
	recordmap.SetFieldValue("x_field_type", string(field.Type), tableschema.String)
	recordmap.SetFieldValue("x_label", field.Label, tableschema.String)

	_, err := journalist.crud.Create("x_schema", recordmap)

	return err
}

func (journalist *StandardJournalist) DeleteField(tableName string, fieldname string) error {
	x_schema := GetTableSchema("x_schema", journalist.crud)

	record := NewRecord(x_schema, journalist.crud)

	_ = record.Add("x_table", query.Equals, tableName)
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
