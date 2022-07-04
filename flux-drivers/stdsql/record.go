package stdsql

import (
	"github.com/amortaza/aceql/flux"
)

func NewRecord(tableName string) (*flux.Record, error) {
	crud := NewCRUD()

	tableschema, err := flux.GetTableSchema(tableName, crud)
	if err != nil {
		return nil, err
	}

	return flux.NewRecord(tableschema, crud), nil
}
