package model

type ErrorValidationResponse map[string]ErrorValidatorMeta

type ErrorValidatorMeta struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
