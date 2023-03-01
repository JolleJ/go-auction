package models

import (
	"errors"
	"time"
)

/*
	ITEM STRUCT
	Id			string 	    	- Id of auction
	itemId 		string			- Id of item being auctioned
	startPrice	float64			- Price where bidding starts at
	startTime	timestamp		- Starttime of auction
	endTime		timestamp		- Endtime of auction
	Status		AuctionStatus	- Status of the auction
	Bids 		[]Bid			- List of bids on the auction



*/

var auctions []Auction

type AuctionStatus string

const (
	AuctionStatusActive AuctionStatus = "active"
	AuctionStatusEnded  AuctionStatus = "ended"
)

type Auction struct {
	Id         string        `json:"id"`
	ItemId     string        `json:"itemId"`
	StartPrice float64       `json:"startPrice"`
	StartTime  time.Time     `json:"startTime"`
	EndTime    time.Time     `json:"endTime"`
	Status     AuctionStatus `json:"auctionStatus"`
	Bids       []Bid         `josn:"bids"`
}

func init() {
	auctions = append(auctions, Auction{Id: "1", ItemId: "1", StartPrice: 100, StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Status: AuctionStatusActive})
}

func ListAuctions() []Auction {
	return auctions
}

func GetAuction(id string) Auction {

	for _, auction := range auctions {
		if auction.Id == id {
			return auction
		}
	}

	return Auction{}
}

func AddAuction(auction Auction) {
	auctions = append(auctions, auction)
}

func AddBidd(bid Bid, itemId string) (bool, error) {
	for _, auction := range auctions {
		if auction.Id == itemId {

			//Validate that the bid is higher than the current one
			//If current bid is higher or if a newer bid has been placed before the request has been processed, return error.
			if auction.Bids[len(auction.Bids)-1].Amount >= bid.Amount || auction.Bids[len(auction.Bids)-1].Time.After(bid.Time) {
				return false, errors.New("Bid is not valid, too low amount of newer bid has been placed.")
			}
			auction.Bids = append(auction.Bids, bid)
		}
	}

}
