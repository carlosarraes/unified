package model

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data interface{}

type MySql struct {
	*gorm.DB
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
	Query    string `json:"query"`
}

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
