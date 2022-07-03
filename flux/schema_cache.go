package flux

import (
	"fmt"
	"github.com/amortaza/aceql/flux/schema_journalist"
	"github.com/amortaza/aceql/flux/tableschema"
	"github.com/amortaza/aceql/logger"
)

//var g_relation_cache = make( map[ string ] *table.Table )

// GetTableSchema will return nil on error
func GetTableSchema(name string, crud CRUD) *tableschema.Table {
	//todo
	//relation, ok := g_relation_cache[ name ]
	//
	//if ok {
	//	return relation
	//}

	if name == "x_schema" {
		return schema_journalist.Get_X_SCHEMA_schema()
		//g_relation_cache[ name ] = schema_journalist.Get_X_SCHEMA_schema()
		//return g_relation_cache[ name ]
	}

	table := tableschema.NewTable(name)

	x_schema := GetTableSchema("x_schema", crud)

	r := NewRecord(x_schema, crud)
	if r == nil {
		return nil
	}

	r.AddEq("x_table", name)
	r.AddEq("x_type", "field")
	_, err := r.Query()

	if err != nil {
		logger.Err(err, logger.SQL)
		return nil
	}

	for {
		ok, err := r.Next()
		if err != nil {
			logger.Error(err.Error(), logger.SQL)
			return nil
		}

		if !ok {
			break
		}

		if err := addField(r, table); err != nil {
			logger.Error(err.Error(), logger.SQL)
			return nil
		}
	}

	//g_relation_cache[ name ] = relation

	return table
}

func addField(r *Record, relation *tableschema.Table) error {
	fieldtype, err := r.Get("x_field_type")
	//fmt.Println( "ft " , fieldtype )
	if err != nil {
		return err
	}

	if fieldtype == string(tableschema.String) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", tableschema.String)

		return nil

	} else if fieldtype == string(tableschema.Number) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", tableschema.Number)

		return nil

	} else if fieldtype == string(tableschema.Bool) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField(field, "TODO", tableschema.Bool)

		return nil
	}

	return fmt.Errorf("unrecognized fieldtype \"%s\" in bsn/schema/schema_cache.go", fieldtype)
}
