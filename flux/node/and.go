package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type And struct {
	Left  Node
	Right Node

	nodeCompiler Compiler
}

func NewAnd(nodeCompiler Compiler) *And {
	return &And{nodeCompiler: nodeCompiler}
}

func (and *And) Compile() (string, error) {
	return and.nodeCompiler.AndCompile(and)
}

func (and *And) Put(kid Node) error {
	if and.Left == nil {
		and.Left = kid
		return nil
	}

	if and.Right == nil {
		and.Right = kid
		return nil
	}

	err := fmt.Errorf("no capacity to Put() a node inside an AND node")

	return logger.Err(err, "???")
}
