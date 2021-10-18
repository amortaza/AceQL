package parser

import "strings"

type (
	StringStack struct {
		top *StackNode
		length int
	}

	StackNode struct {
		text string
		prev *StackNode
	}
)

func newStringStack() *StringStack {
	return &StringStack{nil,0}
}

func (stack *StringStack) Pop() string {
	if stack.length == 0 {
		panic("empty stack in query parser cannot be popped")
	}

	n := stack.top
	stack.top = n.prev
	stack.length--

	return n.text
}

func (stack *StringStack) Push(text string) {
	n := &StackNode{text, stack.top}
	stack.top = n
	stack.length++
}

func (stack *StringStack) Reduce() {
	if stack.length < 1 {
		return
	}

	topText := stack.top.text

	if topText == "=" {
		return

	} else if strings.Index(topText, "\"") > -1 {
		stack.reduce_String()

	} else {
		panic( "reduce unrecognized type, see " + topText)
	}
}

func (stack *StringStack) reduce_String() {
	if stack.length < 3 {
		return
	}

	stringCell := stack.top
	stringText := stringCell.text

	opCell := stringCell.prev
	opText := stringCell.text

	colCell := opCell.prev
	colText := stringCell.text

	if opText == "=" {

	} else {
		panic("reduce_string unrecognized op_text see, " + opText)
	}

	if colText == "" {

	}
}