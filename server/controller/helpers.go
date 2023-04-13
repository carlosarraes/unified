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

func getProductsFromMeliOrBuscape(m model.SearchQuery) ([]model.Product, error) {
	if m.Web == "busca" || m.Web == "meliscr" {
		products, err := utils.Scrape(m.Category, m.Web)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
		}

		return products, nil
	} else {
		resp, err := http.Get(mURL + meliCategoryCodes[m.Category])
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
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

		return products, nil
	}
}
