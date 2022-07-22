package bootstrap

import (
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_ImportSet() *dbschema.Table {
	table := dbschema.NewTable("x_importset")

	table.SetLabel("Import Sets")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_adapter", "Adapter", dbschema.String)
	table.AddField("x_target_table", "Target Table", dbschema.String)
	table.AddField("x_name", "Name", dbschema.String)
	table.AddField("x_mappings", "Mappings", dbschema.String)

	return table
}
