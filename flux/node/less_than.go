package node

import (
	"fmt"
)

type LessThan struct {
	Left    Node
	Right    Node
	OrEquals bool

	nodeCompiler Compiler
}

func NewLessThan(nodeCompiler Compiler) *LessThan {
	return &LessThan{
		OrEquals:     false,
		nodeCompiler: nodeCompiler,
	}
}

func NewLessOrEquals(nodeCompiler Compiler) *LessThan {
	return &LessThan{
		OrEquals:     true,
		nodeCompiler: nodeCompiler,
	}
}

func (lessThan *LessThan) Compile() (string, error) {
	return lessThan.nodeCompiler.LessThanCompile(lessThan)
}

func (lessThan *LessThan) Put(kid Node) error {

	if lessThan.Left == nil {
		lessThan.Left = kid
		return nil
	}

	if lessThan.Right == nil {
		lessThan.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside a LESS THAN node")
}
