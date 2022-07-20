package flux

import (
	"bytes"
	"fmt"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
)

type RecordMap struct {
	Data map[string]*TypedValue
}

func NewRecordMap() *RecordMap {
	return &RecordMap{
		Data: make(map[string]*TypedValue),
	}
}

func (recmap *RecordMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	datamap := recmap.Data

	first := true
	for key, typedValue := range datamap {

		if first {
			first = false
		} else {
			if _, err := buffer.WriteString(","); err != nil {
				return nil, logger.Err(err, "RecordMap.MarshalJSON")
			}
		}

		if typedValue.IsString() {
			asStr := fmt.Sprintf("\"%s\" : \"%s\"", key, typedValue.GetValue())
			if _, err := buffer.WriteString(asStr); err != nil {
				return nil, logger.Err(err, "???")
			}

		} else if typedValue.IsBool() {
			asStr := fmt.Sprintf("\"%s\" : %s", key, typedValue.GetValue())
			if _, err := buffer.WriteString(asStr); err != nil {
				return nil, logger.Err(err, "???")
			}

		} else if typedValue.IsNumber() {
			asStr := fmt.Sprintf("\"%s\" : %s", key, typedValue.GetValue())
			if _, err := buffer.WriteString(asStr); err != nil {
				return nil, logger.Err(err, "???")
			}
		} else {
			return nil, logger.Error("typed value type unrecognized", "RecordMap.MarshalJSON")
		}
	}

	if _, err := buffer.WriteString("}"); err != nil {
		return nil, logger.Err(err, "RecordMap.MarshalJSON")
	}

	logger.Info(string(buffer.Bytes()), "JSON-ENCODING")

	return buffer.Bytes(), nil
}

func (recmap *RecordMap) HasField(fieldName string) bool {
	_, ok := recmap.Data[fieldName]

	return ok
}

func (recmap *RecordMap) IsFieldString(fieldname string) (bool, error) {
	typedValue, ok := recmap.Data[fieldname]
	if !ok {
		return false, logger.Error("field "+fieldname+" not found", "RecordMap.IsFieldString")
	}

	return typedValue.fieldType == dbschema.String, nil
}

func (recmap *RecordMap) IsFieldNumber(fieldname string) (bool, error) {
	typedValue, ok := recmap.Data[fieldname]
	if !ok {
		return false, logger.Error("field "+fieldname+" not found in recordMap", "RecordMap.IsFieldNumber()")
	}

	return typedValue.fieldType == dbschema.Number, nil
}

func (recmap *RecordMap) IsFieldBool(fieldname string) (bool, error) {
	typedValue, ok := recmap.Data[fieldname]
	if !ok {
		return false, logger.Error("field "+fieldname+" not found in recordMap", "RecordMap.IsFieldBool()")
	}

	return typedValue.fieldType == dbschema.Bool, nil
}

func (recmap *RecordMap) GetFieldValue(fieldname string) (string, error) {
	typedValue, ok := recmap.Data[fieldname]
	if !ok {
		return "", logger.Error("field "+fieldname+" not found in recordMap", "RecordMap.GetFieldValue()")
	}

	return typedValue.GetValue(), nil
}

func (recmap *RecordMap) SetFieldValue(fieldname string, value string, fieldType dbschema.FieldType) {
	recmap.Data[fieldname] = NewTypedValue(value, fieldType)
}

func (recmap *RecordMap) Combine(other *RecordMap) *RecordMap {
	result := NewRecordMap()

	for k, v := range recmap.Data {
		result.Data[k] = v
	}

	for k, v := range other.Data {
		result.Data[k] = v
	}

	return result
}
