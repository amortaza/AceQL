package stdsql

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql/compiler"
	"github.com/amortaza/aceql/flux-drivers/stdsql/row_querier"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_generator"
	"github.com/amortaza/aceql/flux-drivers/stdsql/sql_runner"
	"github.com/amortaza/aceql/flux/node"
	"github.com/amortaza/aceql/flux/tableschema"
	"github.com/amortaza/aceql/flux/utils"
)

var global_sqlRunner *sql_runner.SqlRunner

func Init(driverName, dataSourceName string) {
	// "mysql", "clown:1844@/bsn"
	global_sqlRunner = sql_runner.NewSQLRunner(driverName, dataSourceName)
}

type CRUD struct {
	querier   *row_querier.RowQuerier
	compiler  node.Compiler
	sqlRunner *sql_runner.SqlRunner
}

// NewCRUD never returns nil
func NewCRUD() flux.CRUD {
	return &CRUD{
		sqlRunner: global_sqlRunner,
		compiler:  compiler.NewNodeCompiler(),
	}
}

func (crud *CRUD) Compiler() node.Compiler {
	return crud.compiler
}

func (crud *CRUD) Query(table string, fields []*tableschema.Field, root node.Node,
	paginationIndex int, paginationSize int,
	orderBy string, orderByAscending bool) (int, error) {

	crud.querier = row_querier.NewRowQuerier(crud.sqlRunner, table, fields, root)

	return crud.querier.Query(paginationIndex, paginationSize, orderBy, orderByAscending)
}

// Next returns nil if there are no records left
func (crud *CRUD) Next() (*flux.RecordMap, error) {
	return crud.querier.Next()
}

func (crud *CRUD) Create(table string, values *flux.RecordMap) (string, error) {
	sqlGenerator := sql_generator.NewRowInsert_SqlGenerator()

	newId := utils.NewUUID()

	sql := sqlGenerator.GenerateInsertSQL(table, newId, values)

	return newId, crud.sqlRunner.Run(sql)
}

func (crud *CRUD) Update(table string, id string, values *flux.RecordMap) error {
	sqlGenerator := sql_generator.NewRowUpdate_SqlGenerator()

	sql, err := sqlGenerator.GenerateSQL(table, id, values)
	if err != nil {
		return err
	}

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) Delete(table string, id string) error {
	sqlGenerator := sql_generator.NewRowDelete_SqlGenerator()

	sql := sqlGenerator.GenerateDeleteSQL(table, id)

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) CreateTable(name string) error {
	sqlGenerator := sql_generator.NewTableCreate_SqlGenerator()

	sql := sqlGenerator.GenerateCreateTableSQL(name)

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) DeleteTable(name string) error {
	sqlGenerator := sql_generator.NewTableDelete_SqlGenerator()

	sql := sqlGenerator.GenerateDeleteTableSQL(name)

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) CreateField(tablename string, field *tableschema.Field) error {
	sqlGenerator := sql_generator.NewFieldCreate_SqlGenerator()

	sql, err := sqlGenerator.GenerateCreateFieldSQL(tablename, field)
	if err != nil {
		return err
	}

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) DeleteField(tablename string, fieldname string) error {
	sqlGenerator := sql_generator.NewFieldDelete_SqlGenerator()

	sql := sqlGenerator.GenerateDeleteFieldSQL(tablename, fieldname)

	return crud.sqlRunner.Run(sql)
}

func (crud *CRUD) Close() error {
	if crud.querier == nil {
		return nil
	}

	return crud.querier.Close()
}
