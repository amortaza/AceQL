package node

import (
	"fmt"
)

type NumberList struct {
	Numbers []int

	nodeCompiler Compiler
}

func NewNumberList(numbers []int, nodeCompiler Compiler) Node {
	return &NumberList{
		Numbers:      numbers,
		nodeCompiler: nodeCompiler,
	}
}

func (numberList *NumberList) Compile() (string, error) {
	return numberList.nodeCompiler.NumberListCompile(numberList)
}

func (numberList *NumberList) Put(kid Node) error {
	return fmt.Errorf("no capacity to Put() a node inside a LIST OF INTS node")
}
