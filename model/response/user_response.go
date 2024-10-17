package response

type UserRoleResponseDTO struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Role       string `json:"role"`
}
