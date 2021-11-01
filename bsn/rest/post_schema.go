package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/relations"
	"github.com/labstack/echo"
)

func PostSchemaTable(c echo.Context) error {
	tableName := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	fields := (*m)["fields"].([]interface{})
	tableLabel := (*m)["label"].(string)

	relation := makeSchemaObject(tableName, tableLabel, fields)

	schema := stdsql.NewSchema()

	schema.CreateRelation_withFields(relation, true)

	schema.Close()

	return c.JSON(200, "")
}

func makeSchemaObject(tableName string, tableLabel string, fields []interface{}) *relations.Relation {
	relation := relations.NewRelation(tableName)

	relation.SetLabel( tableLabel )

	for _, v := range fields {
		m := v.(map[string]interface{})

		fieldName := m["field"].(string)
		fieldLabel := m["label"].(string)

		fieldType, err := relations.GetFieldTypeByName(m["type"].(string))
		if err != nil {
			logger.Error(err, logger.Main)
			continue
		}

		relation.AddField(fieldName, fieldLabel, fieldType)
	}

	return relation
}
