package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/tableschema"
)

func makeSpecificationFor_ChoiceList() *tableschema.Table {
	table := tableschema.NewTable("x_choice_list")

	table.AddField("x_id", "ID", tableschema.String)
	table.AddField("x_table", "Table", tableschema.String)
	table.AddField("x_field", "Field", tableschema.String)
	table.AddField("x_name", "Name", tableschema.String)
	table.AddField("x_value", "Value", tableschema.String)
	table.AddField("x_order", "Order", tableschema.Number)
	table.AddField("x_active", "Active", tableschema.Bool)

	return table
}

func makeRecordsFor_ChoiceList() []*flux.Record {
	var records []*flux.Record

	rec := stdsql.NewRecord("x_choice_list")
	if rec == nil {
		return nil
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Table")
	rec.Set("x_value", "relation")
	rec.Set("x_order", "1")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	if rec == nil {
		return nil
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "String")
	rec.Set("x_value", "string")
	rec.Set("x_order", "2")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	if rec == nil {
		return nil
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Number")
	rec.Set("x_value", "number")
	rec.Set("x_order", "3")
	rec.Set("x_active", "true")
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	if rec == nil {
		return nil
	}

	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "True/False")
	rec.Set("x_value", "bool")
	rec.Set("x_order", "4")
	rec.Set("x_active", "true")
	records = append(records, rec)

	return records
}
