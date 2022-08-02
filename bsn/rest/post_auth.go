package rest

import (
	"errors"
	"github.com/amortaza/aceql/bsn/authn"
	"github.com/amortaza/aceql/logger"
	"github.com/labstack/echo"
	"strings"
	"time"
)

var SECRET = "\"veryimportantthings\""

// !log
// http://localhost:8000/auth
func PostAuth(c echo.Context) error {
	LOG_SOURCE := "REST.PostAuth()"

	name := c.QueryParam("user")
	pwd := c.QueryParam("password")

	if pwd == "" {
		err := errors.New("invalid password")
		c.String(401, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	auth := authn.NewAuthN(SECRET)
	token, err := auth.CreateToken(name, 1*time.Hour)
	if err != nil {
		c.String(500, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	c.Response().Header().Set("Authorization", token)

	return c.String(200, "token created, see header")
}

// !log
// http://localhost:8000/auth/refresh
func PostAuthRefresh(c echo.Context) error {
	LOG_SOURCE := "REST.PostAuthRefresh()"

	bearer := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(bearer, "Bearer ") {
		err := errors.New("authorization header must start with Bearer")
		c.String(400, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	oldToken := bearer[7:]

	auth := authn.NewAuthN(SECRET)
	newToken, err := auth.RefreshToken(oldToken, 1*time.Hour)
	if err != nil {
		c.String(401, err.Error())
		return logger.Err(err, LOG_SOURCE)
	}

	c.Response().Header().Set("Authorization", newToken)

	return c.String(200, "token refreshed, see header")
}
