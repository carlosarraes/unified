package controller

import (
	"log"
	"net/http"

	"github.com/carlosarraes/unified/server/model"
	"github.com/carlosarraes/unified/server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type App struct {
	DSN string
	DB  model.Data
}

func (a *App) Connect() (*gorm.DB, error) {
	db, err := model.OpenDB(a.DSN)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db, nil
}

func (a *App) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(utils.Cors())

	r.Post("/search", a.search)

	return r
}
