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

func PostCSV(c echo.Context) error {
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
		return logger.Err(err, "Import CSV")
	}

	record, err := stdsql.NewRecord(table)
	if err != nil {
		return err
	}

	defer record.Close()

	for {
		values, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return logger.Err(err, "Import CSV")
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
