package routes

import "github.com/araquach/salon-finance/cmd/handlers"

func costs() {
	R.HandleFunc("/api/costs-by-cat/{salon}/{start}/{end}", handlers.ApiCostsByCat).Methods("GET")
	R.HandleFunc("/api/costs-by-date-range/{start}/{end}", handlers.ApiCostsByDateRange).Methods("GET")
	//for meeting
	R.HandleFunc("/api/costs-by-cat-meeting/{salon}/{start}/{end}", handlers.ApiCostsByCatMeeting).Methods("GET")
	R.HandleFunc("/api/costs-and-takings/{start}/{end}", handlers.ApiCostsAndTakings).Methods("GET")
}
