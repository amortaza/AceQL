package flux

import (
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/schema_journalist"
	"github.com/amortaza/aceql/flux/tableschema"
)

type CRUD interface {
	// Compiler is here for now for convenience, but it really doesn't belong here
	Compiler() node.Compiler

	Query(tableName string, fields []*tableschema.Field, root node.Node, paginationIndex int, paginationSize int, orderBy string, orderByAscending bool) (int, error)
	Next() (*RecordMap, error)

	Create(tableName string, values *RecordMap) (string, error)
	Update(tableName string, id string, values *RecordMap) error
	Delete(tableName string, id string) error

	// schema crud operations
	CreateTable(name string) error
	DeleteTable(name string) error
	CreateField(tableName string, field *tableschema.Field) error
	DeleteField(tableName string, fieldname string) error

	Close() error
}

func NewJournalist(crud CRUD) schema_journalist.Journalist {
	return &StandardJournalist{crud: crud}
}
