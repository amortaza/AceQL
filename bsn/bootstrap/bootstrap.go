package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
)

func Run() error {
	// schema
	// we do not need to bootstrap schema, because it is hard-coded in code!
	// todo test bootstrap

	// choice list
	if err := bootstrap(makeSpecificationFor_ChoiceList(), makeRecordsFor_ChoiceList()); err != nil {
		return err
	}

	return nil
}

func bootstrap(tableschema *dbschema.Table, records []*flux.Record) error {
	schema := stdsql.NewSchema()

	if err := schema.CreateRelation_withName(tableschema.Name(), tableschema.Label(), tableschema.Name() != "x_schema"); err != nil {
		return err
	}

	for _, field := range tableschema.Fields() {
		if err := schema.CreateField(tableschema.Name(), field, tableschema.Name() != "x_schema"); err != nil {
			return err
		}
	}

	for _, record := range records {
		if _, err := record.Insert(); err != nil {
			return err
		}
	}

	return nil
}
