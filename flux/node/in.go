package node

import (
	"fmt"
)

type In struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewIn(nodeCompiler Compiler) *In {
	return &In{nodeCompiler: nodeCompiler}
}

func NewNotIn(nodeCompiler Compiler) *In {
	return &In{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (inNode *In) Compile() (string, error) {
	return inNode.nodeCompiler.InCompile(inNode)
}

func (inNode *In) Put(kid Node) error {

	if inNode.Left == nil {
		inNode.Left = kid
		return nil
	}

	if inNode.Right == nil {
		inNode.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside an IN node")
}
