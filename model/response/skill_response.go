package response

import "nexproject/model"

type GetAllSkillResponseDTO struct {
	Skills       []model.Skill   `json:"Skills"`
	BaseResponse BaseResponseDTO `json:"BaseResponse"`
}
