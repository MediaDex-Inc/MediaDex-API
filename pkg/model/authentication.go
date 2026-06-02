package model

import (
	"errors"
	"net/http"
)

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (token *TokenRequest) Bind(r *http.Request) error {
	if token.RefreshToken == "" {
		return errors.New("refresh token must not be null")
	}
	return nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}