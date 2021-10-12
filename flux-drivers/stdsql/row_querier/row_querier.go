package row_querier

import (
	"database/sql"
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql/compiler"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_runner"
	"github.com/amortaza/aceql/flux/node"
)

type RowQuerier struct {
	rows    *sql.Rows
	columns []string

	sqlRunner      *sql_runner.SqlRunner
	selectCompiler *compiler.SelectCompiler
}

func NewRowQuerier(sqlRunner *sql_runner.SqlRunner, table string, root node.Node) *RowQuerier {
	var selectCompiler = compiler.NewSelectCompiler(table, root)

	return &RowQuerier{
		sqlRunner:      sqlRunner,
		selectCompiler: selectCompiler,
	}
}

func (query *RowQuerier) Query() error {
	sqlstr, err := query.selectCompiler.Compile()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	query.rows, err = query.sqlRunner.Query(sqlstr)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	query.columns, err = query.rows.Columns()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// Next returns nil if there is no records left
func (query *RowQuerier) Next() (*flux.RecordMap, error) {
	has := query.rows.Next()

	if !has {
		return nil, nil
	}

	columns        := make([]interface{}, len(query.columns))
	columnPointers := make([]interface{}, len(query.columns))

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	if err := query.rows.Scan(columnPointers...); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	valuesRecordMap := flux.NewRecordMap()

	for i, name := range query.columns {
		val := columnPointers[i].(*interface{})
		valuesRecordMap.Put(name, *val)
	}

	return valuesRecordMap, nil
}
