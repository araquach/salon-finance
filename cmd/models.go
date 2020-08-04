package main

import "time"

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