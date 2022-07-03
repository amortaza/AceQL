package rest

import (
	"errors"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"strconv"
)

// http://localhost:8000/schema/x_choice_list
func GetSchemaByTable(c echo.Context) error {
	table := c.Param("table")

	r := stdsql.NewRecord("x_schema")
	if r == nil {
		return errors.New("see logs")
	}

	if err := r.Add("x_table", query.Equals, table); err != nil {
		return err
	}

	if _, err := r.Query(); err != nil {
		return err
	}

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, err := r.Next()
		if err != nil {
			return err
		}

		if !hasNext {
			break
		}

		list = append(list, r.GetMap())
	}

	if err := r.Close(); err != nil {
		return err
	}

	size := strconv.Itoa(len(list))

	c.Response().Header().Set("X-Total-Count", size)
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(200, list)
}
