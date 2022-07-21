package dbschema

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
	"strings"
)

type FieldType string

const (
	String FieldType = "String"
	Number           = "Number"
	Bool             = "Bool"
)

func GetFieldTypeByName(name string) (FieldType, error) {
	name = strings.ToLower(name)

	if name == "string" {
		return String, nil
	}

	if name == "number" {
		return Number, nil
	}

	if name == "bool" {
		return Bool, nil
	}

	return "", logger.Error(fmt.Sprintf("no field-type has been defined for '%s'", name), "fieldtype.GetFieldTypeByName")
}
