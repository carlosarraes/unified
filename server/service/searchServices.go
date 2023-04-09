package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/carlosarraes/unified/server/model"
)

func GetFromDb(db model.Database, web, category string) ([]model.Product, error) {
	var search model.SearchHistory
	checkDb := db.Where("web = ? AND category = ?", web, category).First(&search)
	if checkDb.RowsAffected == 0 {
		return nil, errors.New("no results found")
	}

	var products []model.Product
	if err := json.Unmarshal(search.SearchResults, &products); err != nil {
		log.Printf("Error Unmarshal: %v", err)
	}

	return products, nil
}

func SaveToDb(db model.Database, p []model.Product, web, category string) error {
	productsBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	searchHistory := model.SearchHistory{
		Web:           web,
		Category:      category,
		SearchResults: productsBytes,
	}

	result := db.Create(&searchHistory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
