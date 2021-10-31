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

	encodedQuery := c.QueryParam("query")

	paginationIndex := c.QueryParam("index")
	paginationSize := c.QueryParam("size")

	if paginationIndex == "" || paginationSize == "" {
		paginationIndex = "0"
		paginationSize = "100"
	}

	crud := stdsql.NewCRUD()

	r := flux.NewRecord( flux.GetRelation(name,crud), crud )
	defer r.Close()

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err1 := strconv.Atoi(paginationIndex)
	if err1 != nil {
		return err1
	}

	size, err2 := strconv.Atoi(paginationSize)
	if err2 != nil {
		return err2
	}

	r.Pagination(index, size)

	total , _ := r.Query()

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		//fmt.Println( "x_id is " , r.Get("x_id") )

		list = append(list, r.GetMap())
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(http.StatusOK, list)
}
