package rest

import (
	"errors"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

func GetRecordsByQuery(c echo.Context) error {
	name := c.Param("table")

	encodedQuery := c.QueryParam("query")

	if encodedQuery != "" {
		logger.Log("query: "+encodedQuery, "REST:GetRecordsByQuery()")
	}

	orderByAscending := true
	orderBy := c.QueryParam("order_by")
	if orderBy == "" {
		orderBy = c.QueryParam("order_by_desc")
		orderByAscending = false
	}

	paginationIndex := c.QueryParam("index")
	paginationSize := c.QueryParam("size")

	if paginationIndex == "" || paginationSize == "" {
		paginationIndex = "0"
		paginationSize = "100"
	}

	crud := stdsql.NewCRUD()

	r := flux.NewRecord(flux.GetTableSchema(name, crud), crud)
	if r == nil {
		return errors.New("see logs")
	}

	defer r.Close()

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		return err
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		return err
	}

	r.Pagination(index, size)

	if orderByAscending {
		r.SetOrderBy(orderBy)
	} else {
		r.SetOrderByDesc(orderBy)
	}

	total, err := r.Query()
	if err != nil {
		return err
	}

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		list = append(list, r.GetMap())
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(200, list)
}
