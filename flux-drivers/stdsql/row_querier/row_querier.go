package row_querier

import (
	"database/sql"
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql/compiler"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_runner"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/table"
	"strconv"
)

type RowQuerier struct {
	rows   *sql.Rows
	fields []*table.Field

	sqlRunner      *sql_runner.SqlRunner
	selectCompiler *compiler.SelectCompiler
}

func NewRowQuerier(sqlRunner *sql_runner.SqlRunner, table string, fields []*table.Field, root node.Node) *RowQuerier {
	columns := table.FieldsToNames(fields)
	selectCompiler := compiler.NewSelectCompiler(table, columns, root)

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
		//debug logger.Log("Closing DB Connection - successful", "RowQuerier.Close()")
	} else {
		//debug logger.Log("Closing DB Connection - UNSUCCESSFUL", "RowQuerier.Close()")
		logger.Error(err, logger.ERROR)
	}

	return err
}

func (query *RowQuerier) Query(paginationIndex int, paginationSize int, orderBy string, orderByAscending bool) (int, error) {
	sqlstr, sqlstr_forCount, err := query.selectCompiler.Compile(paginationIndex, paginationSize, orderBy, orderByAscending)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}

	//logger.Log( sqlstr, "SQL:RowQuerier.Query()" )

	query.rows, err = query.sqlRunner.Query(sqlstr)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}

	rowcount, err2 := query.sqlRunner.Query(sqlstr_forCount)
	if err2 != nil {
		return -1, fmt.Errorf("%v", err2)
	}

	rowcount.Next()

	count, err3 := readTotal(rowcount)
	if err3 != nil {
		return -1, fmt.Errorf("%v", err3)
	}

	//debug logger.Log( "Closing DB Connection for COUNT(1)", "SQL:RowQuerier.Query()" )
	rowcount.Close()

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
		logger.Error(err, "readTotal() could not scan")
		return -1, err
	}

	value := columns[0]

	total, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		logger.Error(err, "readTotal() could not parse answer")
		return -1, err
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
		logger.Error(err, "RowQuerier.Next() rows.Scan")
		return nil, err
	}

	valuesRecordMap := flux.NewRecordMap()

	for i, field := range query.fields {
		value := columnPointers[i].(*interface{})

		//f mt.Println( field.Name )
		if field.IsString() {
			valuesRecordMap.PutStringByteArray(field.Name, (*value).([]byte))

		} else if field.IsNumber() {
			valuesRecordMap.PutNumberByteArray(field.Name, (*value).([]byte))

		} else if field.IsBool() {
			valuesRecordMap.PutBoolByteArray(field.Name, (*value).([]byte))

		} else {
			err := fmt.Errorf("\"in Next(), field type is unknown, see\"%s : %s\"", field.Name, field.Type)
			logger.Error(err, "RowQuerier.Next()")
			return nil, err
		}
	}

	return valuesRecordMap, nil
}
