package bootstrap

import (
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_BusinessRule() *dbschema.Table {
	table := dbschema.NewTable("x_business_rule")

	table.SetLabel("Business Rules")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_table", "Table", dbschema.String)
	table.AddField("x_script_name", "Script Name", dbschema.String)
	table.AddField("x_active", "Active", dbschema.Bool)

	return table
}
