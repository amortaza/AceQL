package query

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
