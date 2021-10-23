package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8000/table/x_schema/id/0
func DeleteRecordById(c echo.Context) error {

	name := c.Param("table")
	id := c.Param("id")

	r := stdsql.NewRecord(name)
	_ = r.Add("x_id", query.Equals, id)
	_ , _= r.Query()

	ok, _ := r.Next()

	if ok {
		err := r.Delete()
		if err != nil {
			return c.String(500, "something wrong in DeleteRecordById")
		}
	}

	return c.String(http.StatusOK, "")
}
