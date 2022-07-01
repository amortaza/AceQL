package tableschema

import (
	"fmt"
)

type FieldType string

const (
	String FieldType = "String"
	Number           = "Number"
	Bool             = "Bool"
)

func GetFieldTypeByName(name string) (FieldType, error) {
	if name == "String" {
		return String, nil
	}

	if name == "Number" {
		return Number, nil
	}

	if name == "Bool" {
		return Bool, nil
	}

	return "", fmt.Errorf("no field-type has been defined for '%s'", name)
}
