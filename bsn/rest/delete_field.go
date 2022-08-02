package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
func DeleteSchemaField(c echo.Context) error {
	LOG_SOURCE := "REST.DeleteSchemaField()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	table := c.Param("table")
	fieldName := c.Param("field")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	schema := stdsql.NewSchema()
	defer schema.Close()

	if err := schema.DeleteField(table, fieldName); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	return c.JSON(200, "")
}
