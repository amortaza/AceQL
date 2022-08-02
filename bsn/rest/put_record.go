package rest

import (
	"github.com/amortaza/aceql/bsn/cache"
	"github.com/amortaza/aceql/bsn/grpc_script"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// !log
func PutRecord(c echo.Context) error {
	LOG_SOURCE := "REST.PutRecord()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

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

// !log
func updateRecord(name string, id string, m *echo.Map) error {
	LOG_SOURCE := "REST.updateRecord()"

	crud := stdsql.NewCRUD()

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	rec := flux.NewRecord(tableschema, crud)
	defer rec.Close()

	if err := rec.AddPK(id); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	if _, err := rec.Query(); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	hasNext, err := rec.Next()
	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	if !hasNext {
		return nil
	}

	for fieldname, value := range *m {
		// for now assume everything is string
		if err := rec.Set(fieldname, value.(string)); err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}

		// todo make front end aware of field types
	}

	if err := rec.Update(); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	if err := onAfterUpdate(rec); err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	return nil
}

// !log
func onAfterUpdate(rec *flux.Record) error {
	// to do undo
	// logger.Error("onAfterUpdate is commented out", "REST:onAfterUpdate")
	LOG_SOURCE := "REST.onAfterUpdate.onAfterUpdate"

	scriptnames, err := cache.GetOnAfterUpdate_ScriptNames(rec.GetTable())
	if err != nil {
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	grpcMap := rec.GetMapGRPC()

	for _, script := range scriptnames {
		logger.Info("calling script "+script, "REST:onAfterUpdate")
		x_id, err := rec.Get("x_id")
		if err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}

		if err := grpc_script.GRPC_CallBusinessRule("../js/businessrules/", script, rec.GetTable(), x_id, grpcMap, grpcMap); err != nil {
			return logger.PushStackTrace(LOG_SOURCE, err)
		}
	}

	return nil
}
