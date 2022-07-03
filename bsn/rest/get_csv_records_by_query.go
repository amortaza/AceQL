package rest

import (
	"encoding/csv"
	"errors"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

func GetRecordsByQuery_CSV(c echo.Context) error {
	r, err := lookupRecords(c)
	if err != nil {
		return c.String(500, err.Error())
	}
	defer r.Close()

	total, err := r.Query()
	if err != nil {
		return c.String(500, err.Error())
	}

	//
	name := c.Param("table")

	c.Response().Header().Set("Content-Type", "text-csv")
	c.Response().Header().Set("Content-Disposition", "attachment;filename="+name+".csv")

	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))

	// this has to come BEFORE the writers
	c.String(200, "")

	writer := csv.NewWriter(c.Response())
	defer writer.Flush()

	writeRecords(writer, r)

	return nil
}

func writeRecords(writer *csv.Writer, r *flux.Record) error {
	hasNext, err := r.Next()
	if err != nil {
		return err
	}

	if !hasNext {
		//todo at least return the headers
		return nil
	}

	keys, err := writeHeader(writer, r)
	if err != nil {
		return err
	}

	if err := writeRecord(writer, r, keys); err != nil {
		return err
	}

	for {
		hasNext, err := r.Next()
		if err != nil {
			return err
		}

		if !hasNext {
			break
		}

		if err := writeRecord(writer, r, keys); err != nil {
			return err
		}
	}

	return nil
}

func writeHeader(writer *csv.Writer, r *flux.Record) ([]string, error) {
	data := r.GetMap().Data
	keys := make([]string, 0)

	for key, _ := range data {
		keys = append(keys, key)
	}

	if err := writer.Write(keys); err != nil {
		return nil, logger.Error(err.Error(), "CSV Export")
	}

	return keys, nil
}

func writeRecord(writer *csv.Writer, r *flux.Record, keys []string) error {
	values := make([]string, 0)

	for _, key := range keys {
		v, _ := r.Get(key)
		values = append(values, v)
	}

	if err := writer.Write(values); err != nil {
		return logger.Error(err.Error(), "CSV Export")
	}

	return nil
}

func lookupRecords(c echo.Context) (*flux.Record, error) {
	name := c.Param("table")
	encodedQuery := c.QueryParam("query")

	if encodedQuery != "" {
		logger.Log("query: "+encodedQuery, "CSV Export")
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

	r := stdsql.NewRecord(name)
	if r == nil {
		return nil, errors.New("see logs")
	}

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		return nil, logger.Error(err.Error(), "CSV Export")
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		return nil, logger.Error(err.Error(), "CSV Export")
	}

	r.Pagination(index, size)

	if orderByAscending {
		r.SetOrderBy(orderBy)
	} else {
		r.SetOrderByDesc(orderBy)
	}

	return r, nil
}
