package main

import "time"

type Cost struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Date        time.Time  `json:"date" gorm:"type:date"`
	Type        string  `json:"type"`
	Account     string  `json:"account"`
	Description string  `json:"description"`
	Debit       float64 `json:"debit"`
	Balance     float64 `json:"balance"`
	Category    string  `json:"category"`
}

type Taking struct {
	ID       uint    `json:"id" gorm: "primary_key"`
	Date     time.Time  `json:"date" gorm:"type:date"`
	Name     string  `json:"name"`
	Salon    string  `json:"salon"`
	Services float64 `json:"services"`
	Products float64 `json:"products"`
}

type Totals struct {
	S float64
	P float64
	D float64
}

type BankData struct {
	ID					int			`json:"id"`
	TransDate			time.Time	`json:"date"`
	TransType			string		`json:"-"`
	AccountNumber		int			`json:"account"`
	TransDescription 	string		`json:"description"`
	DebitAmount			float32		`json:"amount"`
	CreditAmount		float32		`json:"-"`
	Balance				float32		`json:"-"`
	Category			string		`json:"category"`
}

type Takings struct {
	ID				int			`json:"id"`
	Salon			string		`json:"salon"`
	MonthYear		time.Time	`json:"month_year"`
	FLServices		float32		`json:"fl_services"`
	FLProducts		float32		`json:"fl_products"`
	Services		float32		`json:"services"`
	Products		float32		`json:"products"`
	Total			float32		`json:"total"`
}