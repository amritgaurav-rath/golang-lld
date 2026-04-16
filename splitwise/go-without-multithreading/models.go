package main

// User represents a user in the Splitwise system
type User struct {
	ID    string
	Name  string
	Email string
}

// Group represents a group of friends interacting
type Group struct {
	ID       string
	Name     string
	Members  []*User
	Expenses []*Expense
}

// Expense represents an expense within a group
type Expense struct {
	ID          string
	Amount      float64
	Description string
	PaidBy      *User
	Splits      []Split
}

// Transaction represents a settlement transaction between two users
type Transaction struct {
	ID       string
	Sender   *User
	Receiver *User
	Amount   float64
}
