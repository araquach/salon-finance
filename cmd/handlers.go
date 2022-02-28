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

	type result struct {
		Salon      string  `json:"salon"`
		MonthTotal string  `json:"month"`
		Services   float32 `json:"services"`
		Products   float32 `json:"products"`
		Total      float32 `json:"total"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	var res []result
	if s == "all" {
		db.Raw("SELECT DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month_total ORDER BY month_total", sd, ed).Scan(&res)
	} else {
		db.Raw("SELECT salon, DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE salon = ? AND date BETWEEN ? AND ? GROUP BY salon, month_total ORDER BY month_total", s, sd, ed).Scan(&res)
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiCostsByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t float32

	type Result struct {
		Account  string  `json:"account"`
		Category string  `json:"category"`
		Total    float32 `json:"total"`
		Percent  float32 `json:"percent"`
		Average  float32 `json:"average"`
	}

	type Data struct {
		GrandTotal float32  `json:"grand_total"`
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
		db.Model(&Cost{}).Select("account, category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("account", s).Group("account, category").Find(&res)
	}

	// Calculate total costs
	for _, r := range res {
		t += r.Total * float32(len(res))
	}

	for k, v := range res {
		monthAv := t / float32(mnths)

		(res)[k].Average = v.Total / float32(mnths)
		(res)[k].Percent = (v.Total / monthAv) * 100
	}

	f := Data{
		GrandTotal: t,
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

	type result struct {
		Account    string  `json:"account"`
		MonthTotal string  `json:"month_total"`
		Total      float32 `json:"total"`
	}

	vars := mux.Vars(r)
	s := vars["salon"]
	sd := vars["start"]
	ed := vars["end"]

	var res []result

	if s == "all" {
		db.Raw("SELECT account, DATE_TRUNC('month', date) AS  month_total, sum(debit) AS total FROM costs WHERE date BETWEEN ? AND ? GROUP BY account, month_total", sd, ed).Scan(&res)
	} else {
		db.Raw("SELECT account, DATE_TRUNC('month', date) AS  month_total, sum(debit) AS total FROM costs WHERE account = ? AND date BETWEEN ? AND ? GROUP BY account, month_total", s, sd, ed).Scan(&res)
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
