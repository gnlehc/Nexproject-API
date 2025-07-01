package response

import (
	"nexproject/model"

	"github.com/google/uuid"
)

type GetAllApplicationByJobIDResponseDTO struct {
	ListAppID    []uuid.UUID
	BaseResponse BaseResponseDTO
}

type GetAllApplicationsResponseDTO struct {
	Data       []model.TrApplication
	BaseOutput BaseResponseDTO
}
