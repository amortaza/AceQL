package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relations"
)

func Get_X_SCHEMA_relation() *relations.Relation {
	relation := relations.NewRelation( "x_schema")

	stringType, _ := relations.GetFieldTypeByName( "String")

	relation.AddField( "x_id", "ID", stringType )
	relation.AddField( "x_type", "Type", stringType )
	relation.AddField( "x_table", "Table", stringType )
	relation.AddField( "x_field", "Field", stringType )
	relation.AddField( "x_field_type", "Field Type", stringType )
	relation.AddField( "x_label", "Label", stringType )

	return relation
}

