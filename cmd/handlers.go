package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

func apiTakingsByStylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		Stylist  string
		Products float32
		Services float32
		Total    float32
	}

	vars := mux.Vars(r)
	sd := vars["start"]
	ed := vars["end"]

	var res []result

	db.Model(&Taking{}).Select("name as stylist, sum(services) as services, sum(products) as products, sum(services) + sum(products) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("name").Find(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiCostsByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		Category string
		Total    float32
	}

	vars := mux.Vars(r)
	sd := vars["start"]
	ed := vars["end"]

	var res []result

	db.Model(&Cost{}).Select("category, sum(debit) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("category").Find(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiCostsByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		MonthTotal string
		Total      float32
	}

	var res []result

	db.Raw("SELECT DATE_TRUNC('month', date) AS  month_total, sum(debit) AS total FROM costs GROUP BY month_total").Scan(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiTakingsByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		MonthTotal string
		Services   float32
		Products   float32
		Total      float32
	}

	var res []result

	db.Raw("SELECT DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings GROUP BY month_total").Scan(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
