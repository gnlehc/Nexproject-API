package request

type UserRoleRequestDTO struct {
	Email string `json:"email" binding:"required,email"`
}
