package entities

// Transaction fundamentally resolves explicit Edge/Node relationships on the ledger diagram.
type Transaction struct {
	ID       string
	Sender   *User
	Receiver *User
	Amount   float64
}
