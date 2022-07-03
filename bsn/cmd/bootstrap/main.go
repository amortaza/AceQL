package main

import (
	"github.com/amortaza/aceql/bsn/bootstrap"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
)

func main() {
	stdsql.Init("mysql", "clown:1844@/bsn")

	bootstrap.Run()
}
