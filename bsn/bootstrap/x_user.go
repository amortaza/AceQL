package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
)

func makeSpecificationFor_User() *dbschema.Table {
	table := dbschema.NewTable("x_user")

	table.SetLabel("Users")

	table.AddField("x_id", "ID", dbschema.String)
	table.AddField("x_name", "Name", dbschema.String)
	table.AddField("x_active", "Active", dbschema.Bool)

	return table
}

func makeRecordsFor_User() ([]*flux.Record, error) {
	var records []*flux.Record

	rec, err := stdsql.NewRecord("x_user")
	if err != nil {
		return nil, err
	}

	rec.Set("x_name", "admin")
	rec.Set("x_active", "true")
	records = append(records, rec)

	return records, nil
}
