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

func ApiStylistTakingsMonthByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type result struct {
		Month    string  `json:"month"`
		Products float32 `json:"products"`
		Services float32 `json:"services"`
		Total    float32 `json:"total"`
	}

	vars := mux.Vars(r)
	start := vars["start"]
	end := vars["end"]
	st := vars["stylist"]

	sd, err := time.Parse("2006-01-02", start)
	if err != nil {
		panic(err)
	}
	ed, err := time.Parse("2006-01-02", end)
	if err != nil {
		panic(err)
	}

	var res []result

	db.DB.Raw("SELECT DATE_TRUNC('month',date) AS  month, SUM(services) AS services, SUM(products) AS products, SUM(services) + SUM(products) AS total FROM takings WHERE date >= ? AND date <= ? AND name ILIKE ? GROUP BY month ORDER BY month", sd, ed, st+"%").Scan(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiTakingsByStylist(w http.ResponseWriter, r *http.Request) {
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
		db.DB.Model(&models.Taking{}).Select("name as stylist, salon as salon, sum(services) as services, sum(products) as products, sum(services) + sum(products) as total").Where("date BETWEEN ? AND ?", sd, ed).Group("name, salon").Find(&res)
	} else {
		db.DB.Model(&models.Taking{}).Select("name as stylist, salon as salon, sum(services) as services, sum(products) as products, sum(services) + sum(products) as total").Where("date BETWEEN ? AND ?", sd, ed).Where("salon", s).Group("name, salon").Find(&res)
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func ApiTakingsByDateRange(w http.ResponseWriter, r *http.Request) {
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

	mnths := helpers.MonthsCount(startDate, endDate)

	var res []Result
	var gt GrandTotals

	if s == "all" {
		db.DB.Raw("SELECT DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month_total ORDER BY month_total", sd, ed).Scan(&res)
	} else {
		db.DB.Raw("SELECT salon, DATE_TRUNC('month', date) AS month_total, sum(services) AS services, sum(products) as products, sum(services) + sum(products) as total FROM takings WHERE salon = ? AND date BETWEEN ? AND ? GROUP BY salon, month_total ORDER BY month_total", s, sd, ed).Scan(&res)
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

func ApiTotalsByDateRange(w http.ResponseWriter, r *http.Request) {
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

	db.DB.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'jakata' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&jakata)
	db.DB.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'pk' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&pk)
	db.DB.Raw("SELECT salon, DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE salon = 'base' AND date BETWEEN ? AND ? GROUP BY salon, month ORDER BY month", sd, ed).Scan(&base)
	db.DB.Raw("SELECT DATE_TRUNC('month', date) AS  month, sum(services) + sum(products) AS total FROM takings WHERE date BETWEEN ? AND ? GROUP BY month ORDER BY month", sd, ed).Scan(&totals)

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
