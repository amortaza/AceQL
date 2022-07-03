package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/tableschema"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func PostSchemaField(c echo.Context) error {
	table := c.Param("table")
	fieldName := c.Param("field")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "REST:PostSchemaField()")
	}

	fieldTypeAsString := (*m)["type"].(string)
	fieldLabel := (*m)["label"].(string)

	if fieldLabel == "" {
		fieldLabel = fieldName
	}

	// never nil
	schema := stdsql.NewSchema()

	fieldType, err := tableschema.GetFieldTypeByName(fieldTypeAsString)
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "PostSchemaField()")
	}

	field := &tableschema.Field{Name: fieldName, Label: fieldLabel, Type: fieldType}

	if err := schema.CreateField(table, field, true); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	if err := schema.Close(); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	return c.JSON(200, "")
}
