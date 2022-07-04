package flux

import (
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
)

type StandardJournalist struct {
	crud CRUD
}

func (journalist *StandardJournalist) CreateTable(tableName string, tableLabel string) error {
	recordmap := NewRecordMap()

	recordmap.SetFieldValue("x_type", "relation", dbschema.String)
	recordmap.SetFieldValue("x_table", tableName, dbschema.String)
	recordmap.SetFieldValue("x_label", tableLabel, dbschema.String)
	recordmap.SetFieldValue("x_field", "x_id", dbschema.String)
	recordmap.SetFieldValue("x_field_type", string(dbschema.String), dbschema.String)

	if _, err := journalist.crud.Create("x_schema", recordmap); err != nil {
		return logger.Err(err, "StandardJournalist.CreateTable")
	}

	return nil
}

func (journalist *StandardJournalist) DeleteTable(tableName string) error {
	x_schema, err := GetTableSchema("x_schema", journalist.crud)
	if err != nil {
		return err
	}

	record := NewRecord(x_schema, journalist.crud)

	if err := record.Add("x_table", query.Equals, tableName); err != nil {
		return err
	}

	if _, err := record.Query(); err != nil {
		return err
	}

	var hasNext bool
	var id string

	for {
		if hasNext, err = record.Next(); err != nil {
			return err
		}

		if !hasNext {
			break
		}

		if id, err = record.Get("x_id"); err != nil {
			return err
		}

		if err := journalist.crud.Delete("x_schema", id); err != nil {
			return err
		}
	}

	return nil
}

func (journalist *StandardJournalist) CreateField(tableName string, field *dbschema.Field) error {
	recordmap := NewRecordMap()

	recordmap.SetFieldValue("x_type", "field", dbschema.String)
	recordmap.SetFieldValue("x_table", tableName, dbschema.String)
	recordmap.SetFieldValue("x_field", field.Name, dbschema.String)
	recordmap.SetFieldValue("x_field_type", string(field.Type), dbschema.String)
	recordmap.SetFieldValue("x_label", field.Label, dbschema.String)

	if _, err := journalist.crud.Create("x_schema", recordmap); err != nil {
		return err
	}

	return nil
}

func (journalist *StandardJournalist) DeleteField(tableName string, fieldname string) error {
	x_schema, err := GetTableSchema("x_schema", journalist.crud)
	if err != nil {
		return err
	}

	record := NewRecord(x_schema, journalist.crud)

	if err := record.Add("x_table", query.Equals, tableName); err != nil {
		return err
	}

	if err := record.Add("x_field", query.Equals, fieldname); err != nil {
		return err
	}

	if _, err := record.Query(); err != nil {
		return err
	}

	var hasNext bool
	var id string

	for {
		if hasNext, err = record.Next(); err != nil {
			return err
		}

		if !hasNext {
			break
		}

		if id, err = record.Get("x_id"); err != nil {
			return err
		}

		if err := journalist.crud.Delete("x_schema", id); err != nil {
			return err
		}
	}

	return nil
}
