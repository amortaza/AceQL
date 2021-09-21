package flux

import (
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/relation_type"
	"github.com/amortaza/aceql/flux/schema_journalist"
)

type CRUD interface {
	// Compiler is here for now for convenience, but it really doesn't belong here
	Compiler() node.Compiler

	Query(relationName string, root node.Node) error
	Next() (*RecordMap, error)

	Create(relationName string, values *RecordMap) (string, error)
	Update(relationName string, id string, values *RecordMap) error
	Delete(relationName string, id string) error

	// schema crud operations
	CreateRelation(name string) error
	DeleteRelation(name string) error
	CreateField(relationName string, field *relation_type.Field) error
	DeleteField(relationName string, fieldname string) error
}

func NewJournalist(crud CRUD) schema_journalist.Journalist {
	return &StandardJournalist{crud: crud}
}

