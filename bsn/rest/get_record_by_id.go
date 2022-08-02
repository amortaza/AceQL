package rest

import (
	"encoding/json"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

// !log
// http://localhost:8000/table/x_schema/0
func GetRecordById(c echo.Context) error {
	LOG_SOURCE := "REST.GetRecordById()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	name := c.Param("table")
	id := c.Param("id")

	r, err := stdsql.NewRecord(name)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}
	defer r.Close()

	if err := r.Add("x_id", query.Equals, id); err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	total, err := r.Query()
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	if total == 0 {
		c.Response().Header().Set("X-Total-Count", "0")
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

		return c.String(200, "")
	}

	_, err = r.Next()
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	b, err := json.Marshal(r)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, "???")
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.String(200, string(b))
}
