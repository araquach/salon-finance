package db

import (
	"github.com/jinzhu/gorm"
	"os"
)

func DbConn() (db *gorm.DB) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}
