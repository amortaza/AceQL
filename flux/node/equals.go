package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type Equals struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewEquals(nodeCompiler Compiler) *Equals {
	return &Equals{nodeCompiler: nodeCompiler}
}

func NewNotEquals(nodeCompiler Compiler) *Equals {
	return &Equals{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (equal *Equals) Compile() (string, error) {
	return equal.nodeCompiler.EqualCompile(equal)
}

func (equal *Equals) Put(kid Node) error {
	if equal.Left == nil {
		equal.Left = kid
		return nil
	}

	if equal.Right == nil {
		equal.Right = kid
		return nil
	}

	err := fmt.Errorf("no capacity to Put() a node inside an EQUALS node")
	return logger.Err(err, "???")
}
