package node

import (
	"fmt"
)

type EndsWith struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewEndsWith(nodeCompiler Compiler) *EndsWith {
	return &EndsWith{
		Not:          false,
		nodeCompiler: nodeCompiler,
	}
}

func NewNotEndsWith(nodeCompiler Compiler) *EndsWith {
	return &EndsWith{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (endsWith *EndsWith) Compile() (string, error) {
	return endsWith.nodeCompiler.EndsWithCompile(endsWith)
}

func (endsWith *EndsWith) Put(kid Node) error {

	if endsWith.Left == nil {
		endsWith.Left = kid
		return nil
	}

	if endsWith.Right == nil {
		endsWith.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside an ENDS WITH node")
}
