package model

import (
	"errors"
	"net/http"
)

type FieldRequest struct {
	Name string `json:"name"`
}

type FieldUpdateRequest struct {
	Name *string `json:"name"`
}

func (field *FieldRequest) Bind(r *http.Request) error {
	if field.Name == "" {
		return errors.New("name must not be null")
	}

	return nil
}

type FieldResponse struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}
