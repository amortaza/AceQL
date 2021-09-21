package bootstrap

import (
	"github.com/amortaza/aceql/flux/relation_type"
)

func makeSpecificationFor_Schema() *relation_type.Relation {
	relation := relation_type.NewRelation( "x_schema")

	stringType, _ := relation_type.GetFieldTypeByName( "String")

	relation.AddField( "x_type", stringType )
	relation.AddField( "x_table", stringType )
	relation.AddField( "x_field", stringType )
	relation.AddField( "x_field_type", stringType )

	return relation
}

