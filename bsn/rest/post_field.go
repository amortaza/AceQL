package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
func PostSchemaField(c echo.Context) error {
	LOG_SOURCE := "REST.PostSchemaField()"

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

	fieldTypeAsString := (*m)["type"].(string)
	fieldLabel := (*m)["label"].(string)

	if fieldLabel == "" {
		fieldLabel = fieldName
	}

	// never nil
	schema := stdsql.NewSchema()
	defer schema.Close()

	fieldType, err := dbschema.GetFieldTypeByName(fieldTypeAsString)
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	field := &dbschema.Field{Name: fieldName, Label: fieldLabel, Type: fieldType}

	if err := schema.CreateField(table, field, true); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	return c.JSON(200, "")
}
