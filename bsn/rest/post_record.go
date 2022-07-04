package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func PostRecord(c echo.Context) error {
	name := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, "???")
	}

	id, err := createRecord(name, m)
	if err != nil {
		c.String(500, err.Error())
		return err
	}

	return c.JSON(200, id)
}

func createRecord(name string, m *echo.Map) (string, error) {
	crud := stdsql.NewCRUD()

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		return "", err
	}

	rec := flux.NewRecord(tableschema, crud)
	defer rec.Close()

	for fieldname, value := range *m {
		valueAsString := value.(string)

		if err := rec.Set(fieldname, valueAsString); err != nil {
			return "", err
		}
	}

	id, err := rec.Insert()
	if err != nil {
		return "", err
	}

	return id, nil
}
