package tag

import (
	"net/http"
)

type TagRequest struct {
}

func (a *TagRequest) Bind(r *http.Request) error {

	return nil
}

// Custom response type
type TagResponse struct {
}
