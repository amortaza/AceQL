package query

import (
	"errors"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/logger"
	"strconv"
	"strings"
)

func Parse(encodedQuery string, compiler node.Compiler) (node.Node, error) {
	tokens, err := tokenize(encodedQuery)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, nil
	}

	stack := newLRStack()

	stack.Push("", nil, "", nil, "")

	for _, token := range tokens {
		lctoken := strings.ToLower(token)

		if token == "(" {
			stack.Push("", nil, "", nil, "")

		} else if lctoken == "and" || lctoken == "or" {
			N := stack.Pop()
			stack.Push(token, N, "", nil, "")
			stack.Push("", nil, "", nil, "")

		} else if IsEncodedOps(token) {
			if stack.top.ops == "" {
				stack.top.ops = token
			}
		} else if token == ")" {
			for stack.top.prev != nil {
				if stack.top.prev.IsEmpty() {
					stack.top.prev = stack.top.prev.prev
					break
				}

				newTop := stack.top.prev
				newTop.rightLRNode = stack.top
				stack.top = newTop
				stack.size--
			}
		} else if stack.top.IsLeftEmpty() {
			stack.top.SetLeftText(token)

		} else if stack.top.IsRightEmpty() {
			stack.top.SetRightText(token)
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
		err := logger.Error("expected new top to be empty when collapsing", "???")
		return nil, err
	}

	newTop.rightLRNode = top

	return collapse(newTop)
}

func lrNodeToNode(lrnode *LRNode, compiler node.Compiler) (node.Node, error) {
	if lrnode.ops == "" {
		if lrnode.HasLeft() {
			if lrnode.leftLRNode == nil {
				err := logger.Error("when ops is empty and there is a left node, it MUST be a node", "???")
				return nil, err
			}
			return lrNodeToNode(lrnode.leftLRNode, compiler)
		}
		return nil, nil
	}

	if lrnode.IsLeftEmpty() {
		err := logger.Error("(lrNodeToNode 1) left of LRNode cannot be empty when ops is not empty", "???")
		return nil, err
	}

	if lrnode.IsRightEmpty() {
		// if right is empty, then left MUST be a node - it cannot be a text because "a" is not
		// an expression requires an ops - like "a = 5"
		if lrnode.leftLRNode == nil {
			err := logger.Error("(lrNodeToNode 2) left of LRNode MUST be a node, if right is empty", "???")
			return nil, err
		}
		return lrNodeToNode(lrnode.leftLRNode, compiler)
	}

	parent, err := EncodedOpToNode(lrnode.ops, compiler)
	if err != nil {
		return nil, err
	}

	if lrnode.leftText != "" {
		parent.Put(node.NewColumn(lrnode.leftText, compiler))
	} else {
		kid, err := lrNodeToNode(lrnode.leftLRNode, compiler)
		if err != nil {
			return nil, err
		}
		parent.Put(kid)
	}

	if lrnode.rightText != "" {
		if strings.Index(lrnode.rightText, "'") == 0 || strings.Index(lrnode.rightText, "\"") == 0 {
			unquoted := lrnode.rightText[1 : len(lrnode.rightText)-1]
			parent.Put(node.NewString(unquoted, compiler))

		} else if lrnode.rightText == "true" || lrnode.rightText == "false" {
			parent.Put(node.NewBool(lrnode.rightText == "true", compiler))

		} else {
			numberValue, err := strconv.ParseFloat(lrnode.rightText, 32)
			if err == nil {
				parent.Put(node.NewNumber(float32(numberValue), compiler))
			} else {
				err := logger.Error("Dont know how to handle \""+lrnode.rightText+"\" in parser", "???")
				return nil, err
			}
		}
	} else {
		kid, err := lrNodeToNode(lrnode.rightLRNode, compiler)
		if err != nil {
			return nil, err
		}
		parent.Put(kid)
	}

	return parent, nil
}
