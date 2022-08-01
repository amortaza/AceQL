package rest

import (
	"github.com/amortaza/aceql/bsn/cache"
	"github.com/amortaza/aceql/bsn/grpc_script"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

func PutRecord(c echo.Context) error {
	LOG_SOURCE := "REST:PutRecord"

	name := c.Param("table")
	id := c.Param("id")

	m := &echo.Map{}

	if err := c.Bind(m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	if err := updateRecord(name, id, m); err != nil {
		c.JSON(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	c.JSON(200, "")

	return nil
}

func updateRecord(name string, id string, m *echo.Map) error {
	crud := stdsql.NewCRUD()

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		return err
	}

	rec := flux.NewRecord(tableschema, crud)
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
			return err
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
	// to do undo
	// logger.Error("onAfterUpdate is commented out", "REST:onAfterUpdate")
	//LOG_SOURCE := "CORE:put_record.onAfterUpdate"

	scriptnames, err := cache.GetOnAfterUpdate_ScriptNames(rec.GetTable())
	if err != nil {
		return err
	}

	grpcMap := rec.GetMapGRPC()

	for _, script := range scriptnames {
		logger.Info("calling script "+script, "REST:onAfterUpdate")
		x_id, err := rec.Get("x_id")
		if err != nil {
			return logger.PushStackTrace("put_record.onAfterUpdate()", err)
		}

		if err := grpc_script.GRPC_CallBusinessRule("../js/businessrules/", script, rec.GetTable(), x_id, grpcMap, grpcMap); err != nil {
			return err
		}
	}

	return nil
}
