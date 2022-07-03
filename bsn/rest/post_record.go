package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func PostRecord(c echo.Context) error {

	name := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Err(err, logger.Main)
	}

	id := createRecord(name, m)

	if id == "" {
		return c.JSON(500, "")
	}

	return c.JSON(200, id)
}

func createRecord(name string, m *echo.Map) string {
	crud := stdsql.NewCRUD()
	tableschema := flux.GetTableSchema(name, crud)
	rec := flux.NewRecord(tableschema, crud)
	defer rec.Close()

	for key, value := range *m {
		valueAsString := value.(string)

		rec.Set(key, valueAsString)
	}

	id, err := rec.Insert()
	if err != nil {
		logger.Err(err, "post_record.createRecord()")
		return ""
	}

	return id
}
