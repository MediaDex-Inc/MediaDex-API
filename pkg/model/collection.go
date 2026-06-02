package model

import (
	"errors"
	"net/http"
)

type CollectionRequest struct {
	UserId  uint           `json:"user_id"`
	Name    string         `json:"name"`
	Filters map[string]any `json:"filters"`
}

func (collection *CollectionRequest) Bind(r *http.Request) error {
	if collection.UserId == 0 {
		return errors.New("user_id must not be null")
	}
	if collection.Name == "" {
		return errors.New("name must not be null")
	}
	if collection.Filters == nil {
		return errors.New("filters must not be null")
	}

	return nil
}

type CollectionResponse struct {
	UserId  uint           `json:"user_id"`
	Name    string         `json:"name"`
	Filters map[string]any `json:"filters"`
}