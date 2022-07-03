package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

// NewSchema never returns nil
func NewSchema() *flux.Schema {
	crud := NewCRUD()
	journalist := flux.NewJournalist(crud)

	return flux.NewSchema(journalist, crud)
}
