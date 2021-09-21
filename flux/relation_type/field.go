package relation_type

type Field struct {
	Name string
	Type FieldType
}

func NewField(name string, fieldtype FieldType) *Field {
	return &Field{
		Name: name,
		Type: fieldtype,
	}
}
