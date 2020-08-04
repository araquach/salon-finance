package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

//func apiBankData(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db := dbConn()
//	bd := []BankData{}
//	db.Find(&bd)
//	db.Close()
//
//	json, err := json.Marshal(bd)
//	if err != nil {
//		log.Println(err)
//	}
//	w.Write(json)
//}
//
//func apiCostsCategory(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db := dbConn()
//	c := []BankData{}
//	params := mux.Vars(r)
//	category := params["category"]
//
//	db.Where("category = ?", category).Find(&c)
//	db.Close()
//
//	json, err := json.Marshal(c)
//	if err != nil {
//		log.Println(err)
//	}
//	w.Write(json)
//}
//
//func apiAddCategory(w http.ResponseWriter, r *http.Request) {
//	var bankData BankData
//	params := mux.Vars(r)
//	id := params["id"]
//	json.NewDecoder(r.Body).Decode(&bankData)
//	db := dbConn()
//	db.Model(&bankData).Where("id =" + id).Update("category", bankData.Category)
//	db.Close()
//}
//
//func apiTakings(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db := dbConn()
//	t := []Takings{}
//
//	params := mux.Vars(r)
//	salon := params["salon"]
//
//	if (salon == "All") {
//		db.Find(&t)
//	} else {
//		db.Order("month_year").Where("salon = ?", salon).Find(&t)
//		db.Close()
//	}
//
//	json, err := json.Marshal(t)
//	if err != nil {
//		log.Println(err)
//	}
//	w.Write(json)
//}
//
//func apiMonthlyTakings(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db := dbConn()
//	t := []Takings{}
//
//	params := mux.Vars(r)
//	month := params["month_year"]
//
//	db.Where("month_year = ?", month).Find(&t)
//	db.Close()
//
//	json, err := json.Marshal(t)
//	if err != nil {
//		log.Println(err)
//	}
//	w.Write(json)
//}
