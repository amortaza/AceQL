package schema_journalist

import (
	"github.com/amortaza/aceql/flux/table"
)

func Get_X_SCHEMA_relation() *table.Relation {
	relation := table.NewRelation("x_schema")

	stringType, _ := table.GetFieldTypeByName("String")

	relation.AddField("x_id", "ID", stringType)
	relation.AddField("x_type", "Type", stringType)
	relation.AddField("x_table", "Table", stringType)
	relation.AddField("x_field", "Field", stringType)
	relation.AddField("x_field_type", "Field Type", stringType)
	relation.AddField("x_label", "Label", stringType)

	return relation
}
