package dto

type CreateProvUserRequestDTO struct {
	ProviderID  uint   `json:"provider_id"`
	PhoneNumber string `json:"phone_number"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type CreateProvUserResponseDTO struct {
	ID uint `json:"id"`
}
