package rest

import (
	"encoding/json"
	"errors"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"strconv"
)

// http://localhost:8000/table/x_schema/0
func GetRecordById(c echo.Context) error {
	name := c.Param("table")
	id := c.Param("id")

	r := stdsql.NewRecord(name)
	if r == nil {
		return errors.New("see logs")
	}
	defer r.Close()

	if err := r.Add("x_id", query.Equals, id); err != nil {
		return c.String(500, err.Error())
	}

	total, err := r.Query()
	if err != nil {
		return c.String(500, err.Error())
	}

	if total == 0 {
		c.Response().Header().Set("X-Total-Count", "0")
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

		return c.String(200, "")
	}

	_, err = r.Next()
	if err != nil {
		return c.String(500, err.Error())
	}

	b, err := json.Marshal(r)
	if err != nil {
		return c.String(500, err.Error())
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.String(200, string(b))
}
