package flux

import (
	"fmt"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
)

//var g_relation_cache = make( map[ string ] *table.Table )

func GetTableSchema(name string, crud CRUD) (*dbschema.Table, error) {
	//todo
	//relation, ok := g_relation_cache[ name ]
	//
	//if ok {
	//	return relation
	//}

	if name == "x_schema" {
		return dbschema.Get_X_SCHEMA_schema(), nil
		//g_relation_cache[ name ] = dbschema.Get_X_SCHEMA_schema()
		//return g_relation_cache[ name ]
	}

	table := dbschema.NewTable(name)

	// will never error
	x_schema, _ := GetTableSchema("x_schema", crud)

	r := NewRecord(x_schema, crud)

	if err := r.AddEq("x_table", name); err != nil {
		return nil, err
	}

	if err := r.AddEq("x_type", "field"); err != nil {
		return nil, err
	}

	if _, err := r.Query(); err != nil {
		return nil, err
	}

	for {
		hasNext, err := r.Next()

		if err != nil {
			return nil, logger.Err(err, "SQL")
		}

		if !hasNext {
			break
		}

		if err := addField(r, table); err != nil {
			return nil, logger.Err(err, "SQL")
		}
	}

	//g_relation_cache[ name ] = relation

	return table, nil
}

func addField(r *Record, relation *dbschema.Table) error {
	fieldtype, err := r.Get("x_field_type")

	if err != nil {
		return err
	}

	if fieldtype == string(dbschema.String) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", dbschema.String)

		return nil

	} else if fieldtype == string(dbschema.Number) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", dbschema.Number)

		return nil

	} else if fieldtype == string(dbschema.Bool) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", dbschema.Bool)

		return nil
	}

	return fmt.Errorf("unrecognized fieldtype \"%s\" in bsn/schema/schema_cache.go", fieldtype)
}
