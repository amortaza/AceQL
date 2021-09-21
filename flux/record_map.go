package flux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/relation_type"
)

type RecordMap struct {
	Data map[string]interface{}
}

func (recmap *RecordMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	datamap := recmap.Data

	// we sneak in "id" because React-Admin minimally requires "id"
	datamap["id"] = datamap["x_id"]

	first := true
	for key, valuebytes := range datamap {

		if first {
			first = false
		} else {
			buffer.WriteString(",")
		}

		value := string(valuebytes.([]byte))

		valuemarshalled, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		asStr := fmt.Sprintf("\"%s\" : %s", key, string(valuemarshalled))

		buffer.WriteString(asStr)
	}

	buffer.WriteString("}")

	logger.Log(string(buffer.Bytes()), logger.JsonEncoding)

	return buffer.Bytes(), nil
}

func NewRecordMap() *RecordMap {
	return &RecordMap{
		Data: make(map[string]interface{}),
	}
}

func (recmap *RecordMap) Put(key string, value interface{}) {
	recmap.Data[key] = value
}

func (recmap *RecordMap) Get(key string) (string, error) {
	if !recmap.Has(key) {
		return "", fmt.Errorf("key not '%s' not found in map", key)
	}

	asByteArray, ok := recmap.Data[key].([]byte)
	if !ok {
		return "", fmt.Errorf("value for key'%s' not a string", key)
	}

	return string(asByteArray), nil
}

func (recmap *RecordMap) GetNumber(key string) (float32, error) {
	if !recmap.Has(key) {
		return 0, fmt.Errorf("key not '%s' not found in map", key)
	}

	value, ok := recmap.Data[key].(float32)
	if !ok {
		return 0, fmt.Errorf("value for key'%s' not a number", key)
	}

	return value, nil
}

func (recmap *RecordMap) GetBool(key string) (bool, error) {
	if !recmap.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	value, ok := recmap.Data[key].(bool)
	if !ok {
		return false, fmt.Errorf("value for key'%s' not a bool", key)
	}

	return value, nil
}

func (recmap *RecordMap) Type(key string) (relation_type.FieldType, error) {
	if !recmap.Has(key) {
		return "", fmt.Errorf("key not '%s' not found in map", key)
	}

	value := recmap.Data[key]

	if _, ok := value.(string); ok {
		return relation_type.String, nil
	}

	if _, ok := value.(float32); ok {
		return relation_type.Number, nil
	}

	if _, ok := value.(bool); ok {
		return relation_type.Bool, nil
	}

	return "", fmt.Errorf("cannot determine type of value for key '%s', see '%v'", key, value)
}

func (recmap *RecordMap) Has(key string) bool {
	_, ok := recmap.Data[key]

	return ok
}

func (recmap *RecordMap) Combine(other *RecordMap) *RecordMap {
	result := NewRecordMap()

	for k, v := range recmap.Data {
		result.Put(k, v)
	}

	for k, v := range other.Data {
		result.Put(k, v)
	}

	return result
}
