package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type StartsWith struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewStartsWith(nodeCompiler Compiler) *StartsWith {
	return &StartsWith{
		Not:          false,
		nodeCompiler: nodeCompiler,
	}
}

func NewNotStartsWith(nodeCompiler Compiler) *StartsWith {
	return &StartsWith{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (startsWith *StartsWith) Compile() (string, error) {
	return startsWith.nodeCompiler.StartsWithCompile(startsWith)
}

func (startsWith *StartsWith) Put(kid Node) error {
	if startsWith.Left == nil {
		startsWith.Left = kid
		return nil
	}

	if startsWith.Right == nil {
		startsWith.Right = kid
		return nil
	}

	err := fmt.Errorf("no capacity to Put() a node inside a STARTS WITH node")
	return logger.Err(err, "???")
}
