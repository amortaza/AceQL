package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type Column struct {
	Name string

	nodeCompiler Compiler
}

func NewColumn(name string, nodeCompiler Compiler) *Column {
	return &Column{
		Name:         name,
		nodeCompiler: nodeCompiler,
	}
}

func (column *Column) Compile() (string, error) {
	return column.nodeCompiler.ColumnCompile(column)
}

func (column *Column) Put(kid Node) error {
	err := fmt.Errorf("no capacity to Put() a node inside a COLUMN node")
	return logger.Err(err, "???")
}
