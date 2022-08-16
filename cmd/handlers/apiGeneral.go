package handlers

import (
	"encoding/json"
	"github.com/araquach/salon-finance/cmd/db"
	"github.com/araquach/salon-finance/cmd/models"
	"log"
	"net/http"
)

func ApiGetStylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res []models.TeamMember

	db.DB.Find(&res)

	json, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
