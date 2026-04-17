package entities

// User represents a core entity in the ecosystem, natively storing its own bidirectional debt matrix.
type User struct {
	ID       string
	Name     string
	Email    string
	Balances map[string]float64
}

func NewUser(id, name, email string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Balances: make(map[string]float64),
	}
}
