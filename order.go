package simuexos

import "time"

type OrderType int

const (
	BuyOrder OrderType = iota
	SellOrder
)

// Order represents an order created by an account.
// Orders are placed into their corresponding orderbook and can either be cancalled or settled.
// Current implementation of the exchange will close an open a new order in response to an order modification request.
// Partial fills and related data will be filled on settlement logs while referencing the order, instead of in the order struct,
// therefore lack of data such as remaining amount is not kept here per the implementation details.
type Order struct {
	ID   int
	Type OrderType

	Creator Account

	CreationDate time.Time
	ExpiryDate   time.Time

	Token     Token
	BaseToken Token

	// Price and Amount are currently typed as int, will be switched to proper types once implemented
	Price  int
	Amount int
}
