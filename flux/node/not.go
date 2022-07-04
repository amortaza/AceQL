package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type Not struct {
	Left Node

	nodeCompiler Compiler
}

func NewNot(nodeCompiler Compiler) *Not {
	return &Not{nodeCompiler: nodeCompiler}
}

func (notNode *Not) Compile() (string, error) {
	return notNode.nodeCompiler.NotCompile(notNode)
}

func (notNode *Not) Put(kid Node) error {
	if notNode.Left == nil {
		notNode.Left = kid
		return nil
	}

	err := fmt.Errorf("no capacity to Put() a node inside a NOT node")
	return logger.Err(err, "???")
}
