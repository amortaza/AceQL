package tableschema

type Table struct {
	name   string
	label  string
	fields []*Field

	fieldByName map[string]*Field
}

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

func (table *Table) GetField(name string) *Field {
	return table.fieldByName[name]
}
