package field

import (
	"net/http"
)

type FieldRequest struct {
}

func (a *FieldRequest) Bind(r *http.Request) error {

	return nil
}

// Custom response type
type FieldResponse struct {
}
