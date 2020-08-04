package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplIndex *template.Template
)

func dbConn() (db *gorm.DB) {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
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

func main() {
	var err error
	var dir string

	db := dbConn()
	db.AutoMigrate(&BankData{}, &Takings{})
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

	flag.StringVar(&dir, "dir", "dist", "the directory to serve files from")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api/bankdata", apiBankData).Methods("GET")
	r.HandleFunc("/api/bankdata/{id}", apiAddCategory).Methods("PUT")
	r.HandleFunc("/api/takings/{salon}", apiTakings).Methods("GET")
	r.HandleFunc("/api/monthly/{month_year}", apiMonthlyTakings).Methods("GET")
	r.HandleFunc("/api/costscategory/{category}", apiCostsCategory).Methods("GET")

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":" + port, r)
}
