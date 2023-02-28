package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jollej/go-auction/pgm/models"
)

func ListAllAuctions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auctions, _ := json.Marshal(models.ListAuctions())
	w.WriteHeader(http.StatusOK)
	w.Write(auctions)
}

func GetAuction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	auction, _ := json.Marshal(models.GetAuction(id))
	w.WriteHeader(http.StatusOK)
	w.Write(auction)
}

func AddAuction(w http.ResponseWriter, r *http.Request) {
	var auction models.Auction
	_ = json.NewDecoder(r.Body).Decode(&auction)
	models.AddAuction(auction)
	w.WriteHeader(http.StatusOK)
}
