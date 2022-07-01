package schema_journalist

import (
	"github.com/amortaza/aceql/flux/tableschema"
)

type Journalist interface {
	CreateTable(tableName string, tableLabel string) error
	DeleteTable(tableName string) error

	CreateField(tableName string, field *tableschema.Field) error
	DeleteField(tableName string, fieldname string) error
}
