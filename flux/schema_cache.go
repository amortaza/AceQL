package flux

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/relations"
	"github.com/amortaza/aceql/flux/schema_journalist"
)

var g_relation_cache = make( map[ string ] *relations.Relation )

func GetRelation( name string, crud CRUD) *relations.Relation {
	relation, ok := g_relation_cache[ name ]

	if ok {
		return relation
	}

	if name == "x_schema" {
		g_relation_cache[ name ] = schema_journalist.Get_X_SCHEMA_relation()
		return g_relation_cache[ name ]
	}

	relation = relations.NewRelation( name )

	r := NewRecord(GetRelation("x_schema", crud), crud)
	r.Add( "x_table", query.Equals, name )
	r.Add( "x_type", query.Equals, "field" )
	_, err := r.Query()

	if err != nil {
		logger.Error(err, logger.SQL)
		return nil
	}

	for {
		ok, err := r.Next()
		if err != nil {
			logger.Error(err, logger.SQL)
			return nil
		}

		if !ok {
			break
		}

		if err := addField( r, relation ); err != nil {
			logger.Error(err, logger.SQL)
			return nil
		}
	}

	g_relation_cache[ name ] = relation

	return relation
}

func addField(r *Record, relation *relations.Relation) error {
	fieldtype, err := r.Get("x_field_type")
	//fmt.Println( "ft " , fieldtype )
	if err != nil {
		return err
	}

	if fieldtype == string( relations.String ) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField( field, relations.String )

		return nil

	} else if fieldtype == string( relations.Number ) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField( field, relations.Number )

		return nil

	} else if fieldtype == string( relations.Bool ) {
		field, err := r.Get("x_field")
		if err != nil {
			return err
		}

		relation.AddField( field, relations.Bool )

		return nil
	}

	return fmt.Errorf("unrecognized fieldtype \"%s\" in bsn/schema/schema_cache.go", fieldtype )
}
