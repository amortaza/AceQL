package flux

import (
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/flux/tableschema"
	"github.com/amortaza/aceql/logger"
)

type Record struct {
	tableName string
	fields    []*tableschema.Field

	filterQuery *query.FilterQuery

	values     *RecordMap
	userValues *RecordMap

	fieldnameToFieldType map[string]tableschema.FieldType

	paginationIndex, paginationSize int

	orderBy          string
	orderByAscending bool

	crud CRUD
}

func NewRecord(table *tableschema.Table, crud CRUD) *Record {
	return NewRecord_withDefinition(table.Name(), table.Fields(), crud)
}

// NewRecord_withDefinition beause this is low level, it cannot take "Table" type
// it must take table name and field list (so we can hard code it when bootstrapping)
func NewRecord_withDefinition(tableName string, fields []*tableschema.Field, crud CRUD) *Record {
	rec := &Record{
		filterQuery:    query.NewFilterQuery(crud.Compiler()),
		crud:           crud,
		tableName:      tableName,
		fields:         fields,
		paginationSize: -1,
	}

	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()

	rec.fieldnameToFieldType = make(map[string]tableschema.FieldType)

	for _, field := range fields {
		rec.fieldnameToFieldType[field.Name] = field.Type
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
		m[key] = typedValue.value
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

func (rec *Record) GetFieldType(fieldname string) (tableschema.FieldType, error) {
	fieldType, ok := rec.fieldnameToFieldType[fieldname]

	if !ok {
		return "", logger.Error("Field does not exist, see "+fieldname, "Record.Set()")
	}

	return fieldType, nil
}

func (rec *Record) Set(fieldname string, value string) error {
	fieldType, ok := rec.fieldnameToFieldType[fieldname]
	if !ok {
		return logger.Error("field does not exist, see "+fieldname, "Record.Set()")
	}

	rec.userValues.SetFieldValue(fieldname, value, fieldType)

	return nil
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
	if rec.userValues.HasField(field) {
		return rec.userValues.GetFieldValue(field)
	}

	if rec.values.HasField(field) {
		return rec.values.GetFieldValue(field)
	}

	return "", fmt.Errorf("field '%s' does not exist in record", field)
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

func (rec *Record) AddEq(field string, rhs string) error {
	return rec.filterQuery.Add(field, query.Equals, rhs)
}

func (rec *Record) AddOr(field string, op query.OpType, rhs string) error {
	return rec.filterQuery.AddOr(field, op, rhs)
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
