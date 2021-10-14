package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

func NewRecord(relationName string) *flux.Record {
	crud := NewCRUD()

	return flux.NewRecord(flux.GetRelation(relationName, crud), crud)
}
