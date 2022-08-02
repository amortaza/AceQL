package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

// !log
// http://localhost:8000/schema/x_choice_list
func GetSchemaByTable(c echo.Context) error {
	LOG_SOURCE := "REST.GetSchemaByTable()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	table := c.Param("table")

	r, err := stdsql.NewRecord("x_schema")
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}
	defer r.Close()

	if err := r.Add("x_table", query.Equals, table); err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	if _, err := r.Query(); err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, err := r.Next()
		if err != nil {
			c.String(500, err.Error())
			return logger.Err(err, LOG_SOURCE)
		}

		if !hasNext {
			break
		}

		list = append(list, r.GetMap())
	}

	size := strconv.Itoa(len(list))

	c.Response().Header().Set("X-Total-Count", size)
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(200, list)
}
