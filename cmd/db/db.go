package db

import (
	"github.com/araquach/salon-finance/cmd/finance"
	"github.com/jinzhu/gorm"
	"os"
)

func DbConn() (db *gorm.DB) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate() {
	db := DbConn()
	db.AutoMigrate(&finance.Cost{}, &finance.Taking{})
}
