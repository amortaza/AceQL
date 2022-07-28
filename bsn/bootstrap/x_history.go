package bootstrap

import (
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_History() *dbschema.Table {
	table := dbschema.NewTable("x_history")

	table.SetLabel("History")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_table", "Table", dbschema.String)
	table.AddField("x_field", "Field", dbschema.String)
	table.AddField("x_record_id", "Target Record ID", dbschema.String)
	table.AddField("x_old", "Old Value", dbschema.String)
	table.AddField("x_new", "New Value", dbschema.String)
	table.AddField("x_created", "Created On", dbschema.DateTime)

	return table
}
