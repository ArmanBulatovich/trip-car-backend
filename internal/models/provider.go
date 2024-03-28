package models

type Provider struct {
	ID          uint
	Name        string
	Slug        string
	Metadata    []byte
	BIN         string
	Email       string
	PhoneNumber string
}
