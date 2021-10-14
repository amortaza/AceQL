package flux

import (
	"bytes"
	"fmt"
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/relations"
	"strconv"
)

type RecordMap struct {
	Data map[string] *TypedValue
}

func (recmap *RecordMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	datamap := recmap.Data

	// we sneak in "id" because React-Admin minimally requires "id"
	//datamap["id"] = datamap["x_id"]

	first := true
	for key, typedValue := range datamap {

		if first {
			first = false
		} else {
			buffer.WriteString(",")
		}

		if typedValue.IsString() {
			asStr := fmt.Sprintf("\"%s\" : \"%s\"", key, typedValue.GetString() )
			buffer.WriteString( asStr )

		} else if typedValue.IsBool() {
			asStr := fmt.Sprintf("\"%s\" : %s", key, strconv.FormatBool( typedValue.GetBool() ) )
			buffer.WriteString( asStr )

		} else if typedValue.IsNumber() {
			asStr := fmt.Sprintf("\"%s\" : %f", key, typedValue.GetNumber() )
			buffer.WriteString( asStr )
		} else {
			logger.Error("Typed value type unrecognized", "MarshalJSON()")
		}
	}

	buffer.WriteString("}")

	logger.Log(string(buffer.Bytes()), logger.JsonEncoding)

	return buffer.Bytes(), nil
}

func NewRecordMap() *RecordMap {
	return &RecordMap{
		Data: make(map[string] *TypedValue),
	}
}

func (recmap *RecordMap) PutStringByteArray(key string, bytes []byte) {
	typedValue := &TypedValue{}
	typedValue.SetStringByteArray( bytes )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) PutString(key string, value string) {
	typedValue := &TypedValue{}
	typedValue.SetString( value )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) PutNumberByteArray(key string, bytes []byte) {
	typedValue := &TypedValue{}
	typedValue.SetNumberByteArray( bytes )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) PutNumber(key string, value float32) {
	typedValue := &TypedValue{}
	typedValue.SetNumber( value )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) PutBoolByteArray(key string, bytes []byte) {
	typedValue := &TypedValue{}
	typedValue.SetBoolByteArray( bytes )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) PutBool(key string, value bool) {
	typedValue := &TypedValue{}
	typedValue.SetBool( value )

	recmap.Data[ key ] = typedValue
}

func (recmap *RecordMap) IsString(key string) (bool, error) {
	if !recmap.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue, _ := recmap.Data[ key ]

	return typedValue.fieldType == relations.String, nil
}

func (recmap *RecordMap) IsNumber(key string) (bool, error) {
	if !recmap.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue, _ := recmap.Data[ key ]

	return typedValue.fieldType == relations.Number, nil
}

func (recmap *RecordMap) IsBool(key string) (bool, error) {
	if !recmap.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue, _ := recmap.Data[ key ]

	return typedValue.fieldType == relations.Bool, nil
}

func (recmap *RecordMap) Get(key string) (string, error) {
	if !recmap.Has(key) {
		return "", fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue := recmap.Data[ key ]

	if !typedValue.IsString() {
		err := fmt.Errorf("typed-value is a '%s' and not a string", typedValue.fieldType)
		logger.Error(err, logger.MAIN)

		return "", err
	}

	//fmt.Println( "attempting to get string ", key, typedValue.GetString() ) // debug
	return typedValue.GetString(), nil
}

func (recmap *RecordMap) GetNumber(key string) (float32, error) {
	if !recmap.Has(key) {
		return 0, fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue := recmap.Data[ key ]

	if !typedValue.IsNumber() {
		err := fmt.Errorf("typed-value is a '%s' and not a number", typedValue.fieldType)
		logger.Error(err, logger.MAIN)

		return 0, err
	}

	return typedValue.GetNumber(), nil
}

func (recmap *RecordMap) GetBool(key string) (bool, error) {
	if !recmap.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	typedValue := recmap.Data[ key ]

	if !typedValue.IsBool() {
		err := fmt.Errorf("typed-value is a '%s' and not a bool", typedValue.fieldType)
		logger.Error(err, logger.MAIN)

		return false, err
	}

	fmt.Println( "attempting to get bool ", key, typedValue.GetBool() ) // debug
	return typedValue.GetBool(), nil
}

func (recmap *RecordMap) Has(key string) bool {
	_, ok := recmap.Data[key]

	return ok
}

func (recmap *RecordMap) Combine(other *RecordMap) *RecordMap {
	result := NewRecordMap()

	for k, v := range recmap.Data {
		result.Data[ k ] = v
	}

	for k, v := range other.Data {
		result.Data[ k ] = v
	}

	return result
}
