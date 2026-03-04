package user

import (
	"mediadex/config"
	"net/http"
)

type UserConfig struct {
	*config.Config
}

func New(config *config.Config) *UserConfig {
	return &UserConfig{config}
}

func (config *UserConfig) PostHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *UserConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *UserConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *UserConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *UserConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
