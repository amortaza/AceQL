package rest

import (
	"github.com/amortaza/aceql/bsn/logger"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"strconv"
)

func PutRecord(c echo.Context) error {

	name := c.Param("table")
	id := c.Param("id")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		logger.Error(err, logger.Main)
	}

	e := updateRecord(name, id, m)

	if e != nil {
		c.JSON(500, e.Error())
		return e
	}

	c.JSON(200, nil)

	return nil
}

func updateRecord(name string, id string, m *echo.Map) error {
	crud := stdsql.NewCRUD()
	relation := flux.GetRelation(name, crud)
	rec := flux.NewRecord(relation, crud)
	_ = rec.AddPK(id)
	_, _ = rec.Query()

	b, _ := rec.Next()

	if !b {
		return nil
	}

	for key, value := range *m {
		field := relation.GetField( key )

		//fmt.Println( "ace key ", key ) // debu

		// todo make front end aware of field types
		// 10/13/2021 - daddie
		if field.IsNumber() {
			v64, err := strconv.ParseFloat(value.(string), 32)

			if err != nil {
				logger.Error("Will not set field " + key + " because cannot parse float, see " + value.(string), "rest.updateRecord()")
				continue
			}

			rec.Set(key, float32(v64) )

		} else if field.IsBool() {
			rec.Set(key, value.(string) == "true" )

		} else {
			rec.Set(key, value)
		}
	}

	err := rec.Update()

	rec.Close()

	return err
}
