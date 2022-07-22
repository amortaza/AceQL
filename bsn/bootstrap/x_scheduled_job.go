package bootstrap

import (
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_ScheduledJob() *dbschema.Table {
	table := dbschema.NewTable("x_scheduled_job")

	table.SetLabel("Scheduled Jobs")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_starting_datetime", "Starting Datetime", dbschema.String)
	table.AddField("x_active", "Active", dbschema.Bool)
	table.AddField("x_script_name", "Script Name", dbschema.String)
	table.AddField("x_last_run", "Last Run", dbschema.DateTime)
	table.AddField("x_seconds", "Cadence ", dbschema.Number)

	return table
}
