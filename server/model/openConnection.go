package model

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	dsn = fmt.Sprintf("%s&parseTime=true", dsn)
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
