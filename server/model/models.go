package model

import (
	"gorm.io/gorm"
)

type Database interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
}

type SearchHistory struct {
	gorm.Model
	Web           string `json:"web"`
	Category      string `json:"category"`
	SearchResults []byte `json:"searchResults"`
}

type Product struct {
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Link      string  `json:"permalink"`
	Thumbnail string  `json:"thumbnail"`
}

type SearchQuery struct {
	Web      string `json:"web"`
	Category string `json:"category"`
}
