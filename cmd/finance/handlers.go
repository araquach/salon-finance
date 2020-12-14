package finance

import (
	"encoding/json"
	"github.com/araquach/salon-finance/cmd/db"
	"log"
	"net/http"
)

func ApiTakings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s Total

	db := db.DbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2019-10-03"
	dateTo := "2020-10-03"

	db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&s)

	s.T = s.P + s.S
	s.A = s.T / 12

	json, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiCostsByCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var d CostData
	var c CostByCat
	var e []CostByCat

	db := db.DbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2019-10-03"
	dateTo := "2020-10-03"

	categories := GetCategories()

	db.Table("costs").Select("sum(debit) as a").Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&c)
	d.T = c.A

	for cat, _ := range categories {
		db.Table("costs").Select("sum(debit) as a").Where("category = ?", cat).Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&c)

		percent := (c.A / d.T) * 100
		average := c.A / 12

		e = append(e, CostByCat{cat, c.A, percent, average})
	}
	d.C = e


	json, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)

}
