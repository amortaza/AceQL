package dbschema

func Get_X_SCHEMA_schema() *Table {
	tableschema := NewTable("x_schema")

	stringType, err := GetFieldTypeByName("String")
	if err != nil {
		return nil
	}

	tableschema.AddField("x_id", "ID", stringType)
	tableschema.AddField("x_type", "Type", stringType)
	tableschema.AddField("x_table", "Table", stringType)
	tableschema.AddField("x_field", "Field", stringType)
	tableschema.AddField("x_field_type", "Field Type", stringType)
	tableschema.AddField("x_label", "Label", stringType)

	return tableschema
}
