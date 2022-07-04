package flux

import (
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/flux/node"
)

type CRUD interface {
	// Compiler is here for now for convenience, but it really doesn't belong here
	Compiler() node.Compiler

	Query(tableName string, fields []*dbschema.Field, root node.Node, paginationIndex int, paginationSize int, orderBy string, orderByAscending bool) (int, error)
	Next() (*RecordMap, error)

	Create(tableName string, values *RecordMap) (string, error)
	Update(tableName string, id string, values *RecordMap) error
	Delete(tableName string, id string) error // should log errors

	// schema crud operations
	CreateTable(name string) error
	DeleteTable(name string) error
	CreateField(tableName string, field *dbschema.Field) error
	DeleteField(tableName string, fieldname string) error

	Close() error
}

func NewJournalist(crud CRUD) dbschema.Journalist {
	return &StandardJournalist{crud: crud}
}
