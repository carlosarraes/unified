package controller

import (
	"log"

	"github.com/carlosarraes/unified/server/model"
	"gorm.io/gorm"
)

type App struct {
	DSN string
	DB  model.Database
}

func (a *App) Connect() (*gorm.DB, error) {
	db, err := model.OpenDB(a.DSN)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&model.SearchHistory{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db, nil
}
