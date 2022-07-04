package rest

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/dbschema"
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

	var tableschema *dbschema.Table
	var err error
	tableschema, err = makeTableSchema(tableName, tableLabel, fields)
	if err != nil {
		c.JSON(500, err.Error())
		return err
	}

	schema := stdsql.NewSchema()
	defer schema.Close()

	if err := schema.CreateRelation_withFields(tableschema, true); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	return c.JSON(200, "")
}

func makeTableSchema(tableName string, tableLabel string, fields []interface{}) (*dbschema.Table, error) {
	relation := dbschema.NewTable(tableName)

	relation.SetLabel(tableLabel)

	relation.AddField("x_id", "ID", dbschema.String)

	for _, v := range fields {
		m := v.(map[string]interface{})

		fieldName := m["field"].(string)
		fieldLabel := m["label"].(string)

		fieldType, err := dbschema.GetFieldTypeByName(m["type"].(string))
		if err != nil {
			return nil, err
		}

		relation.AddField(fieldName, fieldLabel, fieldType)
	}

	return relation, nil
}
