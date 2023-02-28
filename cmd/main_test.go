package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	ITEM STRUCT

	id 						- Id of auction
	itemId 					- Id of item being auctioned
	latestBid (timestamp)	- Timestamp of latestBid
	startPrice	float64		- Price where bidding starts at
	startTime
	endTime
	Status					- Auctionstatus
	Bids 					- List of bids on the auction

	type AuctionStatus string

	const (
		AuctionStatusActive   AuctionStatus = "active"
		AuctionStatusEnded    AuctionStatus = "ended"
	)


	BID STRUCT
	id
	auctionId
	userId
	amount
	time

*/

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//Should return all active auctions
func TestlistActiveAuctions(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/auctions/", nil)

	respose := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, respose.Code)
}
