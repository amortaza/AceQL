package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/relation_type"
	"github.com/labstack/echo"
)

func PostSchema(c echo.Context) error {
	name := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	fields := (*m)["fields"].([]interface{})

	relation := makeCollectionObject(name, fields)

	schema := stdsql.NewSchema()

	schema.CreateRelation_withFields(relation, true)

	return c.JSON(200, "")
}

func makeCollectionObject(name string, fields []interface{}) *relation_type.Relation {
	relation := relation_type.NewRelation(name)

	for _, v := range fields {
		m := v.(map[string]interface{})

		fieldname := m["field"].(string)

		fieldtype, err := relation_type.GetFieldTypeByName(m["type"].(string))
		if err != nil {
			logger.Error(err, logger.Main)
			continue
		}

		relation.AddField(fieldname, fieldtype)
	}

	return relation
}
