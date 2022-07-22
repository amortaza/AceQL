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

	if err := bootstrap_BusinessRule(); err != nil {
		return err
	}

	if err := bootstrap_ChoiceList(); err != nil {
		return err
	}

	if err := bootstrap_ImportSet(); err != nil {
		return err
	}

	if err := bootstrap_ScheduledJob(); err != nil {
		return err
	}

	if err := bootstrap_User(); err != nil {
		return err
	}

	return nil
}

func bootstrap_BusinessRule() error {
	var records []*flux.Record

	if err := bootstrap(makeSpecificationFor_BusinessRule(), records); err != nil {
		return err
	}

	return nil
}

func bootstrap_ScheduledJob() error {
	var records []*flux.Record

	if err := bootstrap(makeSpecificationFor_ScheduledJob(), records); err != nil {
		return err
	}

	return nil
}

func bootstrap_ImportSet() error {
	var records []*flux.Record

	if err := bootstrap(makeSpecificationFor_ImportSet(), records); err != nil {
		return err
	}

	return nil
}

func bootstrap_ChoiceList() error {
	records, err := makeRecordsFor_ChoiceList()
	if err != nil {
		return err
	}
	if err := bootstrap(makeSpecificationFor_ChoiceList(), records); err != nil {
		return err
	}

	return nil
}

func bootstrap_User() error {
	records, err := makeRecordsFor_User()
	if err != nil {
		return err
	}
	if err := bootstrap(makeSpecificationFor_User(), records); err != nil {
		return err
	}

	return nil
}

func bootstrap(tableschema *dbschema.Table, records []*flux.Record) error {
	schema := stdsql.NewSchema()

	if err := schema.CreateTable_withName(tableschema.Name(), tableschema.Label(), tableschema.Name() != "x_schema"); err != nil {
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
