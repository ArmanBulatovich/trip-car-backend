package models

type ProvUser struct {
	ID          uint
	Email       string
	PhoneNumber string
	Fullname    string
	ProviderID  uint
	Password    string
	Role        string
}
