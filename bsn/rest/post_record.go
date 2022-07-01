package rest

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strconv"
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
	relation := flux.GetRelation(name, crud)
	rec := flux.NewRecord(relation, crud)

	for key, value := range *m {
		field := relation.GetField(key)

		//fmt.Println( "post field ", field.Name, field.SyntaxType, value ) // debug

		// todo make front end aware of field types
		// 10/13/2021 - daddie
		if field.IsNumber() {
			v64, err := strconv.ParseFloat(value.(string), 32)

			if err != nil {
				logger.Error("Will not set field "+key+" because cannot parse float, see "+value.(string), "rest.updateRecord()")
				continue
			}

			rec.Set(key, float32(v64))

		} else if field.IsBool() {
			rec.Set(key, value.(string) == "true")

		} else {
			rec.Set(key, value)
		}
	}

	id, err := rec.Insert()

	rec.Close()

	return id, err
}
