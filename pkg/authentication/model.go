package authentication

import (
	"net/http"
)

type AuthenticationRequest struct {
}

func (a *AuthenticationRequest) Bind(r *http.Request) error {

	return nil
}

// Custom response type
type AuthenticationResponse struct {
}
