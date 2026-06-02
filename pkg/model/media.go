package model

import (
	"errors"
	"net/http"
	"slices"
	t "time"
)

type MediaRequest struct {
	UserId         uint    `json:"user_id"`
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

func (media *MediaRequest) Bind(r *http.Request) error {
	status := []string{"Planned", "In Progress", "Paused", "Completed", "Abandoned", "For Later"}
	types := []string{"Film", "Shows", "Games", "Books"}

	if media.UserId == 0 {
		return errors.New("user_id must not be null")
	}
	if media.Name == "" {
		return errors.New("name must not be null")
	}
	if !slices.Contains(status, media.Status) {
		return errors.New("status is invalid")
	}
	if !slices.Contains(types, media.MediaType) {
		return errors.New("media_type is invalid")
	}

	return nil
}

type MediaResponse struct {
	ID             uint    `json:"media_id"`
	UserId         uint    `json:"user_id"`
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