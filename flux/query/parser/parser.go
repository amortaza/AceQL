package parser

import (
	"errors"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/query"
)

func Parse( encodedQuery string ) (node.Node, error){
	stack := newLRStack()

	tokens, err := tokenize(encodedQuery)
	if err != nil {
		return nil, err
	}

	for _, token := range tokens {
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
		} else if query.IsEncodedOps( token ) {
			if stack.top.ops == "" {
				stack.top.ops = token
			} else {
				N := stack.Pop()
				stack.Push( token, N, "", nil, "")
			}
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

	return nil, nil
}

