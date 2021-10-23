package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// http://localhost:8000/schema/x_choice_list
func GetSchemaByTable(c echo.Context) error {

	table := c.Param("table")

	r := stdsql.NewRecord("x_schema")
	_ = r.Add("x_type", query.Equals, "field")
	_ = r.Add("x_table", query.Equals, table)
	_, _ = r.Query()

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		list = append(list, r.GetMap())
	}

	size := strconv.Itoa(len(list))

	c.Response().Header().Set("X-Total-Count", size)
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(http.StatusOK, list)
}
