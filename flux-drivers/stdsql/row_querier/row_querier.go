package row_querier

import (
	"database/sql"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql/compiler"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_runner"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/logger"
	"strconv"
)

type RowQuerier struct {
	rows   *sql.Rows
	fields []*dbschema.Field

	sqlRunner      *sql_runner.SqlRunner
	selectCompiler *compiler.SelectCompiler
}

func NewRowQuerier(sqlRunner *sql_runner.SqlRunner, tableName string, fields []*dbschema.Field, root node.Node) *RowQuerier {
	columns := dbschema.FieldsToNames(fields)
	selectCompiler := compiler.NewSelectCompiler(tableName, columns, root)

	return &RowQuerier{
		sqlRunner:      sqlRunner,
		selectCompiler: selectCompiler,

		fields: fields,
	}
}

func (query *RowQuerier) Close() error {
	if query.rows == nil {
		return nil
	}

	err := query.rows.Close()

	if err == nil {
		logger.Log("Closing DB Connection - successful", "RowQuerier.Close()")
	} else {
		logger.Log("Closing DB Connection - UNSUCCESSFUL", "RowQuerier.Close()")
		logger.Err(err, logger.ERROR)
	}

	return err
}

func (query *RowQuerier) Query(paginationIndex int, paginationSize int, orderBy string, orderByAscending bool) (int, error) {
	sqlstr, sqlstr_forCount, err := query.selectCompiler.Compile(paginationIndex, paginationSize, orderBy, orderByAscending)
	if err != nil {
		return -1, err
	}

	logger.Log(sqlstr, "SQL:RowQuerier.Query()")

	query.rows, err = query.sqlRunner.Query(sqlstr)
	if err != nil {
		return -1, err
	}

	rowcount, err2 := query.sqlRunner.Query(sqlstr_forCount)
	if err2 != nil {
		return -1, err2
	}

	defer rowcount.Close()

	rowcount.Next()

	count, err3 := readTotal(rowcount)
	if err3 != nil {
		return -1, err3
	}

	return count, nil
}

// fields should be known well before query is made
//query.fields, err = query.rows.Columns()
//if err != nil {
//	return fmt.Errorf("%v", err)
//}

func readTotal(rows *sql.Rows) (int, error) {
	columnPointers := make([]interface{}, 1)

	columns := []string{"total"}

	columnPointers[0] = &columns[0]

	if err := rows.Scan(columnPointers...); err != nil {
		return -1, logger.Err(err, "readTotal() could not scan")
	}

	value := columns[0]

	total, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return -1, logger.Err(err, "readTotal() could not parse answer")
	}

	return int(total), nil
}

// Next returns nil if there is no records left
func (query *RowQuerier) Next() (*flux.RecordMap, error) {
	has := query.rows.Next()

	if !has {
		return nil, nil
	}

	columns := make([]interface{}, len(query.fields))
	columnPointers := make([]interface{}, len(query.fields))

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	if err := query.rows.Scan(columnPointers...); err != nil {
		return nil, logger.Err(err, "RowQuerier.Next() rows.Scan")
	}

	valuesRecordMap := flux.NewRecordMap()

	for i, field := range query.fields {
		value := columnPointers[i].(*interface{})

		// the value that is set will *always* be string,
		// but the field-type as defined in the schema is respected.
		bytes := (*value).([]byte)
		stringValue := string(bytes)

		valuesRecordMap.SetFieldValue(field.Name, stringValue, field.Type)
	}

	return valuesRecordMap, nil
}
