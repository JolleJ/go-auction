package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jollej/go-auction/pgm/models"
)

func ListAllBidsOnItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auctions, _ := json.Marshal(models.ListAuctions())
	w.WriteHeader(http.StatusOK)
	w.Write(auctions)
}
