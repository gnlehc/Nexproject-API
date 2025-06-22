package request

import (
	"nexproject/model"
	"time"

	"github.com/google/uuid"
)

type JobRequestDTO struct {
	SMEID          uuid.UUID
	JobTitle       string
	JobDescription string
	JobType        string
	Qualification  string
	JobArrangement string
	Wage           string // optional
	Active         bool
	CreatedAt      time.Time
	Location       string
	Skills         []model.Skill
	ProjectID      uuid.UUID
}

type ApplyJobRequest struct {
	TalentID  uuid.UUID
	ProjectID uuid.UUID
	JobID     uuid.UUID
	Status    string
}
