package flux

import (
	"encoding/json"
	"fmt"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
)

/*
***** ANY CHANGES HERE SHOULD ALSO UPDATE Initialize() function below!
 */
type Record struct {
	tableName string
	fields    []*dbschema.Field

	filterQuery *query.FilterQuery

	values     *RecordMap
	userValues *RecordMap

	fieldnameToFieldType map[string]dbschema.FieldType

	paginationIndex, paginationSize int

	orderBy          string
	orderByAscending bool

	crud CRUD
}

// NewRecord NEVER FAILS (0)
// grpc linked
func NewRecord(table *dbschema.Table, crud CRUD) *Record {
	return NewRecord_withDefinition(table.Name(), table.Fields(), crud)
}

// NewRecord_withDefinition NEVER FAILS (0)
// beause this is low level, it cannot take "Table" type
// it must take table name and field list (so we can hard code it when bootstrapping)
func NewRecord_withDefinition(tableName string, fields []*dbschema.Field, crud CRUD) *Record {

	/*
	***** ANY CHANGES HERE SHOULD ALSO UPDATE Initialize() function below!
	 */

	rec := &Record{
		filterQuery:    query.NewFilterQuery(crud.Compiler()),
		crud:           crud,
		tableName:      tableName,
		fields:         fields,
		paginationSize: -1,
	}

	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()

	rec.fieldnameToFieldType = make(map[string]dbschema.FieldType)

	for _, field := range fields {
		rec.fieldnameToFieldType[field.Name] = field.Type
	}

	return rec
}

func (rec *Record) MarshalJSON() ([]byte, error) {
	var bytes []byte
	var err error

	if bytes, err = json.Marshal(rec.GetMap()); err != nil {
		return nil, logger.Err(err, "record.MarshalJSON")
	}

	return bytes, nil
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

// grpc linked
func (rec *Record) Initialize() {
	rec.filterQuery = query.NewFilterQuery(rec.crud.Compiler())
	rec.paginationSize = -1
	rec.paginationIndex = 0
	rec.values = NewRecordMap()
	rec.userValues = NewRecordMap()
	rec.orderBy = ""
	rec.orderByAscending = false
}

// grpc linked
func (rec *Record) GetTable() string {
	return rec.tableName
}

// grpc linked
func (rec *Record) SetOrderByDesc(fields string) error {
	if fields == "" {
		return logger.Error(fmt.Sprintf("%s", "missing parameter \"fields\""), "flux.Record.SetOrderByDesc()")
	}

	rec.orderByAscending = false
	rec.orderBy = fields

	return nil
}

// grpc linked
func (rec *Record) SetOrderBy(fields string) error {
	if fields == "" {
		return logger.Error(fmt.Sprintf("%s", "missing parameter \"fields\""), "flux.Record.SetOrderBy()")
	}

	rec.orderByAscending = true
	rec.orderBy = fields

	return nil
}

// grpc linked
func (rec *Record) GetFieldType(fieldname string) (dbschema.FieldType, error) {
	fieldType, ok := rec.fieldnameToFieldType[fieldname]

	if !ok {
		return "", logger.Error("Field does not exist, see "+fieldname, "Record.GetFieldType")
	}

	return fieldType, nil
}

// grpc linked
func (rec *Record) Set(fieldname string, value string) error {
	fieldType, ok := rec.fieldnameToFieldType[fieldname]
	if !ok {
		return logger.Error("field does not exist, see "+fieldname, "Record.Set")
	}

	rec.userValues.SetFieldValue(fieldname, value, fieldType)

	return nil
}

// grpc linked
func (rec *Record) Close() error {
	if err := rec.crud.Close(); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) Insert() (string, error) {
	var id string
	var err error

	if id, err = rec.crud.Create(rec.tableName, rec.GetMap()); err != nil {
		return "", err
	}

	return id, nil
}

// grpc linked
func (rec *Record) Update() error {
	id, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	if err := rec.crud.Update(rec.tableName, id, rec.GetMap()); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) Delete() error {
	id, err := rec.Get("x_id")
	if err != nil {
		return err
	}

	if err := rec.crud.Delete(rec.tableName, id); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) Query() (int, error) {
	root, err := rec.filterQuery.GetRoot()
	if err != nil {
		return -1, err
	}

	var count int
	if count, err = rec.crud.Query(rec.tableName, rec.fields, root, rec.paginationIndex, rec.paginationSize, rec.orderBy, rec.orderByAscending); err != nil {
		return -1, err
	}

	return count, err
}

// Next will return false when no records left.
// grpc linked
func (rec *Record) Next() (bool, error) {
	rec.userValues = NewRecordMap()

	var err error
	var temp *RecordMap

	temp, err = rec.crud.Next()
	if err != nil {
		return false, err
	}

	rec.values = temp

	if rec.values == nil {
		rec.userValues = nil
		return false, nil
	}

	return true, err
}

// grpc linked
func (rec *Record) Pagination(index, size int) {
	rec.paginationIndex = index
	rec.paginationSize = size
}

// grpc linked
func (rec *Record) Get(field string) (string, error) {
	if rec.userValues.HasField(field) {
		return rec.userValues.GetFieldValue(field)
	}

	if rec.values.HasField(field) {
		return rec.values.GetFieldValue(field)
	}

	return "", logger.Error(fmt.Sprintf("field '%s' does not exist in record", field), "Record.Get")
}

// grpc linked
func (rec *Record) AddPK(id string) error {
	if err := rec.filterQuery.Add("x_id", query.Equals, id); err != nil {
		return logger.Err(err, "Record.AddPK")
	}

	return nil
}

//todo validate encoded query
// grpc linked
func (rec *Record) SetEncodedQuery(encodedQuery string) {
	rec.filterQuery.SetEncodedQuery(encodedQuery)
}

// todo error check
// grpc linked
func (rec *Record) Add(field string, op query.OpType, rhs string) error {
	if err := rec.filterQuery.Add(field, op, rhs); err != nil {
		return err
	}

	return nil
}

// todo error check
// grpc linked
func (rec *Record) AddEq(field string, rhs string) error {
	if err := rec.filterQuery.Add(field, query.Equals, rhs); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) AddOr(field string, op query.OpType, rhs string) error {
	if err := rec.filterQuery.AddOr(field, op, rhs); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) AndGroup() error {
	if err := rec.filterQuery.AndGroup(); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) OrGroup() error {
	if err := rec.filterQuery.OrGroup(); err != nil {
		return err
	}

	return nil
}

// grpc linked
func (rec *Record) Not() {
	rec.filterQuery.Not()
}
