package node

import (
	"fmt"
)

type Bool struct {
	Text string

	nodeCompiler Compiler
}

func NewBool(value bool, nodeCompiler Compiler) Node {

	return &Bool{
		Text:         fmt.Sprintf("%t", value),
		nodeCompiler: nodeCompiler,
	}
}

func (b *Bool) Compile() (string, error) {
	return b.Text, nil
}

func (b *Bool) Put(kid Node) error {
	return fmt.Errorf("no capacity to Put() a node inside a BOOL node")
}
