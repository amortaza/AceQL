package authn

import (
	"errors"
	"fmt"
	"github.com/amortaza/aceql/logger"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var err_ExpiredToken = errors.New("token expired")
var err_InvalidToken = errors.New("invalid token")

type AuthN struct {
	secretKey string
}

func NewAuthN(secretKey string) *AuthN {
	return &AuthN{secretKey: secretKey}
}

func (authn *AuthN) CreateToken(username string, duration time.Duration) (string, error) {
	LOG_SOURCE := "AuthN.CreateToken()"

	payload, err := newPayload(username, duration)
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)
	}

	token, err := authn.createToken(payload)
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)
	}

	logger.Info(fmt.Sprintf("token created for \"%s\"", payload.Username), LOG_SOURCE)

	return token, nil
}

func (authn *AuthN) RefreshToken(oldToken string, duration time.Duration) (string, error) {
	LOG_SOURCE := "AuthN.RefreshToken()"

	payload, err := authn.VerifyToken(oldToken)
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)
	}

	payload.IssuedAt = time.Now()
	payload.ExpiresAt = time.Now().Add(duration)

	newToken, err := authn.createToken(payload)
	if err != nil {
		return "", logger.PushStackTrace(LOG_SOURCE, err)

	}

	logger.Info(fmt.Sprintf("token refreshed for \"%s\"", payload.Username), LOG_SOURCE)

	return newToken, nil
}

func (authn *AuthN) createToken(payload *Payload) (string, error) {
	LOG_SOURCE := "AuthN.createToken()"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(authn.secretKey))
	if err != nil {
		return "", logger.Err(err, LOG_SOURCE)
	}

	return tokenString, nil
}

func (authn *AuthN) VerifyToken(token string) (*Payload, error) {
	LOG_SOURCE := "AuthN.VerifyToken()"
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, logger.Err(err_InvalidToken, LOG_SOURCE)
		}

		return []byte(authn.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, err_ExpiredToken) {
			return nil, logger.Err(err_ExpiredToken, LOG_SOURCE)
		}

		return nil, logger.Err(err_InvalidToken, LOG_SOURCE)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, logger.Err(err_InvalidToken, LOG_SOURCE)
	}

	return payload, nil
}
