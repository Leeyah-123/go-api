package model

type Login struct {
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,min=10"`
}
