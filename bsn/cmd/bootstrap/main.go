package main

import (
	"github.com/amortaza/aceql/bsn/bootstrap"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
)

func main() {
	stdsql.Init("mysql", "clown:1844@/bsn")

	if err := bootstrap.Run(); err != nil {
		logger.Err(err, "bootstrap.main()")
	}
}
