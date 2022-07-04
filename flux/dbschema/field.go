package dbschema

type Field struct {
	Name  string
	Label string
	Type  FieldType
}

func NewField(name string, label string, fieldtype FieldType) *Field {
	return &Field{
		Name:  name,
		Label: label,
		Type:  fieldtype,
	}
}

func (field *Field) IsString() bool {
	return field.Type == String
}

func (field *Field) IsNumber() bool {
	return field.Type == Number
}

func (field *Field) IsBool() bool {
	return field.Type == Bool
}

func FieldsToNames(fields []*Field) []string {
	names := make([]string, len(fields))

	for i, field := range fields {
		names[i] = field.Name
	}

	return names
}
