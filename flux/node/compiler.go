package node

type Compiler interface {
	AndCompile(*And) (string, error)
	ColumnCompile(*Column) (string, error)
	ContainsCompile(*Contains) (string, error)
	EndsWithCompile(*EndsWith) (string, error)
	EqualCompile(*Equals) (string, error)
	GreaterThanCompile(*GreaterThan) (string, error)
	InCompile(*In) (string, error)
	IsNullCompile(*IsNull) (string, error)
	LessThanCompile(*LessThan) (string, error)
	NotCompile(*Not) (string, error)
	NumberCompile(*Number) (string, error)
	NumberListCompile(*NumberList) (string, error)
	OrCompile(*Or) (string, error)
	StartsWithCompile(*StartsWith) (string, error)
	StringCompile(*String) (string, error)
	StringListCompile(*StringList) (string, error)
}
