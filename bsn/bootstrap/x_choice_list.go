package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/relation_type"
)

func makeSpecificationFor_ChoiceList() *relation_type.Relation {
	relation := relation_type.NewRelation( "x_choice_list") 

	stringType, _ := relation_type.GetFieldTypeByName( "String")
	numberType, _ := relation_type.GetFieldTypeByName( "Number")
	boolType, _ := relation_type.GetFieldTypeByName( "Bool")

	relation.AddField( "x_table", stringType )
	relation.AddField( "x_field", stringType )
	relation.AddField( "x_name", stringType )
	relation.AddField( "x_value", stringType )
	relation.AddField( "x_order", numberType )
	relation.AddField( "x_enabled", boolType )

	return relation
}

func makeRecordsFor_ChoiceList() []*flux.Record {
	var records []*flux.Record

	rec := stdsql.NewRecord("x_choice_list")
	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Relation")
	rec.Set("x_value", "relation")
	rec.Set("x_order", "1")
	rec.Set("x_enabled", 1)
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "String")
	rec.Set("x_value", "string")
	rec.Set("x_order", "2")
	rec.Set("x_enabled", 1)
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "Number")
	rec.Set("x_value", "number")
	rec.Set("x_order", "3")
	rec.Set("x_enabled", 1)
	records = append(records, rec)

	rec = stdsql.NewRecord("x_choice_list")
	rec.Set("x_table", "x_schema")
	rec.Set("x_field", "x_type")
	rec.Set("x_name", "True/False")
	rec.Set("x_value", "bool")
	rec.Set("x_order", "4")
	rec.Set("x_enabled", 1)
	records = append(records, rec)

	return records
}
