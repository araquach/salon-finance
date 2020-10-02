package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

func apiTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s Total

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2019-01-01"
	dateTo := "2020-01-01"

	db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&s)

	s.T = s.P + s.S
	s.A = s.T / 12

	json, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiCostsByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var c CostByCat
	var d []CostByCat

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2019-01-01"
	dateTo := "2020-01-01"

	categories := GetCategories()

	for cat, _ := range categories {
		db.Table("costs").Select("sum(debit) as a").Where("category = ?", cat).Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&c)
		d = append(d, CostByCat{cat, c.A})
	}

	json, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)

}
