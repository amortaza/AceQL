package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

// NewSchema NEVER FAILS (0)
func NewSchema() *flux.Schema {
	crud := NewCRUD()
	journalist := flux.NewJournalist(crud)

	return flux.NewSchema(journalist, crud)
}
