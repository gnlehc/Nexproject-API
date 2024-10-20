package request

import "github.com/google/uuid"

type UpdateApplicationStatusRequestDTO struct {
	AppID uuid.UUID
	// SMEID uuid.UUID
	JobID    uuid.UUID
	TalentID uuid.UUID
	StatusID int
}

type GetAllAppIDByJobIDRequestDTO struct {
	JobID uuid.UUID `json:"job_id"`
}

type GetAllAppIDByTalentIDRequestDTO struct {
	TalentID uuid.UUID `json:"talent_id"`
}
