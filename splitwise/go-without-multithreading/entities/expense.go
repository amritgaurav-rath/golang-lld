package entities

// Expense records the orchestration of an upfront financial layout.
type Expense struct {
	ID          string
	Amount      float64
	Description string
	PaidBy      *User
	Splits      []Split
}

// NewExpense stores pre-calculated splits natively seamlessly bypassing any package cyclic dependencies.
func NewExpense(id string, amount float64, description string, paidBy *User, splits []Split) *Expense {
	return &Expense{
		ID:          id,
		Amount:      amount,
		Description: description,
		PaidBy:      paidBy,
		Splits:      splits,
	}
}
