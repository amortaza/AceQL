package flux

import (
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/relations"
	"strconv"
)

type TypedValue struct {
	fieldType relations.FieldType

	valueAsBool bool
	valueAsString string
	valueAsNumber float32
}

func (t *TypedValue) SetStringByteArray(bytes []byte) {
	t.fieldType = relations.String
	t.valueAsString = string(bytes)
}

func (t *TypedValue) SetString(value string) {
	t.fieldType = relations.String
	t.valueAsString = value
}

func (t *TypedValue) GetString() string {
	return t.valueAsString
}

func (t *TypedValue) SetBoolByteArray(bytes []byte ) {
	t.fieldType = relations.Bool
	t.valueAsBool = string(bytes) != "0"
}

func (t *TypedValue) SetBool(value bool ) {
	t.fieldType = relations.Bool
	t.valueAsBool = value
}

func (t *TypedValue) GetBool() bool {
	return t.valueAsBool
}

func (t *TypedValue) SetNumberByteArray(bytes []byte ) {
	t.fieldType = relations.Number

	i, err := strconv.ParseFloat(string(bytes), 32)

	if err != nil {
		logger.Error(err, "TypedValue.GetNumber()")
		panic( "TypedValue.GetNumber()" )
	}

	t.valueAsNumber = float32(i)
}

func (t *TypedValue) SetNumber(value float32 ) {
	t.fieldType = relations.Number
	t.valueAsNumber = value
}

func (t *TypedValue) GetNumber() float32 {
	return t.valueAsNumber
}

func (t *TypedValue) IsString() bool {
	return t.fieldType == relations.String
}

func (t *TypedValue) IsNumber() bool {
	return t.fieldType == relations.Number
}

func (t *TypedValue) IsBool() bool {
	return t.fieldType == relations.Bool
}

