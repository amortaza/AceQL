package tableschema

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

func (table *Table) AddField(name string, label string, fieldtype FieldType) {
	field := NewField(name, label, fieldtype)

	table.fieldByName[name] = field

	table.fields = append(table.fields, field)
}

func (table *Table) GetField(fieldname string) *Field {
	v, ok := table.fieldByName[fieldname]
	if !ok {
		logger.Error(fmt.Sprintf("field not found, see \"%s\"", fieldname), "table.GetField")
	}

	return v
}
