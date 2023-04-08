package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/carlosarraes/unified/server/model"
	"github.com/carlosarraes/unified/server/utils"
)

const mURL = "https://api.mercadolibre.com/sites/MLB/search?category="

var meliCategoryCodes = map[string]string{
	"tv":        "MLB1002",
	"geladeira": "MLB181294",
	"celular":   "MLB1055",
}

type MeliResponse struct {
	Results []model.Product `json:"results"`
}

func (a *App) search(w http.ResponseWriter, r *http.Request) {
	var searchQuery model.SearchQuery
	json.NewDecoder(r.Body).Decode(&searchQuery)

	if searchQuery.Web == "busca" {
		products, err := utils.ScrapeBuscape(searchQuery.Category)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	} else {
		resp, err := http.Get(mURL + meliCategoryCodes[searchQuery.Category])
		if err != nil {
			log.Printf("Error: %v", err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		var meliResponse MeliResponse
		if err = json.Unmarshal(data, &meliResponse); err != nil {
			log.Printf("Error Unmarshal: %v", err)
		}

		var products []model.Product
		for _, product := range meliResponse.Results {
			products = append(products, model.Product{
				Title:     product.Title,
				Price:     product.Price,
				Link:      product.Link,
				Thumbnail: product.Thumbnail,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)

	}
}
