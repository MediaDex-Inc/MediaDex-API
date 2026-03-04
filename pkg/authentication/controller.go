package authentication

import (
	"mediadex/config"
	"net/http"
)

type AuthenticationConfig struct {
	*config.Config
}

func New(config *config.Config) *AuthenticationConfig {
	return &AuthenticationConfig{config}
}

func (config *AuthenticationConfig) PostHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *AuthenticationConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *AuthenticationConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *AuthenticationConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *AuthenticationConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
