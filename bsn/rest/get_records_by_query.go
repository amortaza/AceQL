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

	crud := stdsql.NewCRUD()

	r := flux.NewRecord( flux.GetRelation(name,crud), crud )
	_ = r.Query()

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		//fmt.Println( "x_id is " , r.Get("x_id") )

		list = append(list, r.GetMap())
	}

	size := strconv.Itoa(len(list))

	c.Response().Header().Set("X-Total-Count", size)
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(http.StatusOK, list)
}
