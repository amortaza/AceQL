package flux

import (
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/query"
)

type Record struct {
	relationName string
	filterQuery  *query.FilterQuery

	fieldnames []string

	values     *RecordMap
	userValues *RecordMap

	crud CRUD
}

func NewRecord(relationName string, crud CRUD) *Record {
	rec := &Record{
		filterQuery:  query.NewFilterQuery(crud.Compiler()),
		crud:         crud,
		relationName: relationName,
	}

	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()

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

func (rec *Record) Set(field string, value interface{}) {
	rec.userValues.Put(field, value)
}

func (rec *Record) Insert() (string, error) {
	return rec.crud.Create(rec.relationName, rec.GetMap())
}

func (rec *Record) Update() error {
	pk, err := rec.Get("x_id")
	if err != nil {
		return err
	}

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

	return rec.crud.Query(rec.relationName, root)
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
