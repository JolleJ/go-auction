package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jollej/go-auction/pgm/models"
)

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
func TestListActiveAuctions(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/auctions/", nil)

	respose := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, respose.Code)
}

func TestGetAuction(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/auctions/1", nil)

	respose := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, respose.Code)
}

func TestAddAuction(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	auction := &models.Auction{Id: "1", ItemId: "1", StartPrice: 100, StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Status: models.AuctionStatusActive}
	auctionJson, err := json.Marshal(auction)
	jsonRequest := bytes.NewReader(auctionJson)

	if err != nil {
		log.Panic(err)
	}
	req, _ := http.NewRequest("POST", "/auctions", jsonRequest)
	respose := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, respose.Code)
}
