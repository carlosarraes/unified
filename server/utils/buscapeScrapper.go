package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/carlosarraes/unified/server/model"
	"github.com/gocolly/colly"
)

func ScrapeBuscape(s string) ([]model.Product, error) {
	var products []model.Product
	c := colly.NewCollector()

	c.OnHTML("a[data-testid='product-card::card']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		name := e.ChildText("h2[data-testid='product-card::name']")
		price := e.ChildText("p[data-testid='product-card::price']")
		thumbnail := e.ChildAttr("img[alt*='Imagem de']", "src")

		price = price[3:]
		price = strings.ReplaceAll(price, ".", "")
		price = strings.ReplaceAll(price, ",", ".")
		price64, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Printf("Error converting: %v", err)
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

	url := fmt.Sprintf("https://www.buscape.com.br/%s", s)

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return products, nil
}
