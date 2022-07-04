package node

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type StringList struct {
	Strings []string

	nodeCompiler Compiler
}

func NewStringList(strings []string, nodeCompiler Compiler) Node {
	return &StringList{
		Strings:      strings,
		nodeCompiler: nodeCompiler,
	}
}

func (stringList *StringList) Compile() (string, error) {
	return stringList.nodeCompiler.StringListCompile(stringList)
}

func (stringList *StringList) Put(kid Node) error {
	err := fmt.Errorf("no capacity to Put() a node inside a LIST OF STRINGS node")
	return logger.Err(err, "???")
}
