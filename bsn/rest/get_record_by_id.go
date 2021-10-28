package rest

import (
	"encoding/json"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8000/table/x_schema/0
func GetRecordById(c echo.Context) error {

	name := c.Param("table")
	id := c.Param("id")

	r := stdsql.NewRecord(name)
	_ = r.Add("x_id", query.Equals, id)
	_ , _ = r.Query()

	_, _ = r.Next()

	b, _ := json.Marshal(r)

	r.Close()

	return c.String(http.StatusOK, string(b))
}
