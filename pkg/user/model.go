package user

import (
	"net/http"
)

type UserRequest struct {
}

func (a *UserRequest) Bind(r *http.Request) error {

	return nil
}

// Custom response type
type UserResponse struct {
}
