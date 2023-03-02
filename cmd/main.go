package main

import (
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/handlers"
	"github.com/araquach/salon-finance/cmd/migrations"
	"github.com/araquach/salon-finance/cmd/routes"
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

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db.DBInit(dsn)

	migrations.LoadTakings()
	// migrations.LoadCosts()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	handlers.TplIndex = template.Must(template.ParseFiles("index.gohtml"))

	routes.Router()

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, forceSsl(&routes.R))
}

func forceSsl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
