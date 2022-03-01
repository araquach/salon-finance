package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	rand.Seed(time.Now().UnixNano())

	v := string(rand.Intn(30))

	if err := tplIndex.Execute(w, v); err != nil {
		panic(err)
	}
}

func apiTakingsByStylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		Stylist  string  `json:"stylist"`
		Salon    string  `json:"salon"`
		Products float32 `json:"products"`
		Services float32 `json:"services"`
		Total    float32 `json:"total"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	var res []result

	if s == "all" {
		db.Model(&Taking{}).Select("name as stylist, salon as salon, sum(services) as services, sum(products) as products, sum(services) + sum(products) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("name, salon").Find(&res)
	} else {
		db.Model(&Taking{}).Select("name as stylist, salon as salon, sum(services) as services, sum(products) as products, sum(services) + sum(products) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("salon", s).Group("name, salon").Find(&res)
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiTakingsByDateRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Result struct {
		MonthTotal string  `json:"month"`
		Services   float32 `json:"services"`
		Products   float32 `json:"products"`
		Total      float32 `json:"total"`
	}

	type GrandTotals struct {
		Services   float32 `json:"services"`
		Products   float32 `json:"products"`
		GrandTotal float32 `json:"grand_total"`
		Yearly     float32 `json:"yearly"`
	}

	type Data struct {
		Salon       string      `json:"salon"`
		Months      int         `json:"months"`
		Figures     []Result    `json:"figures"`
		GrandTotals GrandTotals `json:"grand_totals"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	startDate, err := time.Parse("2006-01-02", sd)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("2006-01-02", ed)
	if err != nil {
		panic(err)
	}

	mnths := monthsCount(startDate, endDate)

	var res []Result
	var gt GrandTotals

	if s == "all" {
		db.Raw("SELECT DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month_total ORDER BY month_total", sd, ed).Scan(&res)
	} else {
		db.Raw("SELECT salon, DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE salon = ? AND date BETWEEN ? AND ? GROUP BY salon, month_total ORDER BY month_total", s, sd, ed).Scan(&res)
	}

	// Calculate total income
	for _, r := range res {
		gt.Services += r.Services
		gt.Products += r.Products
		gt.GrandTotal += r.Total
		gt.Yearly = gt.GrandTotal / float32(mnths) * 12
	}

	f := Data{
		Salon:       s,
		GrandTotals: gt,
		Months:      mnths,
		Figures:     res,
	}

	json, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiTotalsByDateRange(w http.ResponseWriter, r *http.Request) {
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

	db.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'jakata' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&jakata)
	db.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'pk' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&pk)
	db.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'base' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&base)
	db.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&totals)

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

func apiCostsByCat(w http.ResponseWriter, r *http.Request) {
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

	startDate, err := time.Parse("2006-01-02", sd)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("2006-01-02", ed)
	if err != nil {
		panic(err)
	}

	mnths := monthsCount(startDate, endDate)

	var res []Result

	if s == "all" {
		db.Model(&Cost{}).Order("total desc").Select("category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("category").Find(&res)
	} else {
		db.Model(&Cost{}).Order("total desc").Select("account, category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("account", s).Group("account, category").Find(&res)
	}

	// Calculate total costs
	for _, r := range res {
		t += r.Total
	}

	for k, v := range res {
		//remove Izzys Wage and loans from drawings
		if v.Category == "drawings" {
			(res)[k].Total = (res)[k].Total - ((2000 + 450.23 + 300) * float32(mnths))
		}
		if v.Category == "loans" {
			(res)[k].Total = (res)[k].Total + ((450 + 300) * float32(mnths))
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

func apiCostsByDateRange(w http.ResponseWriter, r *http.Request) {
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

	db.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '06517160' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&jakata)
	db.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '02017546' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&pk)
	db.Raw("SELECT account as salon, DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE account = '17623364' AND date BETWEEN ? AND ? GROUP BY account, month ORDER BY month", sd, ed).Scan(&base)
	db.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(debit) AS total FROM costs WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&totals)

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
