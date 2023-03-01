# Auction
## Purpose
This application is one in muliple that combined completes a auction service. The auction application handles the CRUD operations of any items/auctions added to the service, and exposes this functionality through api.

## TechStack
The basic layout for this application is as follows:
    
    mongodb -> go-chi-restapi -> golang-apigateway (TODO) -> angular UI (TODO)


## Development
This application tries to its best ability to follow a test first based development. That means that you should first create a basic test that can confirm that you new functionality works as intenden. See Test section.

There main components that complete this application is:


1. Every new entity feature that needs to be added should first and foremost created in models folder. Example, I want to create a feature regarding books. Then i need to create a book entity in the models folder. The entitys in the models folder serves two purposes. 1. It must contain the entity structure. 2. It should only handle the basic operations for that entity such as fetching, adding, updating towards a db or other third party services. 

2. Router folder. This folder should only contain the entity routes and naming should be based on the entity that you are creating. Example: I want to create a book enity, then then router-file would be book-routes.

3. Every router file should have a controller in controllers folder that handles any business logic needed and call the underlying repository/entity-model. The router should call the controller.
## Testing

To run the tests simply run in main folder:
    
    go test -v


## Upcoming functionality
Auction listing: A component that lists all the auctions that are currently active. It should display the item being auctioned, the current highest bid, the time remaining for the auction, and any other relevant details.

Bidding system: A component that allows users to place bids on an auction. It should validate bids and ensure that only the highest bid is displayed.

Timer: A component that tracks the time remaining for an auction. It should display the remaining time for each auction and alert users when an auction is about to end.

Payment system: A component that allows the winning bidder to pay for the item. It should handle payment transactions securely and efficiently.

Notification system: A component that notifies users when they have been outbid or when an auction they are interested in is about to end.

Auction management system: A component that allows the auctioneer to manage the auctions. It should allow them to create new auctions, set the starting price and bidding increments, and end auctions when they have concluded.

User authentication and authorization system: A component that ensures that only authorized users can place bids, create auctions, and perform other actions on the platform.

By incorporating these components, you can create a comprehensive auction application that can handle active auctions effectively

Bid validation: Before accepting a bid, the system should validate it to ensure that it meets the minimum bid increment and the current highest bid. If a bid is lower than the minimum increment or lower than the current highest bid, the system should reject the bid and notify the bidder.

Automatic bidding: The system should also allow for automatic bidding, which allows users to enter the maximum amount they are willing to pay for an item. The system will then automatically place bids on the user's behalf until their maximum bid is reached.

Bid history: The system should keep a record of all bids placed on an auction, including the bidder's username, the amount of the bid, and the time the bid was placed. This history should be available to all users.

Outbid notifications: The system should notify users when they have been outbid on an auction, allowing them to increase their bid if they wish.

Bid cancellation: Users should be able to cancel a bid they have placed, but only if the auction is still ongoing and if their bid is not the current highest bid.

Reserve prices: The system should also support reserve prices, which are minimum prices set by the auctioneer. If the highest bid does not meet the reserve price, the auction will not be sold.

Bid increment rules: The system should support bid increment rules, which are the minimum amounts by which a bid must increase. This helps to ensure a fair and competitive bidding process.

Conflict of two bids at the same time?
Timestamps: Add a timestamp to each bid to track when it was submitted. If two bids are submitted at the same time, the bid with the earliest timestamp should be accepted, while the other bid should be rejected. You can then notify the bidder whose bid was not accepted, explaining that their bid was not accepted due to a conflict.

Bid lockout: Implement a bid lockout feature, which prevents multiple bids from being submitted at the exact same time. This feature would temporarily lock out the ability to submit a new bid for a brief period (e.g., a few seconds) after a bid is submitted to ensure that no two bids can be submitted simultaneously.

Automated bidding: In an automated bidding system, the system would automatically place the next highest bid on behalf of the bidder who was not accepted due to a conflict. This approach ensures that the auction continues without interruption.