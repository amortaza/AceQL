package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
)

func PostRecord(c echo.Context) error {

	name := c.Param("table")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	id, _ := createRecord(name, m)

	return c.JSON(200, id)
}

func createRecord(name string, m *echo.Map) (string, error) {
	crud := stdsql.NewCRUD()
	rec := flux.NewRecord(flux.GetRelation(name, crud), crud)

	for key, value := range *m {
		rec.Set(key, value)
	}

	return rec.Insert()
}
