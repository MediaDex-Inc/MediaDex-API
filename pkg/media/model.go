package media

import (
	"errors"
	"net/http"
	"slices"
	t "time"
)

type MediaRequest struct {
	UserId         int     `json:"user_id"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	MediaType      string  `json:"media_type"`
	ImgURL         *string `json:"img_url"`
	Rating         *int    `json:"rating"`
	Notes          *string `json:"notes"`
	Description    *string `json:"description"`
	Genre          *string `json:"genre"`
	StartDate      *t.Time `json:"start_date"`
	CompletionDate *t.Time `json:"completion_date"`
}

func (a *MediaRequest) Bind(r *http.Request) error {

	status := []string{"Planned", "In Progress", "Paused", "Completed", "Abandoned", "For Later"}
	types := []string{"Film", "Shows", "Games", "Books"}

	if &a.UserId == nil || a.UserId == 0 {
		return errors.New("No valid user id has been submited !")
	}
	if &a.Name == nil || a.Name == "" {
		return errors.New("No valid name has been submited !")
	}
	if &a.Status == nil || slices.Contains(status, a.Status) {
		return errors.New("No valid status has been submited !")
	}
	if &a.MediaType == nil || slices.Contains(types, a.MediaType) {
		return errors.New("No valid media type has benn submited !")
	}

	return nil
}

// Custom response type
type MediaResponse struct {
	UserId         int     `json:"user_id"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	MediaType      string  `json:"media_type"`
	ImgURL         *string `json:"img_url"`
	Rating         *int    `json:"rating"`
	Notes          *string `json:"notes"`
	Description    *string `json:"description"`
	Genre          *string `json:"genre"`
	StartDate      *t.Time `json:"start_date"`
	CompletionDate *t.Time `json:"completion_date"`
}
