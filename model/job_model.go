package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	JobID          uuid.UUID
	JobTitle       string
	JobDescription string
	JobType        string
	CreatedAt      time.Time
	Qualification  string
	Wage           string
}
