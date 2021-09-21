package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
)

func PutRecord(c echo.Context) error {

	name := c.Param("table")
	id := c.Param("id")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	_ = updateRecord(name, id, m)

	return c.JSON(200, "wow")
}

func updateRecord(name string, id string, m *echo.Map) error {
	rec := flux.NewRecord(name, stdsql.NewCRUD())
	_ = rec.AddPrimaryKey(id)
	_ = rec.Query()

	b, _ := rec.Next()

	if !b {
		return nil
	}

	for key, value := range *m {
		rec.Set(key, value)
	}

	return rec.Update()
}
