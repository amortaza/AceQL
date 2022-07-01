package schema_journalist

import (
	"github.com/amortaza/aceql/flux/tableschema"
)

func Get_X_SCHEMA_relation() *tableschema.Table {
	relation := tableschema.NewTable("x_schema")

	stringType, _ := tableschema.GetFieldTypeByName("String")

	relation.AddField("x_id", "ID", stringType)
	relation.AddField("x_type", "Type", stringType)
	relation.AddField("x_table", "Table", stringType)
	relation.AddField("x_field", "Field", stringType)
	relation.AddField("x_field_type", "Field Type", stringType)
	relation.AddField("x_label", "Label", stringType)

	return relation
}
