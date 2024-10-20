package response

import "loom/model"

type GetAllJobResponseDTO struct {
	Jobs         []model.Job
	BaseResponse BaseResponseDTO
}

type GetJobResponseDTO struct {
	Data         model.Job
	BaseResponse BaseResponseDTO
}
