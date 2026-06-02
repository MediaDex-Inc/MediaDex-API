package tag

import (
	"errors"
	"net/http"
)

type TagRequest struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Color  string `json:"colors"`
}

func (a *TagRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("Name is required !")
	}
	if a.Color == "" {
		return errors.New("Color is required !")
	}
	return nil
}

// Custom response type
type TagResponse struct {
	ID     uint   `json:"tag_id"`
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Color  string `json:"colors"`
}
