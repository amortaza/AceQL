package query

type (
	LRStack struct {
		top  *LRNode
		size int
	}

	LRNode struct {
		leftLRNode *LRNode
		rightLRNode *LRNode
		leftText string
		rightText string
		ops string
		prev *LRNode
	}
)

func newLRStack() *LRStack {
	return &LRStack{nil,0}
}

func (stack *LRStack) Pop() *LRNode {
	if stack.size == 0 {
		return nil
	}

	n := stack.top
	stack.top = n.prev
	stack.size--

	return n
}

func (stack *LRStack) Push(ops string, leftLRNode *LRNode, leftText string, rightLRNode *LRNode, rightText string) {
	n := &LRNode{leftLRNode: leftLRNode, rightLRNode: rightLRNode, leftText: leftText, rightText: rightText, ops: ops}
	stack.PushNode(n)
}
func (stack *LRStack) PushNode(n *LRNode) {
	n.prev = stack.top
	stack.top = n
	stack.size++
}

func (stack *LRStack) IsEmpty() bool {
	return stack.size == 0
}

func (lrnode *LRNode) HasOps() bool {
	return lrnode.ops != ""
}

func (lrnode *LRNode) SetOps(ops string) {
	lrnode.ops = ops
}

func (lrnode *LRNode) IsOpsGroup() bool {
	return lrnode.ops == "and" || lrnode.ops == "or"
}

func (lrnode *LRNode) IsOpsEmpty() bool {
	return lrnode.ops == ""
}

func (lrnode *LRNode) IsLeftEmpty() bool {
	return lrnode.leftText == "" && lrnode.leftLRNode == nil
}

func (lrnode *LRNode) HasLeft() bool {
	return lrnode.leftText != "" || lrnode.leftLRNode != nil
}

func (lrnode *LRNode) IsRightEmpty() bool {
	return lrnode.rightText == "" && lrnode.rightLRNode == nil
}

func (lrnode *LRNode) HasRight() bool {
	return lrnode.rightText != "" || lrnode.rightLRNode != nil
}

func (lrnode *LRNode) SetLeftText(text string) {
	lrnode.leftText = text
}

func (lrnode *LRNode) SetRightText(text string) {
	lrnode.rightText = text
}
