package response

import "nexproject/model"

type GetAllProjectsResponse struct {
	Projects     []model.Project
	BaseResponse BaseResponseDTO
}

type GetProjectByID struct {
	Project      model.Project
	BaseResponse BaseResponseDTO
}
