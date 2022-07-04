package flux

import (
	"github.com/amortaza/aceql/flux/dbschema"
)

type TypedValue struct {
	fieldType dbschema.FieldType
	value     string
}

func NewTypedValue(value string, fieldType dbschema.FieldType) *TypedValue {
	return &TypedValue{value: value, fieldType: fieldType}
}

func (t *TypedValue) IsString() bool {
	return t.fieldType == dbschema.String
}

func (t *TypedValue) IsNumber() bool {
	return t.fieldType == dbschema.Number
}

func (t *TypedValue) IsBool() bool {
	return t.fieldType == dbschema.Bool
}

func (t *TypedValue) SetValue(value string, fieldType dbschema.FieldType) {
	t.fieldType = fieldType
	t.value = value
}

func (t *TypedValue) GetValue() string {
	return t.value
}
