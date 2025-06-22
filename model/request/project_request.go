package request

import (
	"nexproject/model"

	"github.com/google/uuid"
)

type AddProjectRequest struct {
	ProjectID          uuid.UUID
	ProjectName        string
	ProjectDescription string
	Jobs               []model.Job
	SMEID              uuid.UUID
}
