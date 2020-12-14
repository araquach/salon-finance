package main

import (
	"flag"
	"github.com/araquach/salon-finance/cmd/finance"
	"github.com/araquach/salon-finance/cmd/stock"
	"github.com/gorilla/mux"
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

func main() {
	var err error
	var dir string

	finance.LoadCosts()
	finance.LoadTakings()
	stock.LoadProfessional()
	stock.LoadRetail()

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
	// API routes
	r.HandleFunc("/api/takings", finance.ApiTakings).Methods("GET")
	r.HandleFunc("/api/costs", finance.ApiCostsByCat).Methods("GET")
	// Main Routes
	r.HandleFunc("/{category}/{name}", index)
	r.HandleFunc("/{name}", index)
	r.HandleFunc("/", index).Methods("GET")

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, r)
}
