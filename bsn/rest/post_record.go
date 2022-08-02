package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
func PostRecord(c echo.Context) error {
	LOG_SOURCE := "REST.PostRecord()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	name := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	id, err := createRecord(name, m)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	return c.JSON(200, id)
}

// !log
func createRecord(name string, m *echo.Map) (string, error) {
	LOG_SOURCE := "REST.createRecord()"

	crud := stdsql.NewCRUD()

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)
	}

	rec := flux.NewRecord(tableschema, crud)
	defer rec.Close()

	for fieldname, value := range *m {
		valueAsString := value.(string)

		if err := rec.Set(fieldname, valueAsString); err != nil {
			return "", logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	id, err := rec.Insert()
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)
	}

	return id, nil
}
