package query

import (
	"errors"
	"github.com/amortaza/aceql/flux/node"
	"strings"
)

type OpType string

const (
	Equals         OpType = "Equals"
	NotEquals      OpType = "NotEquals"
	LessThan       OpType = "LessThan"
	LessOrEqual    OpType = "LessOrEqual"
	GreaterThan    OpType = "GreaterThan"
	GreaterOrEqual OpType = "GreaterOrEqual"

	StartsWith    OpType = "StartsWith"
	NotStartsWith OpType = "NotStartsWith"
	EndsWith      OpType = "EndsWith"
	NotEndsWith   OpType = "NotEndsWith"
	Contains      OpType = "Contains"
	NotContains   OpType = "NotContains"
)

// IsEncodedOps : Ops is Equals, EncodedOps is =
func IsEncodedOps( s string ) bool {
	s = strings.ToLower(s)

	return s == "=" ||
		   s == "!=" ||
		s == "<"  ||
		s == "<="  ||
		s == ">"  ||
		s == ">="  ||

		s == string(StartsWith)  ||
		s == string(NotStartsWith)  ||
		s == string(EndsWith)  ||
		s == string(NotEndsWith)  ||
		s == string(Contains)  ||
		s == string(NotContains)
}

func EncodedOpToNode(op string, compiler node.Compiler) (node.Node, error) {
	op = strings.ToLower(op)

	if op == "=" {
		return node.NewEquals(compiler), nil
	} else if op == "!=" {
		return node.NewNotEquals(compiler), nil
	} else if op == "<" {
		return node.NewLessThan(compiler), nil
	} else if op == "<=" {
		return node.NewLessOrEquals(compiler), nil
	} else if op == ">" {
		return node.NewGreaterThan(compiler), nil
	} else if op == ">=" {
		return node.NewGreaterOrEquals(compiler), nil
	} else if op == "startswith" {
		return node.NewStartsWith(compiler), nil
	} else if op == "notstartswith" {
		return node.NewNotStartsWith(compiler), nil
	} else if op == "endswith" {
		return node.NewEndsWith(compiler), nil
	} else if op == "notendswith" {
		return node.NewNotEndsWith(compiler), nil
	} else if op == "contains" {
		return node.NewContains(compiler), nil
	} else if op == "notcontains" {
		return node.NewNotContains(compiler), nil
	} else if op == "and" {
		return node.NewAnd(compiler), nil
	} else if op == "or" {
		return node.NewOr(compiler), nil
	}

	return nil, errors.New("EncodedOpToNode() does not recognize ---" + op + "---")
}