package relation_type

type Relation struct {
	name   string
	fields []*Field
}

func NewRelation(name string) *Relation {
	return &Relation{
		name:   name,
		fields: nil,
	}
}

func (relation *Relation) Name() string {
	return relation.name
}

func (relation *Relation) Fields() []*Field {
	return relation.fields
}

func (relation *Relation) AddField(name string, fieldtype FieldType) {
	field := NewField(name, fieldtype)

	relation.fields = append(relation.fields, field)
}
