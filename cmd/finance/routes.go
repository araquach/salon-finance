package finance

import "github.com/gorilla/mux"

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/api/takings", ApiTakings).Methods("GET")
	r.HandleFunc("/api/costs", ApiCostsByCat).Methods("GET")
}