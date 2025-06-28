package response

import "nexproject/model"

type GetLearningResponseDTO struct {
	BaseResponseDTO
	Learnings []model.Learning
}

type GetLearningDetailResponseDTO struct {
	BaseResponseDTO
	Data model.Learning
}
