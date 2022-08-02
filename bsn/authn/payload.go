package authn

import (
	"errors"
	"github.com/amortaza/aceql/logger"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID
	Username  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return errors.New("JWT Token has expired")
	}

	return nil
}

func newPayload(username string, duration time.Duration) (*Payload, error) {
	LOG_SOURCE := "authn.NewPayload()"

	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, logger.Err(err, LOG_SOURCE)
	}

	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}
