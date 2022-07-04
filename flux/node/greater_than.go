package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type GreaterThan struct {
	Left     Node
	Right    Node
	OrEquals bool

	nodeCompiler Compiler
}

func NewGreaterThan(nodeCompiler Compiler) *GreaterThan {
	return &GreaterThan{
		OrEquals:     false,
		nodeCompiler: nodeCompiler,
	}
}

func NewGreaterOrEquals(nodeCompiler Compiler) *GreaterThan {
	return &GreaterThan{
		OrEquals:     true,
		nodeCompiler: nodeCompiler,
	}
}

func (greaterThan *GreaterThan) Compile() (string, error) {
	return greaterThan.nodeCompiler.GreaterThanCompile(greaterThan)
}

func (greaterThan *GreaterThan) Put(kid Node) error {
	if greaterThan.Left == nil {
		greaterThan.Left = kid
		return nil
	}

	if greaterThan.Right == nil {
		greaterThan.Right = kid
		return nil
	}

	err := fmt.Errorf("no capacity to Put() a node inside a GREATHER THAN node")
	return logger.Err(err, "???")
}
