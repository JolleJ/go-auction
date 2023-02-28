package models

import "time"

/*
	ITEM STRUCT
	Id			string 	    	- Id of auction
	itemId 		string			- Id of item being auctioned
	startPrice	float64			- Price where bidding starts at
	startTime	timestamp		- Starttime of auction
	endTime		timestamp		- Endtime of auction
	Status		AuctionStatus	- Status of the auction
	Bids 		[]Bid			- List of bids on the auction

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
