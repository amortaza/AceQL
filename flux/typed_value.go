package flux

import (
	"github.com/amortaza/aceql/flux/tableschema"
)

type TypedValue struct {
	fieldType tableschema.FieldType
	value     string
}

func NewTypedValue(value string, fieldType tableschema.FieldType) *TypedValue {
	return &TypedValue{value: value, fieldType: fieldType}
}

func (t *TypedValue) IsString() bool {
	return t.fieldType == tableschema.String
}

func (t *TypedValue) IsNumber() bool {
	return t.fieldType == tableschema.Number
}

func (t *TypedValue) IsBool() bool {
	return t.fieldType == tableschema.Bool
}

func (t *TypedValue) SetValue(value string, fieldType tableschema.FieldType) {
	t.fieldType = fieldType
	t.value = value
}

func (t *TypedValue) GetValue() string {
	return t.value
}
