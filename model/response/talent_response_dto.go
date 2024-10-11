package response

import "github.com/google/uuid"

type TalentLoginResponseDTO struct {
	StatusCode int
	Message    string
	Token      string
	Data       TalentLoginData
}

type TalentLoginData struct {
	TalentID uuid.UUID
	Email    string
}
