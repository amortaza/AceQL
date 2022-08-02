package rest

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_script"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
)

// test adapter blank

// !log
// http://localhost:8000/importset/{adapter}
func ImportSet(c echo.Context) error {
	LOG_SOURCE := "REST.ImportSet()"

	if err := confirmAccess(c); err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	importsetName := c.Param("importset_name")

	if importsetName == "" {
		err := fmt.Errorf("missing parameter \"adapter\"")
		c.String(400, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	grpc_script.GRPC_ImportSet(importsetName)

	return c.String(200, "")
}
