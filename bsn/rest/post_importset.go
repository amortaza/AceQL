package rest

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_client"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// test adapter blank

// http://localhost:8000/importset/{adapter}
func ImportSet(c echo.Context) error {
	importsetName := c.Param("importset_name")

	if importsetName == "" {
		err := fmt.Errorf("missing parameter \"adapter\"")
		c.String(400, err.Error())
		return logger.Err(err, "???")
	}

	grpc_client.GRPC_ImportSet(importsetName)

	return c.String(200, "")
}
