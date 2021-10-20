package query

import "strings"

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

func ContainsEncodedOps( s string ) bool {
	return strings.Index(s, "=") > -1 ||
		strings.Index(s, "!=") > -1 ||
		strings.Index(s, "<") > -1 ||
		strings.Index(s, "<=") > -1 ||
		strings.Index(s, ">") > -1 ||
		strings.Index(s, ">=") > -1 ||

		strings.Index(s, string(StartsWith)) > -1 ||
		strings.Index(s, string(NotStartsWith)) > -1 ||
		strings.Index(s, string(EndsWith)) > -1 ||
		strings.Index(s, string(NotEndsWith)) > -1 ||
		strings.Index(s, string(Contains)) > -1 ||
		strings.Index(s, string(NotContains)) > -1
}