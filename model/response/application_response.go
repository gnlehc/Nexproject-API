package response

import (
	"github.com/google/uuid"
)

type GetAllApplicationByJobIDResponseDTO struct {
	ListAppID    []uuid.UUID
	BaseResponse BaseResponseDTO
}
