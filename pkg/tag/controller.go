package tag

import (
	"mediadex/config"
	"net/http"
)

type TagConfig struct {
	*config.Config
}

func New(config *config.Config) *TagConfig {
	return &TagConfig{config}
}

func (config *TagConfig) PostHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *TagConfig) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *TagConfig) GetAllHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *TagConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func (config *TagConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}
