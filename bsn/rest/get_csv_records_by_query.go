package rest

import (
	"encoding/csv"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
)

// !log
func GetRecordsByQuery_CSV(c echo.Context) error {
	LOG_SOURCE := "REST.GetRecordsByQuery_CSV()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	r, err := lookupRecords(c)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}
	defer r.Close()

	total, err := r.Query()
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	name := c.Param("table")

	c.Response().Header().Set("Content-Type", "text-csv")
	c.Response().Header().Set("Content-Disposition", "attachment;filename="+name+".csv")

	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	c.Response().Header().Set("X-Total-Count", strconv.Itoa(total))

	// this has to come BEFORE the writers
	c.String(200, "")

	writer := csv.NewWriter(c.Response())
	defer writer.Flush()

	if err := writeRecords(writer, r); err != nil {
		// not writting 500 stream on purpose
		return logger.Err(err, LOG_SOURCE)
	}

	return nil
}

// !log
func writeRecords(writer *csv.Writer, r *flux.Record) error {
	LOG_SOURCE := "REST.writeRecords()"

	hasNext, err := r.Next()

	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	if !hasNext {
		//todo at least return the headers
		return nil
	}

	keys, err := writeHeader(writer, r)
	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	if err := writeRecord(writer, r, keys); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	for {
		hasNext, err := r.Next()

		if err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}

		if !hasNext {
			break
		}

		if err := writeRecord(writer, r, keys); err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	return nil
}

// !log
func writeHeader(writer *csv.Writer, r *flux.Record) ([]string, error) {
	LOG_SOURCE := "REST.writeHeader()"

	data := r.GetMap().Data
	keys := make([]string, 0)

	for key, _ := range data {
		keys = append(keys, key)
	}

	if err := writer.Write(keys); err != nil {
		return nil, logger.PushStackTrace(LOG_SOURCE, err)
	}

	return keys, nil
}

// !log
func writeRecord(writer *csv.Writer, r *flux.Record, keys []string) error {
	LOG_SOURCE := "REST.writeRecord()"

	values := make([]string, 0)

	var v string
	var err error

	for _, key := range keys {
		if v, err = r.Get(key); err != nil {
			return err
		}
		values = append(values, v)
	}

	if err := writer.Write(values); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	return nil
}

// !log
func lookupRecords(c echo.Context) (*flux.Record, error) {
	LOG_SOURCE := "REST.lookupRecords()"

	name := c.Param("table")
	encodedQuery := c.QueryParam("query")

	if encodedQuery != "" {
		logger.Info("query: "+encodedQuery, "CSV Export")
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

	r, err := stdsql.NewRecord(name)
	if err != nil {
		return nil, logger.PushStackTrace(LOG_SOURCE, err)
	}

	if encodedQuery != "" {
		r.SetEncodedQuery(encodedQuery)
	}

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		return nil, logger.PushStackTrace(LOG_SOURCE, err)
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		return nil, logger.PushStackTrace(LOG_SOURCE, err)
	}

	r.Pagination(index, size)

	if orderByAscending {
		if err := r.SetOrderBy(orderBy); err != nil {
			return nil, logger.PushStackTrace(LOG_SOURCE, err)
		}
	} else {
		if err := r.SetOrderByDesc(orderBy); err != nil {
			return nil, logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	return r, nil
}
