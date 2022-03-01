package main

import "time"

type Cost struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date" gorm:"type:date"`
	Type        string    `json:"type"`
	Account     string    `json:"account"`
	Description string    `json:"description"`
	Debit       float64   `json:"debit"`
	Category    string    `json:"category"`
	SubCat      string    `json:"sub_cat"`
}

type Taking struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Date     time.Time `json:"date" gorm:"type:date"`
	Name     string    `json:"name"`
	Salon    string    `json:"salon"`
	Services float64   `json:"services"`
	Products float64   `json:"products"`
}


