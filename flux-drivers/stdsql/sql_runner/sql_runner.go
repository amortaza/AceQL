package sql_runner

import (
	"database/sql"
	"fmt"
	"github.com/amortaza/aceql/flux-drivers/logger"
	"time"

	// need to have driver available otherwise sql.Open() fails
	_ "github.com/go-sql-driver/mysql"
)

type SqlRunner struct {
	db       *sql.DB
	lastPing time.Time

	driverName, dataSourceName string
}

func NewSQLRunner(driverName string, dataSourceName string) *SqlRunner {
	return &SqlRunner{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}
}

func (runner *SqlRunner) Run(sql string) error {
	if err := runner.ping(); err != nil {
		return err
	}

	logger.Log(sql, logger.SQL)

	_, err := runner.db.Exec(sql)

	return err
}

func (runner *SqlRunner) Query(sql string) (*sql.Rows, error) {
	if err := runner.ping(); err != nil {
		return nil, err
	}

	return runner.db.Query(sql)
}

func (runner *SqlRunner) ping() error {
	if runner.db == nil {
		var err error

		runner.db, err = sql.Open(runner.driverName, runner.dataSourceName)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		runner.lastPing = time.Now()

		return nil
	}

	if time.Since(runner.lastPing) < 1*time.Minute {
		return nil
	}

	runner.lastPing = time.Now()

	err := runner.db.Ping()

	return err
}
