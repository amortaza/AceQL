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

	r, err := stdsql.NewRecord(tablename)
	if err != nil {
		c.String(500, err.Error())
		return err
	}
	defer r.Close()

	if err := r.Add("x_id", query.Equals, id); err != nil {
		c.String(500, err.Error())
		return err
	}

	if _, err := r.Query(); err != nil {
		c.String(500, err.Error())
		return err
	}

	ok, err := r.Next()
	if err != nil {
		c.String(500, err.Error())
		return err
	}

	if ok {
		if err := r.Delete(); err != nil {
			c.String(500, err.Error())
			return err
		}
	}

	return c.String(200, "")
}
