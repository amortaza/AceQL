package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relations"
)

type Journalist interface {
	CreateTable(relationName string) error
	DeleteTable(relationName string) error

	CreateField(relationName string, field *relations.Field) error
	DeleteField(relationName string, fieldname string) error
}

