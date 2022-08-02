package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
// PostSchemaTable is where a Table is defined! We are creating a table here.
func PostSchemaTable(c echo.Context) error {
	LOG_SOURCE := "REST.PostSchemaField()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	tableName := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	fields := (*m)["fields"].([]interface{})
	tableLabel := (*m)["label"].(string)

	var tableschema *dbschema.Table
	var err error
	tableschema, err = makeTableSchema(tableName, tableLabel, fields)
	if err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	schema := stdsql.NewSchema()
	defer schema.Close()

	if err := schema.CreateRelation_withFields(tableschema, true); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	return c.JSON(200, "")
}

// !log
func makeTableSchema(tableName string, tableLabel string, fields []interface{}) (*dbschema.Table, error) {
	LOG_SOURCE := "REST.makeTableSchema()"

	relation := dbschema.NewTable(tableName)

	relation.SetLabel(tableLabel)

	relation.AddField("x_id", "ID", dbschema.String)

	for _, v := range fields {
		m := v.(map[string]interface{})

		fieldName := m["field"].(string)
		fieldLabel := m["label"].(string)

		fieldType, err := dbschema.GetFieldTypeByName(m["type"].(string))
		if err != nil {
			return nil, logger.PushStackTrace(LOG_SOURCE, err)
		}

		relation.AddField(fieldName, fieldLabel, fieldType)
	}

	return relation, nil
}
