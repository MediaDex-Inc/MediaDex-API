package collection

import (
	"mediadex/config"
	"net/http"
)

type CollectionConfig struct {
	*config.Config
}

func New(config *config.Config) *CollectionConfig {
	return &CollectionConfig{config}
}

func (config *CollectionConfig) PostHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *CollectionConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *CollectionConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *CollectionConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *CollectionConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
