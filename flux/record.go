package flux

import (
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/relations"
)

type Record struct {
	relationName string
	fields [] *relations.Field

	filterQuery  *query.FilterQuery

	values     *RecordMap
	userValues *RecordMap

	nameToType map[string] *relations.FieldType

	crud CRUD
}

func NewRecord(relation *relations.Relation, crud CRUD) *Record {
	return NewRecord_withDefinition( relation.Name(), relation.Fields(), crud )
}

// beause this is low level, it cannot take "Relation" type
// it must take table name and field list (so we can hard code it when bootstrapping)
func NewRecord_withDefinition(relationName string, fields [] *relations.Field, crud CRUD) *Record {
	rec := &Record{
		filterQuery:  query.NewFilterQuery(crud.Compiler()),
		crud:         crud,
		relationName: relationName,
		fields: fields,
	}

	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()

	rec.nameToType = make(map[string] *relations.FieldType)

	for _, field := range fields {
		rec.nameToType[ field.Name ] = &field.Type
	}

	return rec
}

func (rec *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(rec.GetMap())
}

func (rec *Record) RelationName() string {
	return rec.relationName
}

func (rec *Record) GetMap() *RecordMap {
	return rec.values.Combine(rec.userValues)
}

func (rec *Record) Set( fieldname string, value interface{} ) {
	fieldType, ok := rec.nameToType[ fieldname ]
	if !ok {
		logger.Error("Field does not exist, see " + fieldname, "Record.Set()")
		return
	}

	if *fieldType == relations.String {
		rec.userValues.PutString(fieldname, value.(string))

	} else if *fieldType == relations.Number {
		rec.userValues.PutNumber(fieldname, value.(float32))

	} else if *fieldType == relations.Bool {
		rec.userValues.PutBool(fieldname, value.(bool))

	} else {
		logger.Error("Field type unrecognized, see " + fieldname + " : " + string(*fieldType), "Record.Set()")
	}
}

func (rec *Record) Insert() (string, error) {
	return rec.crud.Create(rec.relationName, rec.GetMap())
}

func (rec *Record) Update() error {
	pk, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	//v, _ := rec.Get("x_name")
	//fmt.Println( "bro ", v )

	return rec.crud.Update(rec.relationName, pk, rec.GetMap())
}

func (rec *Record) Delete() error {
	pk, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	return rec.crud.Delete(rec.relationName, pk)
}

func (rec *Record) Query() error {
	root, err := rec.filterQuery.GetRoot()
	if err != nil {
		return err
	}

	return rec.crud.Query(rec.relationName, rec.fields, root)
}

// Next will return false when no records left.
func (rec *Record) Next() (bool, error) {
	rec.userValues = NewRecordMap()

	var err error

	rec.values, err = rec.crud.Next()

	if rec.values == nil {
		rec.userValues = nil
		return false, nil
	}

	return true, err
}

func (rec *Record) Get(field string) (string, error) {
	if rec.userValues.Has(field) {
		return rec.userValues.Get(field)
	}

	if rec.values.Has(field) {
		return rec.values.Get(field)
	}

	return "", fmt.Errorf("field '%s' does not exist in record", field)
}

func (rec *Record) GetNumber(field string) (float32, error) {
	if rec.userValues.Has(field) {
		return rec.userValues.GetNumber(field)
	}

	if rec.values.Has(field) {
		return rec.values.GetNumber(field)
	}

	return 0, fmt.Errorf("field '%s' does not exist in record", field)
}

func (rec *Record) GetBool(field string) (bool, error) {
	if rec.userValues.Has(field) {
		return rec.userValues.GetBool(field)
	}

	if rec.values.Has(field) {
		return rec.values.GetBool(field)
	}

	return true, fmt.Errorf("field '%s' does not exist in record", field)
}

func (rec *Record) AddPrimaryKey(id string) error {
	return rec.filterQuery.Add("x_id", query.Equals, id)
}

func (rec *Record) Add(field string, op query.OpType, rhs string) error {
	return rec.filterQuery.Add(field, op, rhs)
}

func (rec *Record) AddNumber(field string, op query.OpType, rhs float32) error {
	return rec.filterQuery.AddNumber(field, op, rhs)
}

func (rec *Record) AddOr(field string, op query.OpType, rhs string) error {
	return rec.filterQuery.AddOr(field, op, rhs)
}

func (rec *Record) AddOrNumber(field string, op query.OpType, rhs float32) error {
	return rec.filterQuery.AddOrNumber(field, op, rhs)
}

func (rec *Record) AndGroup() error {
	return rec.filterQuery.AndGroup()
}

func (rec *Record) OrGroup() error {
	return rec.filterQuery.OrGroup()
}

func (rec *Record) Not() {
	rec.filterQuery.Not()
}
