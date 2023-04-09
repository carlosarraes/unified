package controller

import (
	"net/http"

	"github.com/carlosarraes/unified/server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(utils.Cors())

	r.Post("/search", a.search)

	return r
}
