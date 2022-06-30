package flux

import (
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/relations"
)

type Record struct {
	tableName string
	fields    []*relations.Field

	filterQuery *query.FilterQuery

	values     *RecordMap
	userValues *RecordMap

	nameToType map[string]*relations.FieldType

	paginationIndex, paginationSize int

	orderBy          string
	orderByAscending bool

	crud CRUD
}

func NewRecord(relation *relations.Relation, crud CRUD) *Record {
	return NewRecord_withDefinition(relation.Name(), relation.Fields(), crud)
}

// beause this is low level, it cannot take "Relation" type
// it must take table name and field list (so we can hard code it when bootstrapping)
func NewRecord_withDefinition(relationName string, fields []*relations.Field, crud CRUD) *Record {
	rec := &Record{
		filterQuery:    query.NewFilterQuery(crud.Compiler()),
		crud:           crud,
		tableName:      relationName,
		fields:         fields,
		paginationSize: -1,
	}

	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()

	rec.nameToType = make(map[string]*relations.FieldType)

	for _, field := range fields {
		rec.nameToType[field.Name] = &field.Type
	}

	return rec
}

func (rec *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(rec.GetMap())
}

func (rec *Record) GetTable() string {
	return rec.tableName
}

func (rec *Record) GetMap() *RecordMap {
	return rec.values.Combine(rec.userValues)
}

// GetMapGRPC : With GRPC we can only handle map[string]string hence this function.
func (rec *Record) GetMapGRPC() map[string]string {
	rmap := rec.GetMap()
	m := make(map[string]string)

	for key, typedValue := range rmap.Data {
		// For now we are only going to handle strings.
		if typedValue.IsString() {
			m[key] = typedValue.valueAsString
		} else if typedValue.IsNumber() {
			m[key] = "Number NOT Supported yet!"
		} else if typedValue.IsBool() {
			m[key] = "Bool NOT Supported yet!"
		} else {
			m[key] = "This is impossible in record.go"
		}
	}

	return m
}

func (rec *Record) SetOrderByDesc(fields string) {
	rec.orderByAscending = false
	rec.orderBy = fields
}

func (rec *Record) SetOrderBy(fields string) {
	rec.orderByAscending = true
	rec.orderBy = fields
}

func (rec *Record) Set(fieldname string, value interface{}) {
	fieldType, ok := rec.nameToType[fieldname]
	if !ok {
		logger.Error("Field does not exist, see "+fieldname, "Record.Set()")
		return
	}

	if *fieldType == relations.String {
		rec.userValues.PutString(fieldname, value.(string))

	} else if *fieldType == relations.Number {
		rec.userValues.PutNumber(fieldname, value.(float32))

	} else if *fieldType == relations.Bool {
		rec.userValues.PutBool(fieldname, value.(bool))

	} else {
		logger.Error("Field type unrecognized, see "+fieldname+" : "+string(*fieldType), "Record.Set()")
	}
}

func (rec *Record) Close() error {
	return rec.crud.Close()
}

func (rec *Record) Insert() (string, error) {
	return rec.crud.Create(rec.tableName, rec.GetMap())
}

func (rec *Record) Update() error {
	pk, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	return rec.crud.Update(rec.tableName, pk, rec.GetMap())
}

func (rec *Record) Delete() error {
	pk, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	return rec.crud.Delete(rec.tableName, pk)
}

func (rec *Record) Query() (int, error) {
	root, err := rec.filterQuery.GetRoot()
	if err != nil {
		return -1, err
	}

	return rec.crud.Query(rec.tableName, rec.fields, root, rec.paginationIndex, rec.paginationSize, rec.orderBy, rec.orderByAscending)
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

func (rec *Record) Pagination(index, size int) {
	rec.paginationIndex = index
	rec.paginationSize = size
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

func (rec *Record) AddPK(id string) error {
	return rec.filterQuery.Add("x_id", query.Equals, id)
}

func (rec *Record) SetEncodedQuery(encodedQuery string) {
	rec.filterQuery.SetEncodedQuery(encodedQuery)
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
