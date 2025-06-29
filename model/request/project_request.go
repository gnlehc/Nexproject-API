package request

import (
	"github.com/google/uuid"
)

type AddProjectRequest struct {
	ProjectID          uuid.UUID
	ProjectName        string
	ProjectDescription string
	Jobs               []AddJobRequestDTO
	SMEID              uuid.UUID
}
