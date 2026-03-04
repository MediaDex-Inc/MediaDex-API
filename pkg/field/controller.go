package field

import (
	"mediadex/config"
	"net/http"
)

type FieldConfig struct {
	*config.Config
}

func New(config *config.Config) *FieldConfig {
	return &FieldConfig{config}
}

func (config *FieldConfig) PostHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *FieldConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *FieldConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *FieldConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *FieldConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
