package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_ChoiceList() *dbschema.Table {
	table := dbschema.NewTable("x_choice_list")

	table.SetLabel("Choice Lists")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_table", "Table", dbschema.String)
	table.AddField("x_field", "Field", dbschema.String)
	table.AddField("x_name", "Name", dbschema.String)
	table.AddField("x_value", "Value", dbschema.String)
	table.AddField("x_order", "Order", dbschema.Number)
	table.AddField("x_active", "Active", dbschema.Bool)

	return table
}

func makeRecordsFor_ChoiceList() ([]*flux.Record, error) {
	var records []*flux.Record

	rec, err := stdsql.NewRecord("x_choice_list")
	if err != nil {
		return nil, err
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Table")
	rec.Set("x_value", "relation")
	rec.Set("x_order", "1")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec, err = stdsql.NewRecord("x_choice_list")
	if err != nil {
		return nil, err
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "String")
	rec.Set("x_value", "string")
	rec.Set("x_order", "2")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec, err = stdsql.NewRecord("x_choice_list")
	if err != nil {
		return nil, err
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Number")
	rec.Set("x_value", "number")
	rec.Set("x_order", "3")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec, err = stdsql.NewRecord("x_choice_list")
	if err != nil {
		return nil, err
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "True/False")
	rec.Set("x_value", "bool")
	rec.Set("x_order", "4")
	rec.Set("x_active", "true")
	records = append(records, rec)

	return records, nil
}
