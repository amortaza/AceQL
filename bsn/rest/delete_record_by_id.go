package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
)

// DeleteRecordById http://localhost:8000/table/x_schema/id/0
func DeleteRecordById(c echo.Context) error {
	tablename := c.Param("table")
	id := c.Param("id")

	r := stdsql.NewRecord(tablename)
	if r == nil {
		return c.String(500, "see logs")
	}

	if err := r.Add("x_id", query.Equals, id); err != nil {
		return c.String(500, err.Error())
	}

	if _, err := r.Query(); err != nil {
		return c.String(500, err.Error())
	}

	ok, err := r.Next()
	if err != nil {
		return c.String(500, err.Error())
	}

	if ok {
		if err := r.Delete(); err != nil {
			return c.String(500, err.Error())
		}
	}

	if err := r.Close(); err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, "")
}
