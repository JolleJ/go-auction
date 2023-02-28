package models

import "time"

/*
BID STRUCT
Id			string - Bid id
AuctionId	string - Auction id that the bid belongs to
UserId      string - Userid of the bid placer
Amount		string - Amount bid
Time		string - Time bid was placed
*/
type Bid struct {
	Id        string    `json:"id"`
	AuctionId string    `json:"auctionId"`
	UserId    string    `json:"userId"`
	Amount    float64   `json:"amount"`
	Time      time.Time `json:"time"`
}
