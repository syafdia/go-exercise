package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

type Event interface {
	AggregateID() ID
	OccuredAt() time.Time
}
