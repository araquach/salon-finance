package routes

import (
	"flag"
	"github.com/araquach/salon-finance/cmd/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

var R mux.Router

func Router() {
	var dir string

	flag.StringVar(&dir, "dir", "dist", "the directory to serve files from")
	flag.Parse()

	R.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(dir))))
	// API routes
	R.HandleFunc("/api/team-members", handlers.ApiGetStylists).Methods("GET")

	costs()
	takings()

	// Main Routes
	R.HandleFunc("/{category}/{name}", handlers.Index)
	R.HandleFunc("/{name}", handlers.Index)
	R.HandleFunc("/", handlers.Index).Methods("GET")

	return
}
