package field

import (
	"errors"
	"net/http"
)

type FieldRequest struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}

func (a *FieldRequest) Bind(r *http.Request) error {
	if a.UserId == 0 {
		return errors.New("No valid user_id has been submited !")
	}
	if a.Name == "" {
		return errors.New("No valid name has been submited !")
	}

	return nil
}

// Custom response type
type FieldResponse struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}
