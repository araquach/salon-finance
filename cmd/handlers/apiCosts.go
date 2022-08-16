package handlers

import (
	"encoding/json"
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/helpers"
	"github.com/araquach/salon-finance/cmd/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ApiCostsByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t float32

	type Result struct {
		Category string  `json:"category"`
		Total    float32 `json:"total"`
		Percent  float32 `json:"percent"`
		Average  float32 `json:"average"`
	}

	type Data struct {
		Salon      string   `json:"salon"`
		GrandTotal float32  `json:"grand_total"`
		ByYear     float32  `json:"by_year"`
		Months     int      `json:"months"`
		Figures    []Result `json:"figures"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	var acc string

	switch s {
	case "jakata":
		acc = "06517160"
	case "pk":
		acc = "02017546"
	case "base":
		acc = "17623364"
	}

	startDate, err := time.Parse("2006-01-02", sd)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("2006-01-02", ed)
	if err != nil {
		panic(err)
	}

	mnths := helpers.MonthsCount(startDate, endDate)

	var res []Result

	if s == "all" {
		db.DB.Model(&models.Cost{}).Order("total desc").Select("category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("category").Find(&res)
	} else {
		db.DB.Model(&models.Cost{}).Order("total desc").Select("account, category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("account", acc).Group("account, category").Find(&res)
	}

	// Calculate total costs
	for _, r := range res {
		t += r.Total
	}

	for k, v := range res {
		//remove Izzys Wage and loans from drawings
		if v.Category == "drawings" {
			(res)[k].Total = (res)[k].Total - ((2000 + 450.23 + 291) * float32(mnths))
		}
		if v.Category == "loans" {
			(res)[k].Total = (res)[k].Total + ((450 + 291) * float32(mnths))
		}
		if v.Category == "wages" {
			(res)[k].Total = (res)[k].Total + (2000 * float32(mnths))
		}
		(res)[k].Average = (res)[k].Total / float32(mnths)
		(res)[k].Percent = ((res)[k].Total / t) * 100
	}

	f := Data{
		Salon:      s,
		GrandTotal: t,
		ByYear:     t / float32(mnths) * 12,
		Months:     mnths,
		Figures:    res,
	}

	json, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

// For Meeting - excludes drawings
func ApiCostsByCatMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t float32

	type Result struct {
		Category string  `json:"category"`
		Total    float32 `json:"total"`
		Percent  float32 `json:"percent"`
		Average  float32 `json:"average"`
	}

	type Data struct {
		Salon      string   `json:"salon"`
		GrandTotal float32  `json:"grand_total"`
		ByYear     float32  `json:"by_year"`
		Months     int      `json:"months"`
		Figures    []Result `json:"figures"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	var acc string

	switch s {
	case "jakata":
		acc = "06517160"
	case "pk":
		acc = "02017546"
	case "base":
		acc = "17623364"
	}

	startDate, err := time.Parse("2006-01-02", sd)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("2006-01-02", ed)
	if err != nil {
		panic(err)
	}

	mnths := helpers.MonthsCount(startDate, endDate)

	var res []Result

	if s == "all" {
		db.DB.Model(&models.Cost{}).Order("total desc").Select("category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("category").Find(&res)
	} else {
		db.DB.Model(&models.Cost{}).Order("total desc").Select("account, category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("account", acc).Group("account, category").Find(&res)
	}

	// Calculate total costs
	for _, r := range res {
		t += r.Total
	}

	for k, v := range res {
		//remove Izzys Wage and loans from drawings
		if v.Category == "drawings" {
			(res)[k].Total = 0
		}
		if v.Category == "loans" {
			(res)[k].Total = (res)[k].Total + ((450 + 291) * float32(mnths))
		}
		if v.Category == "wages" {
			(res)[k].Total = (res)[k].Total + (2000 * float32(mnths))
		}
		(res)[k].Average = (res)[k].Total / float32(mnths)
		(res)[k].Percent = ((res)[k].Total / t) * 100
		if v.Category == "costs" {
			(res)[k].Total = 0
		}
	}

	f := Data{
		Salon:      s,
		GrandTotal: t,
		ByYear:     t / float32(mnths) * 12,
		Months:     mnths,
		Figures:    res,
	}

	json, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiCostsByDateRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type JakataTotals struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type PKTotals struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type BaseTotals struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type TotalTotals struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type Data struct {
		JakataTotals []JakataTotals `json:"jakata"`
		PKTotals     []PKTotals     `json:"pk"`
		BaseTotals   []BaseTotals   `json:"base"`
		TotalTotals  []TotalTotals  `json:"all"`
	}

	vars := mux.Vars(r)
	sd := vars["start"]
	ed := vars["end"]

	var jakata []JakataTotals
	var pk []PKTotals
	var base []BaseTotals
	var totals []TotalTotals

	db.DB.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '06517160' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&jakata)
	db.DB.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '02017546' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&pk)
	db.DB.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '17623364' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&base)
	db.DB.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&totals)

	f := Data{
		JakataTotals: jakata,
		PKTotals:     pk,
		BaseTotals:   base,
		TotalTotals:  totals,
	}

	json, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiCostsAndTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Costs struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type Takings struct {
		Month string  `json:"month"`
		Total float32 `json:"total"`
	}

	type Result struct {
		Costs   []Costs   `json:"costs"`
		Takings []Takings `json:"takings"`
	}

	vars := mux.Vars(r)
	sd := vars["start"]
	ed := vars["end"]

	var c []Costs
	var t []Takings
	var res Result

	db.DB.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&c)
	db.DB.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&t)

	res = Result{
		Costs:   c,
		Takings: t,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
