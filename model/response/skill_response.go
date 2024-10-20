package response

import "loom/model"

type GetAllSkillResponseDTO struct {
	Skills       []model.Skill   `json:"Skills"`
	BaseResponse BaseResponseDTO `json:"BaseResponse"`
}
