package main

import "time"

type Cost struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Date        time.Time `json:"date" gorm:"type:date"`
	Type        string    `json:"type"`
	Account     string    `json:"account"`
	Description string    `json:"description"`
	Debit       float64   `json:"debit"`
	Balance     float64   `json:"balance"`
	Category    string    `json:"category"`
}

type Taking struct {
	ID       uint    `json:"id" gorm: "primary_key"`
	Date     string  `json:"date" gorm:"type:date"`
	Name     string  `json:"name"`
	Salon    string  `json:"salon"`
	Services float64 `json:"services"`
	Products float64 `json:"products"`
}

type Total struct {
	C string
	S float64 `json:"services"`
	P float64 `json:"products"`
	D float64
}

type GetCost struct {
	C float64 `json:"cost"`
	D string `json:"category"`
}
