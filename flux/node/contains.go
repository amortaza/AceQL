package node

import (
	"fmt"
)

type Contains struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewContains(nodeCompiler Compiler) *Contains {
	return &Contains{
		Not:          false,
		nodeCompiler: nodeCompiler,
	}
}

func NewNotContains(nodeCompiler Compiler) *Contains {
	return &Contains{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (contains *Contains) Compile() (string, error) {
	return contains.nodeCompiler.ContainsCompile(contains)
}

func (contains *Contains) Put(kid Node) error {

	if contains.Left == nil {
		contains.Left = kid
		return nil
	}

	if contains.Right == nil {
		contains.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside a CONTAINS node")
}
