package query

import (
	"errors"
	"fmt"
	"github.com/amortaza/aceql/flux/node"
	"strings"
)

func Parse( encodedQuery string, compiler node.Compiler ) (node.Node, error) {
	tokens, err := tokenize(encodedQuery)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, nil
	}

	stack := newLRStack()

	for _, token := range tokens {
		lctoken := strings.ToLower(token)

		if token == "(" {
			stack.Push("", nil, "", nil, "")

		} else if lctoken == "and" || lctoken == "or" {
			N := stack.Pop()
			stack.Push(token, N, "", nil, "")

		} else if IsEncodedOps( token ) {
			if stack.top.ops == "" {
				stack.top.ops = token
			} else {
				N := stack.Pop()
				stack.Push( token, N, "", nil, "")
			}
		} else if token == ")" {
			for stack.top.prev != nil {
				if stack.top.prev.HasRight() {
					break
				}

				newTop := stack.top.prev
				newTop.rightLRNode = stack.top
				stack.top = newTop
				stack.size--

				if stack.top.IsOpsEmpty() {
					if stack.top.HasLeft() {
						return nil, errors.New("(2) Parse() encoded query is malformed, see ---" + encodedQuery + "---")
					} else {
						stack.top.leftText = "true"
						stack.top.ops = "AND"
					}
				}
			}
		} else {
			if stack.IsEmpty() {
				stack.Push("", nil, token, nil, "")

			} else {
				if stack.top.IsOpsGroup() {
					if !stack.top.IsRightEmpty() {
						fmt.Println("I dont know what to do (1)") // debug
						return nil, errors.New("(1) Parse() encoded query is malformed, see ---" + encodedQuery + "---")
					}

					stack.Push("", nil, token, nil, "")

				} else if stack.top.IsLeftEmpty() {
					stack.top.SetLeftText( token )

				} else if stack.top.IsRightEmpty() {
					stack.top.SetRightText( token )

				} else {
					fmt.Println("I dont know what to do ") // debug
					return nil, errors.New("(2) Parse() encoded query is malformed, see ---" + encodedQuery + "---")
				}
			}
		}
	}

	lrnode, err := collapse(stack.top)
	if err != nil {
		return nil, err
	}

	return lrNodeToNode(lrnode, compiler)
}

func Parse2( encodedQuery string, compiler node.Compiler ) (node.Node, error) {
	tokens, err := tokenize(encodedQuery)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, nil
	}

	stack := newLRStack()

	for _, token := range tokens {
		lctoken := strings.ToLower(token)

		if token == "(" {
			stack.Push("", nil, "", nil, "")

		} else if token == ")" {
			if stack.IsEmpty() {
				return nil, errors.New("encoded query is malformed (1), see ---" + encodedQuery + "---")
			}

			N := stack.Pop()
			if stack.IsEmpty() {
				stack.PushNode(N)
			} else if stack.top.IsLeftEmpty() {
				stack.top.leftLRNode = N
			} else if stack.top.IsRightEmpty() {
				stack.top.rightLRNode = N
			} else {
				return nil, errors.New("encoded query is malformed (2), see ---" + encodedQuery + "---")
			}
		} else if IsEncodedOps( token ) {
			if stack.top.ops == "" {
				stack.top.ops = token
			} else {
				N := stack.Pop()
				stack.Push( token, N, "", nil, "")
			}
		} else if lctoken == "and" || lctoken == "or" {
			N := stack.Pop()
			stack.Push( token, N, "", nil, "")
			stack.Push( "", nil, "", nil, "")
		} else {
			if stack.IsEmpty() {
				stack.Push( "", nil, token, nil, "")
			} else {
				if stack.top.IsLeftEmpty() {
					stack.top.SetLeftText( token )
				} else if stack.top.IsRightEmpty() {
					stack.top.SetRightText( token )
				} else {
					return nil, errors.New("encoded query is malformed (3), see ---" + encodedQuery + "---")
				}
			}
		}
	}

	lrnode, err := collapse(stack.top)
	if err != nil {
		return nil, err
	}

	return lrNodeToNode(lrnode, compiler)
}

func collapse(top *LRNode) (*LRNode, error) {
	if top.prev == nil {
		if top.HasOps() && top.HasLeft() && top.HasRight() {
			return top, nil

		} else if top.HasOps() {
			return nil, errors.New("when there is ops, cannot have left or right empty when collapsing")

		} else {
			// no ops
			if top.HasLeft() && top.HasRight() {
				return nil, errors.New("when there is no ops, cannot have both left and right when collapsing")
			} else if top.HasLeft() {
				if top.leftLRNode == nil {
					return nil, errors.New("must be left NODE when collapsing")
				}
				return top.leftLRNode, nil

			} else if top.HasRight() {
				if top.rightLRNode == nil {
					return nil, errors.New("must be right NODE when collapsing")
				}
				return top.rightLRNode, nil
			} else {
				return nil, errors.New("not possible")
			}
		}
	}

	newTop := top.prev

	if newTop.IsLeftEmpty() && newTop.IsRightEmpty() && !newTop.HasOps() {
		top.prev = newTop.prev
		return collapse(top)
	}

	if newTop.HasRight() {
		return nil, errors.New("expected new top to be empty when collapsing")
	}

	newTop.rightLRNode = top

	return collapse(newTop)
}

func lrNodeToNode(lrnode *LRNode, compiler node.Compiler) (node.Node, error) {
	if lrnode.ops == "" {
		if lrnode.HasLeft() {
			if lrnode.leftLRNode == nil {
				return nil, errors.New("when ops is empty and there is a left node, it MUST be a node")
			}
			return lrNodeToNode(lrnode.leftLRNode, compiler)
		}
		return nil, nil
	}

	if lrnode.IsLeftEmpty() {
		return nil, errors.New("(lrNodeToNode 1) left of LRNode cannot be empty when ops is not empty")
	}

	if lrnode.IsRightEmpty() {
		// if right is empty, then left MUST be a node - it cannot be a text because "a" is not
		// an expression requires an ops - like "a = 5"
		if lrnode.leftLRNode == nil {
			return nil, errors.New("(lrNodeToNode 2) left of LRNode MUST be a node, if right is empty")
		}
		return lrNodeToNode( lrnode.leftLRNode, compiler )
	}

	parent, err := EncodedOpToNode(lrnode.ops, compiler)
	if err != nil {
		return nil, err
	}

	if lrnode.leftText != "" {
		parent.Put( node.NewColumn( lrnode.leftText, compiler ) )
	} else {
		kid, err := lrNodeToNode( lrnode.leftLRNode, compiler )
		if err != nil {
			return nil, err
		}
		parent.Put( kid )
	}

	if lrnode.rightText != "" {
		parent.Put( node.NewString( lrnode.rightText, compiler ) )
	} else {
		kid, err := lrNodeToNode( lrnode.rightLRNode, compiler )
		if err != nil {
			return nil, err
		}
		parent.Put( kid )
	}

	return parent, nil
}

