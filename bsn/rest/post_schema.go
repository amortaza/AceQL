package rest

import (
	"errors"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/tableschema"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// PostSchemaTable is where a Table is defined! We are creating a table here.
func PostSchemaTable(c echo.Context) error {
	tableName := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "REST:PostSchemaTable()")
	}

	fields := (*m)["fields"].([]interface{})
	tableLabel := (*m)["label"].(string)

	relation := makeTableSchema(tableName, tableLabel, fields)
	if relation == nil {
		c.JSON(500, "see logs")
		return errors.New("see logs")
	}

	schema := stdsql.NewSchema()

	if err := schema.CreateRelation_withFields(relation, true); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	if err := schema.Close(); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	return c.JSON(200, "")
}

func makeTableSchema(tableName string, tableLabel string, fields []interface{}) *tableschema.Table {
	relation := tableschema.NewTable(tableName)

	relation.SetLabel(tableLabel)

	relation.AddField("x_id", "ID", tableschema.String)

	for _, v := range fields {
		m := v.(map[string]interface{})

		fieldName := m["field"].(string)
		fieldLabel := m["label"].(string)

		fieldType, err := tableschema.GetFieldTypeByName(m["type"].(string))
		if err != nil {
			logger.Err(err, "REST:makeTableSchema()")
			continue
		}

		relation.AddField(fieldName, fieldLabel, fieldType)
	}

	return relation
}
