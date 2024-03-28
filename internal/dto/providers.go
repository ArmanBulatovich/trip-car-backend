package dto

type CreateProviderRequestDTO struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Metadata    interface{} `json:"metadata"`
	BIN         string      `json:"bin"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phone_number"`
}

type CreateProviderResponseDTO struct {
	ID uint `json:"id"`
}
