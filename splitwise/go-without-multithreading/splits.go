package main

// Split represents the abstract split of an expense
type Split interface {
	GetUser() *User
	GetAmount() float64
	SetAmount(amount float64)
}

// baseSplit provides common fields for all Split types
type baseSplit struct {
	User   *User
	Amount float64
}

func (b *baseSplit) GetUser() *User {
	return b.User
}

func (b *baseSplit) GetAmount() float64 {
	return b.Amount
}

func (b *baseSplit) SetAmount(amount float64) {
	b.Amount = amount
}

// EqualSplit distributes the expense equally
type EqualSplit struct {
	baseSplit
}

func NewEqualSplit(user *User) *EqualSplit {
	return &EqualSplit{
		baseSplit: baseSplit{User: user},
	}
}

// ExactSplit distributes the expense by exact specified amounts
type ExactSplit struct {
	baseSplit
}

func NewExactSplit(user *User, amount float64) *ExactSplit {
	return &ExactSplit{
		baseSplit: baseSplit{User: user, Amount: amount},
	}
}

// PercentSplit distributes the expense by a percentage
type PercentSplit struct {
	baseSplit
	Percent float64
}

func NewPercentSplit(user *User, percent float64) *PercentSplit {
	return &PercentSplit{
		baseSplit: baseSplit{User: user},
		Percent:   percent,
	}
}
