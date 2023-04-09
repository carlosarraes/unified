package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosarraes/unified/server/model"
	"github.com/carlosarraes/unified/server/services"
	"github.com/carlosarraes/unified/server/utils"
)

func (a *App) search(w http.ResponseWriter, r *http.Request) {
	var searchQuery model.SearchQuery
	json.NewDecoder(r.Body).Decode(&searchQuery)

	if searchQuery.Web == "" || searchQuery.Category == "" {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	mySql := a.DB
	fromDb, err := services.GetFromDb(mySql, searchQuery.Web, searchQuery.Category)
	if err == nil {
		log.Println("Got from db")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fromDb)
		return
	}

	products, err := getProductsFromMeliOrBuscape(searchQuery)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, "Error getting products")
		return
	}

	if err := services.SaveToDb(mySql, products, searchQuery.Web, searchQuery.Category); err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, "Error saving to db")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
