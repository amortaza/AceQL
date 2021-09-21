package utils

import (
	"encoding/hex"
	"github.com/google/uuid"
)

func NewUUID() string {
	uid := uuid.New()
	str := hex.EncodeToString(uid[:])

	return str
}
