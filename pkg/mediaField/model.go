package mediaField

import (
	"errors"
	"net/http"
)

type MediaFieldRequest struct {
	FieldID uint   `json:"field_id"`
	MediaID uint   `json:"media_id"`
	Value   string `json:"value"`
}

func (a *MediaFieldRequest) Bind(r *http.Request) error {

	if a.FieldID <= 0 {
		return errors.New("No valid field_id has been submited !")
	}
	if a.MediaID <= 0 {
		return errors.New("No valid media_id has been submited !")
	}
	if a.Value == "" {
		return errors.New("No valid value has been submited !")
	}

	return nil
}

// Custom response type
type MediaFieldResponse struct {
	FieldID uint   `json:"field_id"`
	MediaID uint   `json:"media_id"`
	Value   string `json:"value"`
}
