package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func DeleteSchemaField(c echo.Context) error {
	table := c.Param("table")
	fieldName := c.Param("field")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	schema := stdsql.NewSchema()

	schema.DeleteField(table, fieldName)

	schema.Close()

	return c.JSON(200, "")
}
