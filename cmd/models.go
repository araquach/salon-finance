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
	ID       uint      `json:"id" gorm:"primary_key"`
	Date     time.Time `json:"date" gorm:"type:date"`
	Name     string    `json:"name"`
	Salon    string    `json:"salon"`
	Services float64   `json:"services"`
	Products float64   `json:"products"`
}

type Total struct {
	S float64 `json:"services"`
	P float64 `json:"products"`
	T float64 `json:"total"`
	A float64 `json:"average"`
	
}

type CostByCat struct {
	C string `json:"category"`
	A float64 `json:"amount"`
}
