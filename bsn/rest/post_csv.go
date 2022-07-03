package rest

import (
	"encoding/csv"
	"errors"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
)

func PostCSV(c echo.Context) error {
	table := c.Param("table")

	file, err := c.FormFile("myfile")
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Error(err.Error(), "PostCSV")
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Error(err.Error(), "PostCSV")
	}

	defer src.Close()

	if err := importCSV(table, src); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	return c.JSON(200, "")
}

func importCSV(table string, src multipart.File) error {
	reader := csv.NewReader(src)

	headers, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return logger.Error(err.Error(), "Import CSV")
	}

	record := stdsql.NewRecord(table)
	if record == nil {
		return errors.New("see logs")
	}

	defer record.Close()

	for {
		values, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return logger.Error(err.Error(), "Import CSV")
		}

		if err := importRow(record, headers, values); err != nil {
			return err
		}
	}

	return nil
}

func importRow(record *flux.Record, headers []string, values []string) error {
	for i, header := range headers {
		if err := record.Set(header, values[i]); err != nil {
			return err
		}
	}

	if _, err := record.Insert(); err != nil {
		return err
	}

	//todo
	//record.Initialize()

	return nil
}
