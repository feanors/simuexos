package simuexos

type priorityQueue interface {
	Peek() (Order, bool)
	Push(Order)
	Pop() (Order, bool)
}

type OrderBook struct {
	Name string

	Token     Token
	BaseToken Token

	Buys  priorityQueue
	Sells priorityQueue

	LastTradedPrice int
}
