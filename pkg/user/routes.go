package user

import (
	"mediadex/config"

	"github.com/go-chi/chi/v5"
)

// Routes the user table
func Routes(config *config.Config) chi.Router {

	// Init Router
	userConfig := New(config)
	router := chi.NewRouter()

	// Routes
	router.Get("/me", userConfig.GetMeHandler)
	router.Patch("/me", userConfig.UpdateHandler)
	router.Delete("/me", userConfig.DeleteHandler)

	return router
}
