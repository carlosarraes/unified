package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosarraes/unified/server/model"
	"github.com/carlosarraes/unified/server/utils"
)

func (a *App) search(w http.ResponseWriter, r *http.Request) {
	var searchQuery model.SearchQuery
	json.NewDecoder(r.Body).Decode(&searchQuery)

	if searchQuery.Web == "" || searchQuery.Category == "" {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var search model.SearchHistory
	checkDb := a.DB.Where("web = ? AND category = ?", searchQuery.Web, searchQuery.Category).First(&search)

	if checkDb.RowsAffected > 0 {
		fromDb := getFromDb(search)
		log.Println("from db")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fromDb)
	} else {
		products, err := getProductsFromMeliOrBuscape(searchQuery)
		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, "Error getting products")
			return
		}

		if err := saveToDb(a, products, searchQuery.Web, searchQuery.Category); err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, "Error saving to db")
			log.Printf("Error saving to db: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
