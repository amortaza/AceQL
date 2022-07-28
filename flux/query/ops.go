package query

import (
	"fmt"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/logger"
	"strings"
)

type OpType string

const (
	Equals          OpType = "Equals"
	NotEquals       OpType = "NotEquals"
	LessThan        OpType = "LessThan"
	LessOrEquals    OpType = "LessOrEquals"
	GreaterThan     OpType = "GreaterThan"
	GreaterOrEquals OpType = "GreaterOrEquals"

	StartsWith    OpType = "StartsWith"
	NotStartsWith OpType = "NotStartsWith"
	EndsWith      OpType = "EndsWith"
	NotEndsWith   OpType = "NotEndsWith"
	Contains      OpType = "Contains"
	NotContains   OpType = "NotContains"
)

func GetOpTypeByName(name string) (OpType, error) {
	name = strings.ToLower(name)

	if name == "equals" {
		return Equals, nil
	}

	if name == "notequals" {
		return NotEquals, nil
	}

	if name == "lessthan" {
		return LessThan, nil
	}

	if name == "lessorequals" {
		return LessOrEquals, nil
	}

	if name == "greaterthan" {
		return GreaterThan, nil
	}

	if name == "greaterorequals" {
		return GreaterOrEquals, nil
	}

	if name == "startswith" {
		return StartsWith, nil
	}

	if name == "notstartswith" {
		return NotStartsWith, nil
	}

	if name == "endswith" {
		return EndsWith, nil
	}

	if name == "notendswith" {
		return NotEndsWith, nil
	}

	if name == "contains" {
		return Contains, nil
	}

	if name == "notcontains" {
		return NotContains, nil
	}

	return "", logger.Error(fmt.Sprintf("unrecognized OpType \"%s\"", name), "query.GetOpTypeByName()")
}

// IsEncodedOps : Ops is Equals, EncodedOps is =
func IsEncodedOps(s string) bool {
	s = strings.ToLower(s)

	return s == "=" ||
		s == "!=" ||
		s == "<" ||
		s == "<=" ||
		s == ">" ||
		s == ">=" ||
		s == strings.ToLower(string(Equals)) ||
		s == strings.ToLower(string(NotEquals)) ||
		s == strings.ToLower(string(LessThan)) ||
		s == strings.ToLower(string(LessOrEquals)) ||
		s == strings.ToLower(string(GreaterThan)) ||
		s == strings.ToLower(string(GreaterOrEquals)) ||
		s == strings.ToLower(string(StartsWith)) ||
		s == strings.ToLower(string(NotStartsWith)) ||
		s == strings.ToLower(string(EndsWith)) ||
		s == strings.ToLower(string(NotEndsWith)) ||
		s == strings.ToLower(string(Contains)) ||
		s == strings.ToLower(string(NotContains))
}

func EncodedOpToNode(op string, compiler node.Compiler) (node.Node, error) {
	op = strings.ToLower(op)

	if op == "=" || op == "equals" {
		return node.NewEquals(compiler), nil
	} else if op == "!=" || op == "notequals" {
		return node.NewNotEquals(compiler), nil
	} else if op == "<" || op == "lessthan" {
		return node.NewLessThan(compiler), nil
	} else if op == "<=" || op == "lessorequals" {
		return node.NewLessOrEquals(compiler), nil
	} else if op == ">" || op == "greaterthan" {
		return node.NewGreaterThan(compiler), nil
	} else if op == ">=" || op == "greaterorequals" {
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

	err := logger.Error("EncodedOpToNode() does not recognize ---"+op+"---", "???")
	return nil, err
}
