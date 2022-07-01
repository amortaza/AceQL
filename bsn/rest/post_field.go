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
		logger.Error(err, logger.Main)
	}

	fieldTypeAsString := (*m)["type"].(string)
	fieldLabel := (*m)["label"].(string)

	if fieldLabel == "" {
		fieldLabel = fieldName
	}

	schema := stdsql.NewSchema()

	fieldType, err := tableschema.GetFieldTypeByName(fieldTypeAsString)
	if err != nil {
		logger.Error(err, "PostSchemaField()")
		return err
	}

	field := &tableschema.Field{Name: fieldName, Label: fieldLabel, Type: fieldType}

	schema.CreateField(table, field, true)

	schema.Close()

	return c.JSON(200, "")
}
