package main

import (
	"github.com/amortaza/aceql/bsn/rest"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	stdsql.Init("mysql", "clown:1844@/bsn")

	//debug
	//scheduler.StartScheduler()

	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORS())

	// get
	e.GET("/table/:table", rest.GetRecordsByQuery)
	e.GET("/table/:table/:id", rest.GetRecordById)

	e.GET("/schema/:table", rest.GetSchemaByTable)

	// create a table
	e.POST("/schema/:table", rest.PostSchemaTable)
	e.POST("/schema/table/:table/field/:field", rest.PostSchemaField)

	e.POST("/table/:table", rest.PostRecord)

	e.PUT("/table/:table/:id", rest.PutRecord)

	// delete field
	e.DELETE("/schema/table/:table/field/:field", rest.DeleteSchemaField)

	// delete record
	e.DELETE("/table/:table/id/:id", rest.DeleteRecordById)

	// upload CSV
	e.POST("/csv/:table", rest.PostCSV)

	// importset
	e.POST("/importset/:importset_name", rest.ImportSet)

	// get CSV
	e.GET("/csv/:table", rest.GetRecordsByQuery_CSV)

	e.Logger.Fatal(e.Start(":8000"))
}
