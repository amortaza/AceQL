package rest

import (
	"encoding/json"
	"github.com/amortaza/aceql/flux-drivers/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// http://localhost:8000/table/x_schema/0
func GetRecordById(c echo.Context) error {

	name := c.Param("table")
	id := c.Param("id")

	r := stdsql.NewRecord(name)
	defer r.Close()

	_ = r.Add("x_id", query.Equals, id)
	total , err := r.Query()
	if err != nil {
		logger.Error(err, "GetRecordById()")
		return c.String(http.StatusInternalServerError, "")
	}

	if total == 0 {
		c.Response().Header().Set("X-Total-Count", "0")
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

		return c.String(http.StatusOK, "")
	}

	_, err = r.Next()
	if err != nil {
		logger.Error(err, "GetRecordById()")
		return c.String(http.StatusInternalServerError, "")
	}

	b, err := json.Marshal(r)
	if err != nil {
		logger.Error(err, "GetRecordById()")
		return c.String(http.StatusInternalServerError, "")
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	return c.String(http.StatusOK, string(b))
}
