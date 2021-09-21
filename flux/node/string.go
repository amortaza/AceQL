package node

import "fmt"

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
	return fmt.Errorf("no capacity to Put() a node inside a STRING node")
}
