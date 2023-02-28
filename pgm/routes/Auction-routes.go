package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jollej/go-auction/pgm/controllers"
)

var RegisterAuctionRoutes = func(router *chi.Mux) {
	router.Route("/auctions", func(r chi.Router) {
		r.Get("/", controllers.ListAllAuctions)
		r.Get("/{id}", controllers.GetAuction)
		r.Post("/", controllers.AddAuction)
	})

	router.Route("/auctions/{id}/bids", func(r chi.Router) {
		//List all bids
		r.Get("/", controllers.ListAllAuctions)
		//Place bid on auction
		r.Post("/", controllers.AddAuction)
	})
}
