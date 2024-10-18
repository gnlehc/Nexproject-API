package response

import (
	"loom/model"

	"github.com/google/uuid"
)

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

type GetAllTalentResponseDTO struct {
	Talents      []model.Talent
	BaseResponse BaseResponseDTO
}

type TalentDetailResponseDTO struct {
	Data         model.Talent
	BaseResponse BaseResponseDTO
}
