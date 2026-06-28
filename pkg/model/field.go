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
	FieldId uint   `json:"field_id"`
	UserId  uint   `json:"user_id"`
	Name    string `json:"name"`
}

// FieldValueResponse is used when fetching fields with their associated value for a given media.
type FieldValueResponse struct {
	FieldID uint   `json:"field_id"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
