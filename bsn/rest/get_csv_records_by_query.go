package rest

import (
	"encoding/csv"
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetRecordsByQuery_CSV(c echo.Context) error {

	r, _ := lookupRecords(c)
	defer r.Close()

	total, err := r.Query()
	if err != nil {
		return err
	}

	//
	name := c.Param("table")

	c.Response().Header().Set("Content-Type", "text-csv")
	c.Response().Header().Set("Content-Disposition", "attachment;filename="+name+".csv")

	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))

	c.String(http.StatusOK, "")

	writer := csv.NewWriter(c.Response())

	writeRecords(writer, r)

	writer.Flush()

	return nil
}

func writeRecords(writer *csv.Writer, r *flux.Record) {
	hasNext, _ := r.Next()

	if !hasNext {
		//todo at least return the headers
		return
	}

	keys := writeHeader(writer, r)
	writeRecord(writer, r, keys)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		writeRecord(writer, r, keys)
	}
}

func writeHeader(writer *csv.Writer, r *flux.Record) []string {
	data := r.GetMap().Data
	keys := make([]string, 0)

	for key, _ := range data {
		keys = append(keys, key)
	}

	writer.Write(keys)

	return keys
}

func writeRecord(writer *csv.Writer, r *flux.Record, keys []string) {
	values := make([]string, 0)

	for _, key := range keys {
		v, _ := r.Get(key)
		values = append(values, v)
	}

	writer.Write(values)
}

func lookupRecords(c echo.Context) (*flux.Record, error) {
	name := c.Param("table")
	encodedQuery := c.QueryParam("query")

	if encodedQuery != "" {
		fmt.Println("query: " + encodedQuery) // debug
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

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		return nil, err
	}

	r.Pagination(index, size)

	if orderByAscending {
		r.SetOrderBy(orderBy)
	} else {
		r.SetOrderByDesc(orderBy)
	}

	return r, nil
}
