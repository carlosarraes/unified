package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/carlosarraes/unified/server/model"
	"github.com/gocolly/colly"
)

var webUrl = map[string]string{
	"busca":   "https://www.buscape.com.br/",
	"meliscr": "https://lista.mercadolivre.com.br/",
}

func Scrape(s, w string) ([]model.Product, error) {
	var products []model.Product
	var card string
	c := colly.NewCollector()

	if w == "busca" {
		card = "a[data-testid='product-card::card']"
	} else {
		card = "li.ui-search-layout__item"
	}

	c.OnHTML(card, func(e *colly.HTMLElement) {
		var link, name, price, thumbnail string
		var price64 float64

		if w == "busca" {
			link = e.Attr("href")
			name = e.ChildText("h2[data-testid='product-card::name']")
			price = e.ChildText("p[data-testid='product-card::price']")
			thumbnail = e.ChildAttr("img[alt*='Imagem de']", "src")

			price64 = convertToFloat(price)
		} else {
			link = e.ChildAttr("a.ui-search-link", "href")
			name = e.ChildText("h2.ui-search-item__title")
			price = e.ChildText("div.ui-search-price__second-line > span.price-tag > span.price-tag-amount")
			thumbnail = e.ChildAttr("img.ui-search-result-image__element", "src")

			price64 = extractPrice(price)
		}

		product := model.Product{
			Title:     name,
			Price:     price64,
			Link:      link,
			Thumbnail: thumbnail,
		}
		products = append(products, product)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	url := webUrl[w] + s

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func convertToFloat(s string) float64 {
	s = s[3:]
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Error converting: %v", err)
	}
	return f
}

func extractPrice(priceText string) float64 {
	re := regexp.MustCompile(`R\$\d+(?:\.\d{3})*`)
	prices := re.FindAllString(priceText, -1)
	price := prices[0][2:]
	price = strings.ReplaceAll(price, ".", "")

	price64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0
	}

	return price64
}
