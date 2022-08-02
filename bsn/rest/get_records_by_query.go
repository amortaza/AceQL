package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

// !log
func GetRecordsByQuery(c echo.Context) error {
	LOG_SOURCE := "REST.GetRecordsByQuery()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	name := c.Param("table")

	encodedQuery := c.QueryParam("query")

	if encodedQuery != "" {
		logger.Info("query: "+encodedQuery, "REST:GetRecordsByQuery()")
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

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	r := flux.NewRecord(tableschema, crud)
	defer r.Close()

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	r.Pagination(index, size)

	if orderBy != "" {
		if orderByAscending {
			if err := r.SetOrderBy(orderBy); err != nil {
				return err
			}
		} else {
			if err := r.SetOrderByDesc(orderBy); err != nil {
				return err
			}
		}
	}

	total, err := r.Query()
	if err != nil {
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

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

	return c.JSON(200, list)
}
