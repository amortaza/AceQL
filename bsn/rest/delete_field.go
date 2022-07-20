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
		c.JSON(500, err.Error())
		return logger.Err(err, "REST")
	}

	schema := stdsql.NewSchema()
	defer schema.Close()

	if err := schema.DeleteField(table, fieldName); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	return c.JSON(200, "")
}
