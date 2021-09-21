package node

import (
	"fmt"
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
	return fmt.Errorf("no capacity to Put() a node inside a COLUMN node")
}
