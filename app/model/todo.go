package model

type Todo struct {
	Title string `json:"title" validate:"required,max=30"`
}
