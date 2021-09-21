package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

func NewSchema() *flux.Schema {
	crud := NewCRUD()
	journalist := flux.NewJournalist( crud )

	return flux.NewSchema( journalist, crud )
}
