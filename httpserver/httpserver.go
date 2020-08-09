package httpserver

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/rbonnat/codecademy/configuration"
	"github.com/rbonnat/codecademy/httpserver/controller"
	"github.com/rbonnat/codecademy/httpserver/middleware"
)

// Run Initialize router and launch http server
func Run(ctx context.Context, cfg *configuration.Configuration) error {
	r := chi.NewRouter()

	// Middlewares
	r.Use(jwtauth.Verifier(cfg.TokenAuth))
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(jwtauth.Authenticator)
	r.Use(middleware.Authorize)

	// Initialize routes
	r.Route("/cat/picture", func(r chi.Router) {
		// Upload a cat picture
		r.Post("/", controller.HandleInsertPic(cfg.FileStore, cfg.DBStore))

		r.Route("/{ID}", func(r chi.Router) {
			// Fetch a particular cat picture by its ID
			r.Get("/", controller.HandleGetPic(cfg.FileStore, cfg.DBStore))
			// Delete a cat picture
			r.Delete("/", controller.HandleDeletePic(cfg.FileStore, cfg.DBStore))
			// Update a previously uploaded cat picture
			r.Put("/", controller.HandleUpdatePic(cfg.FileStore, cfg.DBStore))
		})
	})

	//Fetch a list of the uploaded cat pictures
	r.Get("/cat/pictures", controller.HandleGetPics(cfg.FileStore, cfg.DBStore))

	// Launch http server
	log.Printf("Server listening on port: '%s'", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Printf("Cannot launch http server: '%v'", err)
	}

	return err
}
