package schema_journalist

import (
	"github.com/amortaza/aceql/flux/tableschema"
)

func Get_X_SCHEMA_schema() *tableschema.Table {
	tableSchema := tableschema.NewTable("x_schema")

	stringType, err := tableschema.GetFieldTypeByName("String")
	if err != nil {
		return nil
	}

	tableSchema.AddField("x_id", "ID", stringType)
	tableSchema.AddField("x_type", "Type", stringType)
	tableSchema.AddField("x_table", "Table", stringType)
	tableSchema.AddField("x_field", "Field", stringType)
	tableSchema.AddField("x_field_type", "Field Type", stringType)
	tableSchema.AddField("x_label", "Label", stringType)

	return tableSchema
}
