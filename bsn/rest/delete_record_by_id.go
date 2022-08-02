package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
// DeleteRecordById http://localhost:8000/table/x_schema/id/0
func DeleteRecordById(c echo.Context) error {
	LOG_SOURCE := "REST.DeleteRecordById()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	tablename := c.Param("table")
	id := c.Param("id")

	r, err := stdsql.NewRecord(tablename)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}
	defer r.Close()

	if err := r.Add("x_id", query.Equals, id); err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	if _, err := r.Query(); err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	ok, err := r.Next()
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	if ok {
		if err := r.Delete(); err != nil {
			c.String(500, err.Error())
			return logger.Err(err, LOG_SOURCE)
		}
	}

	return c.String(200, "")
}
