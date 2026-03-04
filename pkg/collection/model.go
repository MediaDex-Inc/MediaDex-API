package collection

import (
	"net/http"
)

type CollectionRequest struct {
}

func (a *CollectionRequest) Bind(r *http.Request) error {

	return nil
}

// Custom response type
type CollectionResponse struct {
}
