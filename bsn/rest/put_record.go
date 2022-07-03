package rest

import (
	"github.com/amortaza/aceql/bsn/cache"
	"github.com/amortaza/aceql/bsn/grpcclient"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func PutRecord(c echo.Context) error {
	name := c.Param("table")
	id := c.Param("id")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, logger.Main)
	}

	err2 := updateRecord(name, id, m)

	if err2 != nil {
		c.JSON(500, err2.Error())
		return logger.Err(err2, logger.Main)
	}

	c.JSON(200, nil)

	return nil
}

func updateRecord(name string, id string, m *echo.Map) error {
	crud := stdsql.NewCRUD()
	relation := flux.GetTableSchema(name, crud)
	rec := flux.NewRecord(relation, crud)
	_ = rec.AddPK(id)
	_, _ = rec.Query()

	b, _ := rec.Next()

	if !b {
		return nil
	}

	for key, value := range *m {
		// for now assume everything is string
		rec.Set(key, value.(string))

		//fmt.Println( "ace key ", key ) // debu

		// todo make front end aware of field types
		// 10/13/2021 - daddie
		/*
			field := relation.GetField( key )

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
		*/
	}

	err := rec.Update()

	rec.Close()

	onAfterUpdate(rec)

	return err
}

func onAfterUpdate(rec *flux.Record) {
	grpcMap := rec.GetMapGRPC()

	scriptnames := cache.GetOnAfterUpdate_ScriptNames(rec.GetTable())

	for _, script := range scriptnames {
		grpcclient.GRPC_OnRecordUpdate(script, grpcMap)
	}
}
