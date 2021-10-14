package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relations"
)

func Get_X_SCHEMA_relation() *relations.Relation {
	relation := relations.NewRelation( "x_schema")

	stringType, _ := relations.GetFieldTypeByName( "String")

	relation.AddField( "x_id", stringType )
	relation.AddField( "x_type", stringType )
	relation.AddField( "x_table", stringType )
	relation.AddField( "x_field", stringType )
	relation.AddField( "x_field_type", stringType )

	return relation
}

