package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetRecordsByQuery(c echo.Context) error {
	name := c.Param("table")

	r := flux.NewRecord(name, stdsql.NewCRUD())
	_ = r.Query()

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
