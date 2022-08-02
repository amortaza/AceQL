package rest

import (
	"encoding/csv"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
)

// !log
func PostCSV(c echo.Context) error {
	LOG_SOURCE := "REST.PostCSV()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	table := c.Param("table")

	file, err := c.FormFile("myfile")
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "PostCSV")
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "PostCSV")
	}

	defer src.Close()

	if err := importCSV(table, src); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	return c.JSON(200, "")
}

// !log
func importCSV(table string, src multipart.File) error {
	LOG_SOURCE := "REST.importCSV()"

	reader := csv.NewReader(src)

	headers, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	record, err := stdsql.NewRecord(table)
	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	defer record.Close()

	for {
		values, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return logger.PushStackTrace(LOG_SOURCE, err)
		}

		if err := importRow(record, headers, values); err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	return nil
}

// !log
func importRow(record *flux.Record, headers []string, values []string) error {
	LOG_SOURCE := "REST.importRow()"

	for i, header := range headers {
		if err := record.Set(header, values[i]); err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	if _, err := record.Insert(); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	record.Initialize()

	return nil
}
