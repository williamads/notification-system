package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	User    User      `json:"user"`
}
