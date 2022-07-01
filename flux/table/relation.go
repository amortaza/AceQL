package table

type Relation struct {
	name   string
	label  string
	fields []*Field

	fieldByName map[string]*Field
}

func NewRelation(name string) *Relation {
	return &Relation{
		name:        name,
		fields:      nil,
		fieldByName: make(map[string]*Field),
	}
}

func (relation *Relation) Name() string {
	return relation.name
}

func (relation *Relation) Label() string {
	return relation.label
}

func (relation *Relation) SetLabel(label string) {
	relation.label = label
}

func (relation *Relation) Fields() []*Field {
	return relation.fields
}

func (relation *Relation) AddField(name string, label string, fieldtype FieldType) {
	field := NewField(name, label, fieldtype)

	relation.fieldByName[name] = field

	relation.fields = append(relation.fields, field)
}

func (relation *Relation) GetField(name string) *Field {
	return relation.fieldByName[name]
}
