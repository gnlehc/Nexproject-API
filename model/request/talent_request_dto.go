package request

type TalentLoginRequestDTO struct {
	Email    string
	Password string
}

type TalentRegisterRequestDTO struct {
	Email       string
	Password    string
	FullName    string
	PhoneNumber string
}
