package response

import (
	"nexproject/model"

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

type EditTalentDetailResponseDTO struct {
	Data         model.Talent
	BaseResponse BaseResponseDTO
}

type TalentPortfolioResponseDTO struct {
	Data         []model.Portofolio
	BaseResponse BaseResponseDTO
}
type GetTalentSkillsResponseDTO struct {
	Skills       []model.Skill
	BaseResponse BaseResponseDTO
}
