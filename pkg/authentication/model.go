package authentication

import (
	"errors"
	"net/http"
)

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (t *TokenRequest) Bind(r *http.Request) error {
	if t.RefreshToken == "" {
		return errors.New("refresh_token is required")
	}
	return nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}
