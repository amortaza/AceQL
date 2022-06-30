package rest

import (
	"encoding/csv"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
)

func PostCSV(c echo.Context) error {
	table := c.Param("table")

	file, err := c.FormFile("myfile")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	importCSV(table, src)

	return c.JSON(200, "")
}

func importCSV(table string, src multipart.File) error {
	reader := csv.NewReader(src)
	headers, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	record := stdsql.NewRecord(table)
	defer record.Close()

	for {
		values, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		importRow(record, headers, values)
	}

	return nil
}

func importRow(record *flux.Record, headers []string, values []string) {
	for i, header := range headers {
		record.Set(header, values[i])
	}

	record.Insert()
	//todo
	//record.Initialize()
}
