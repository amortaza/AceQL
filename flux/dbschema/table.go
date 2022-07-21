package dbschema

import (
	"fmt"
	"github.com/amortaza/aceql/logger"
)

type Table struct {
	name   string
	label  string
	fields []*Field

	fieldByName map[string]*Field
}

// NewTable never nil
func NewTable(name string) *Table {
	return &Table{
		name:        name,
		label:       name,
		fields:      nil,
		fieldByName: make(map[string]*Field),
	}
}

func (table *Table) Name() string {
	return table.name
}

func (table *Table) Label() string {
	return table.label
}

func (table *Table) SetLabel(label string) {
	table.label = label
}

func (table *Table) Fields() []*Field {
	return table.fields
}

func (table *Table) AddField(name string, label string, fieldtype FieldType) error {
	if name == "" {
		return logger.Error(fmt.Sprintf("missing name parameter"), "AddField()")
	}

	if label == "" {
		return logger.Error(fmt.Sprintf("missing label parameter"), "AddField()")
	}

	if fieldtype == "" {
		return logger.Error(fmt.Sprintf("missing field type parameter"), "AddField()")
	}

	field := NewField(name, label, fieldtype)

	table.fieldByName[name] = field

	table.fields = append(table.fields, field)

	return nil
}

func (table *Table) GetField(fieldname string) (*Field, error) {
	v, ok := table.fieldByName[fieldname]
	if !ok {
		return nil, logger.Error(fmt.Sprintf("field not found, see \"%s\"", fieldname), "table.GetField")
	}

	return v, nil
}
