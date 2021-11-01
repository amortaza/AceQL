package schema_journalist

import (
	"github.com/amortaza/aceql/flux/relations"
)

type Journalist interface {
	CreateTable(tableName string, tableLabel string) error
	DeleteTable(tableName string) error

	CreateField(tableName string, field *relations.Field) error
	DeleteField(tableName string, fieldname string) error
}

