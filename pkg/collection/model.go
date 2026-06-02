package collection

import (
	"errors"
	"net/http"
)

type CollectionRequest struct {
	UserId  uint     `json:"user_id"`
	Name    string   `json:"name"`
	Filters []string `json:"filters"`
}

func (a *CollectionRequest) Bind(r *http.Request) error {
	if a.UserId == 0 {
		return errors.New("No valid user_id has been submited !")
	}
	if a.Name == "" {
		return errors.New("No valid name has been submited !")
	}
	if a.Filters == nil {
		return errors.New("No valid filters has been submited !")
	}

	return nil
}

// Custom response type
type CollectionResponse struct {
	UserId  uint     `json:"user_id"`
	Name    string   `json:"name"`
	Filters []string `json:"filters"`
}
