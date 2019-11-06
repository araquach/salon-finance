package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	tplIndex *template.Template
)

type BankData struct {
	ID					int		`json:"id"`
	TransDate		time.Time	`json:"trans_date"`
	TransType			string	`json:"trans_type"`
	AccountNumber		int		`json:"account_number"`
	TransDescription 	string	`json:"trans_description"`
	DebitAmount			float32	`json:"debit_amount"`
	CreditAmount		float32	`json:"credit_amount"`
	Balance				float32	`json:"balance"`
	Category			string	`json:"category"`
}

type JakataTakings struct {
	ID				int			`json:"id"`
	MonthYear		time.Time	`json:"month_year"`
	FLServices		float32		`json:"fl_services"`
	FLProducts		float32		`json:"fl_products"`
	Services		float32		`json:"services"`
	Products		float32		`json:"products"`
}

type PKTakings struct {
	ID				int			`json:"id"`
	MonthYear		time.Time	`json:"month_year"`
	FLServices		float32		`json:"fl_services"`
	FLProducts		float32		`json:"fl_products"`
	Services		float32		`json:"services"`
	Products		float32		`json:"products"`
}

type BaseTakings struct {
	ID				int			`json:"id"`
	MonthYear		time.Time	`json:"month_year"`
	Services		float32		`json:"services"`
	Products		float32		`json:"products"`
}

func dbConn() (db *gorm.DB) {
	dbhost     := os.Getenv("DB_HOST")
	dbport     := os.Getenv("DB_PORT")
	dbuser     := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname     := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}

func apiBankData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	bd := []BankData{}
	db.Find(&bd)
	db.Close()

	json, err := json.Marshal(bd)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiJakata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	jak := []JakataTakings{}
	db.Find(&jak)
	db.Close()

	json, err := json.Marshal(jak)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	pk := []PKTakings{}
	db.Find(&pk)
	db.Close()

	json, err := json.Marshal(pk)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func apiBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	base := []BaseTakings{}
	db.Find(&base)
	db.Close()

	json, err := json.Marshal(base)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func main() {
	var err error

	db := dbConn()

	db.AutoMigrate(&BankData{}, &JakataTakings{}, &PKTakings{}, &BaseTakings{})

	db.LogMode(true)

	db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplIndex = template.Must(template.ParseFiles(
		"views/index.gohtml"))
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api/bankdata", apiBankData).Methods("GET")
	r.HandleFunc("/api/jakata", apiJakata).Methods("GET")
	r.HandleFunc("/api/pk", apiPK).Methods("GET")
	r.HandleFunc("/api/base", apiBase).Methods("GET")


	// Styles
	assetHandler := http.FileServer(http.Dir("./dist/"))
	assetHandler = http.StripPrefix("/dist/", assetHandler)
	r.PathPrefix("/dist/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./dist/"))
	jsHandler = http.StripPrefix("/dist/", jsHandler)
	r.PathPrefix("/js/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./public/img/"))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", imageHandler))

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":" + port, r)
}
