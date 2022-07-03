package rest

import (
	"errors"
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
		return logger.Err(err, "REST:PutRecord")
	}

	if err := updateRecord(name, id, m); err != nil {
		c.JSON(500, err.Error())
		return err
	}

	c.JSON(200, "")

	return nil
}

func updateRecord(name string, id string, m *echo.Map) error {
	crud := stdsql.NewCRUD()
	relation := flux.GetTableSchema(name, crud)
	if relation == nil {
		return errors.New("see logs")
	}

	rec := flux.NewRecord(relation, crud)
	if rec == nil {
		return errors.New("see logs")
	}

	defer rec.Close()

	if err := rec.AddPK(id); err != nil {
		return err
	}

	if _, err := rec.Query(); err != nil {
		return err
	}

	hasNext, err := rec.Next()
	if err != nil {
		return err
	}

	if !hasNext {
		return nil
	}

	for fieldname, value := range *m {
		// for now assume everything is string
		if err := rec.Set(fieldname, value.(string)); err != nil {
			continue
		}

		// todo make front end aware of field types
	}

	if err := rec.Update(); err != nil {
		return err
	}

	if err := onAfterUpdate(rec); err != nil {
		return err
	}

	return nil
}

func onAfterUpdate(rec *flux.Record) error {
	grpcMap := rec.GetMapGRPC()

	scriptnames := cache.GetOnAfterUpdate_ScriptNames(rec.GetTable())
	if scriptnames == nil {
		return errors.New("see logs")
	}

	for _, script := range scriptnames {
		if err := grpcclient.GRPC_OnRecordUpdate(script, grpcMap); err != nil {
			return err
		}
	}

	return nil
}
