package bootstrap

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/tableschema"
)

func Run() error {
	//var empty []*flux.Record

	// schema
	//if err := bootstrap(makeSpecificationFor_Schema(), empty); err != nil {
	//	return err
	//}

	// choice list
	if err := bootstrap(makeSpecificationFor_ChoiceList(), makeRecordsFor_ChoiceList()); err != nil {
		return err
	}

	return nil
}

func bootstrap(relation *tableschema.Table, records []*flux.Record) error {
	schema := stdsql.NewSchema()

	if err := schema.CreateRelation_withName(relation.Name(), relation.Label(), relation.Name() != "x_schema"); err != nil {
		return err
	}

	for _, field := range relation.Fields() {
		if err := schema.CreateField(relation.Name(), field, relation.Name() != "x_schema"); err != nil {
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
