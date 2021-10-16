package main

import (
	"github.com/amortaza/aceql/bsn/rest"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	stdsql.Init( "mysql", "clown:1844@/bsn")

	e := echo.New()

	e.Use(middleware.CORS())

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

	e.Logger.Fatal(e.Start(":8000"))
}

