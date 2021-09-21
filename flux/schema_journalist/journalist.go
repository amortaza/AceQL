package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relation_type"
)

type Journalist interface {
	CreateRelation(relationName string) error
	DeleteRelation(relationName string) error

	CreateField(relationName string, field *relation_type.Field) error
	DeleteField(relationName string, fieldname string) error
}

