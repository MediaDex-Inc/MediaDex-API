package media

import (
	"errors"
	"net/http"
	"slices"
	t "time"
)

type MediaRequest struct {
	userId         int     `json:"user_id"`
	name           string  `json:"name"`
	status         string  `json:"status"`
	mediaType      string  `json:"media_type"`
	imgURL         *string `json:"img_url"`
	rating         *int    `json:"rating"`
	notes          *string `json:"notes"`
	description    *string `json:"description"`
	genre          *string `json:"genre"`
	startDate      *t.Time `json:"start_date"`
	completionDate *t.Time `json:"completion_date"`
}

func (a *MediaRequest) Bind(r *http.Request) error {

	status := []string{"Planned", "In Progress", "Paused", "Completed", "Abandoned", "For Later"}
	types := []string{"Film", "Shows", "Games", "Books"}

	if &a.userId == nil || a.userId == 0 {
		return errors.New("No valid user id has been submited !")
	}
	if &a.name == nil || a.name == "" {
		return errors.New("No valid name has been submited !")
	}
	if &a.status == nil || slices.Contains(status, a.status) {
		return errors.New("No valid status has been submited !")
	}
	if &a.mediaType == nil || slices.Contains(types, a.mediaType) {
		return errors.New("No valid media type has benn submited !")
	}

	return nil
}

// Custom response type
type MediaResponse struct {
	userId         int     `json:"user_id"`
	name           string  `json:"name"`
	status         string  `json:"status"`
	mediaType      string  `json:"media_type"`
	imgURL         *string `json:"img_url"`
	rating         *int    `json:"rating"`
	notes          *string `json:"notes"`
	description    *string `json:"description"`
	genre          *string `json:"genre"`
	startDate      *t.Time `json:"start_date"`
	completionDate *t.Time `json:"completion_date"`
}
