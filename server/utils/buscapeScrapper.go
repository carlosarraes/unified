package utils

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Link  string `json:"link"`
}

func ScrapeBuscape(s string) ([]Product, error) {
	var products []Product
	c := colly.NewCollector()

	c.OnHTML("a[data-testid='product-card::card']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		name := e.ChildText("h2[data-testid='product-card::name']")
		price := e.ChildText("p[data-testid='product-card::price']")

		product := Product{
			Name:  name,
			Price: price,
			Link:  link,
		}
		fmt.Println(product)
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
