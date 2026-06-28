package model

import (
	"errors"
	"net/http"
)

type CollectionRequest struct {
	Name    string `json:"name"`
	Filters string `json:"filters"`
}

type CollectionUpdateRequest struct {
	Name    *string `json:"name"`
	Filters *string `json:"filters"`
}

func (collection *CollectionRequest) Bind(r *http.Request) error {
	if collection.Name == "" {
		return errors.New("name must not be null")
	}
	if collection.Filters == "" {
		return errors.New("filters must not be null")
	}

	return nil
}

type CollectionResponse struct {
	CollectionId uint   `json:"collection_id"`
	UserId       uint   `json:"user_id"`
	Name         string `json:"name"`
	Filters      string `json:"filters"`
}
