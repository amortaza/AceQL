package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/relations"
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

	fieldType, err := relations.GetFieldTypeByName(fieldTypeAsString)
	if err != nil {
		logger.Error(err, "PostSchemaField()")
		return err
	}

	field := &relations.Field{Name: fieldName, Label: fieldLabel, Type: fieldType}

	schema.CreateField(table, field, true)

	schema.Close()

	return c.JSON(200, "")
}
