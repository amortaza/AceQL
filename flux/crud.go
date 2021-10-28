package flux

import (
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/relations"
	"github.com/amortaza/aceql/flux/schema_journalist"
)

type CRUD interface {
	// Compiler is here for now for convenience, but it really doesn't belong here
	Compiler() node.Compiler

	Query(relationName string, fields []*relations.Field, root node.Node, paginationIndex int, paginationSize int) (int,error)
	Next() (*RecordMap, error)

	Create(relationName string, values *RecordMap) (string, error)
	Update(relationName string, id string, values *RecordMap) error
	Delete(relationName string, id string) error

	// schema crud operations
	CreateTable(name string) error
	DeleteTable(name string) error
	CreateField(relationName string, field *relations.Field) error
	DeleteField(relationName string, fieldname string) error

	Close() error
}

func NewJournalist(crud CRUD) schema_journalist.Journalist {
	return &StandardJournalist{crud: crud}
}

