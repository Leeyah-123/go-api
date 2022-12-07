package model

type User struct {
	// ID       string `json:"id" validate:"isdefault,uuid"`
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,min=10"`
}
