package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

func apiCostsCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var d Totals
	var c GetCost

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	cat := "Wages"

	db.Table("costs").Select("sum(debit) as d").Where("category = ?", cat).Scan(&c)
	json, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)

}

func apiTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s GetTotal
	// var p Total
	// var d Total

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	stylist := "Natalie Sharpe"

	dateFrom := "2019-01-01"
	dateTo := "2020-01-01"

	db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("date >= ? AND date <= ?", dateFrom, dateTo).Where("name = ?", stylist).Scan(&s)

	// db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("name = ?", stylist).Scan(&s)

	//s.C = stylist
	s.T = s.P + s.S

	json, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiTotalTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s GetTotal

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

	c := []GetCostsByCat{}

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	// cat := "Wages"

	categories := GetCategories()

	keys := reflect.ValueOf(categories).MapKeys()

	for _, cat := range keys {
		db.Table("costs").Select("sum(debit) as a").Where("category = ?", cat).Scan(&c)

		json, err := json.Marshal(c)
		if err != nil {
			log.Println(err)
		}
		w.Write(json)
	}
}
