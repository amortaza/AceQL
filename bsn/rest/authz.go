package rest

import (
	"errors"
	"github.com/amortaza/aceql/bsn/authn"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strings"
)

func confirmAccess(c echo.Context) error {
	LOG_SOURCE := "REST.confirmAccess()"

	bearer := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(bearer, "Bearer ") {
		err := errors.New("authorization header must start with Bearer")
		c.String(400, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	token := bearer[7:]

	auth := authn.NewAuthN(SECRET)

	_, err := auth.VerifyToken(token)
	if err != nil {
		c.String(401, err.Error())
		return logger.PushStackTrace(LOG_SOURCE, err)
	}

	return nil
}
