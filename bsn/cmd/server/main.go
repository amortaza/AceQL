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

	// create a table
	e.POST("/schema/:table", rest.PostSchema)
	e.POST("/table/:table", rest.PostRecord)

	e.PUT("/table/:table/:id", rest.PutRecord)

	e.Logger.Fatal(e.Start(":8000"))
}

