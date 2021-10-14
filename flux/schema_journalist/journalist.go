package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relations"
)

type Journalist interface {
	CreateRelation(relationName string) error
	DeleteRelation(relationName string) error

	CreateField(relationName string, field *relations.Field) error
	DeleteField(relationName string, fieldname string) error
}

