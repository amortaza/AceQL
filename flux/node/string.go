package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type String struct {
	Text string

	nodeCompiler Compiler
}

func NewString(text string, nodeCompiler Compiler) Node {
	return &String{
		Text:         text,
		nodeCompiler: nodeCompiler,
	}
}

func (stringNode *String) Compile() (string, error) {
	return stringNode.nodeCompiler.StringCompile(stringNode)
}

func (stringNode *String) Put(kid Node) error {
	err := fmt.Errorf("no capacity to Put() a node inside a STRING node")
	return logger.Err(err, "???")
}
