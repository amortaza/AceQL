package sql_runner

import (
	"database/sql"
	"github.com/amortaza/aceql/logger"
	"sync"
	"time"

	// need to have driver available otherwise sql.Open() fails
	_ "github.com/go-sql-driver/mysql"
)

var lock sync.Mutex

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
	lock.Lock()
	defer lock.Unlock()

	if err := runner.ping(); err != nil {
		return err
	}

	logger.Log(sql, "SQL:SqlRunner.Run()")

	_, err := runner.db.Exec(sql)
	if err != nil {
		return logger.Err(err, "SqlRunner.Run")
	}

	return err
}

func (runner *SqlRunner) Query(sql string) (*sql.Rows, error) {
	lock.Lock()
	defer lock.Unlock()

	if err := runner.ping(); err != nil {
		return nil, logger.Err(err, "SqlRunner.Query")
	}

	logger.Log(sql, "SQL:SqlRunner.Query()")

	return runner.db.Query(sql)
}

func (runner *SqlRunner) ping() error {
	if runner.db == nil {
		var err error

		runner.db, err = sql.Open(runner.driverName, runner.dataSourceName)
		if err != nil {
			return logger.Err(err, "SqlRunner.ping")
		}

		runner.lastPing = time.Now()

		return nil
	}

	if time.Since(runner.lastPing) < 1*time.Minute {
		return nil
	}

	runner.lastPing = time.Now()

	err := runner.db.Ping()
	if err != nil {
		return logger.Err(err, "SqlRunner.ping")
	}

	return nil
}
