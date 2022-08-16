package routes

import "github.com/araquach/salon-finance/cmd/handlers"

func takings() {
	R.HandleFunc("/api/takings-by-stylist/{salon}/{start}/{end}", handlers.ApiTakingsByStylist).Methods("GET")
	R.HandleFunc("/api/takings-by-date-range/{salon}/{start}/{end}", handlers.ApiTakingsByDateRange).Methods("GET")
	R.HandleFunc("/api/stylist-takings-month-by-month/{start}/{end}/{stylist}", handlers.ApiStylistTakingsMonthByMonth).Methods("GET")
	R.HandleFunc("/api/totals-by-date-range/{start}/{end}", handlers.ApiTotalsByDateRange).Methods("GET")
}
