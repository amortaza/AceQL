package row_querier

import (
	"database/sql"
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql/compiler"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_runner"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/relations"
)

type RowQuerier struct {
	rows   *sql.Rows
	fields [] *relations.Field

	sqlRunner      *sql_runner.SqlRunner
	selectCompiler *compiler.SelectCompiler
}

func NewRowQuerier(sqlRunner *sql_runner.SqlRunner, table string, fields [] *relations.Field, root node.Node) *RowQuerier {
	columns := relations.FieldsToNames( fields )
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

	return query.rows.Close()
}

func (query *RowQuerier) Query(paginationIndex int, paginationSize int) error {
	sqlstr, err := query.selectCompiler.Compile(paginationIndex, paginationSize)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	logger.Log( sqlstr, "SQL:RowQuerier.Query()" )

	query.rows, err = query.sqlRunner.Query(sqlstr)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// fields should be known well before query is made
	//query.fields, err = query.rows.Columns()
	//if err != nil {
	//	return fmt.Errorf("%v", err)
	//}

	return nil
}

// Next returns nil if there is no records left
func (query *RowQuerier) Next() (*flux.RecordMap, error) {
	has := query.rows.Next()

	if !has {
		return nil, nil
	}

	columns        := make([]interface{}, len(query.fields))
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
		value := columnPointers[ i ].(*interface{})

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
