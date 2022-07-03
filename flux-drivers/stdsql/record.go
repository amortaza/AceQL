package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

func NewRecord(tableName string) *flux.Record {
	crud := NewCRUD()

	tableSchema := flux.GetTableSchema(tableName, crud)

	return flux.NewRecord(tableSchema, crud)
}
