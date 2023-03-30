package simuexos

type priorityQueue interface {
	peek() Order
	push(Order)
	pop() Order
}

type OrderBook struct {
	Name string

	Token     Token
	BaseToken Token

	Buys  priorityQueue
	Sells priorityQueue

	LastTradedPrice int
}
