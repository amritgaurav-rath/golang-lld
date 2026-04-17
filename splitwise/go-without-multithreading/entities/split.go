package entities

// Split is a concrete, immutable data carrier resolving precisely how much a User owes in a given Expense.
type Split struct {
	User   *User
	Amount float64
}
