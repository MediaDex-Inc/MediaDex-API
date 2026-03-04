package authentication

import (
	"mediadex/config"

	"github.com/go-chi/chi/v5"
)

// Routes the authentication table
func Routes(config *config.Config) chi.Router {

	// Init Router
	authenticationConfig := New(config)
	router := chi.NewRouter()

	// Routes
	router.Post("/", authenticationConfig.PostHandler)
	router.Get("/{id}", authenticationConfig.GetByIdHandler)
	router.Get("/", authenticationConfig.GetAllHandler)
	router.Patch("/{id}", authenticationConfig.UpdateHandler)
	router.Delete("/{id}", authenticationConfig.DeleteHandler)

	return router
}
