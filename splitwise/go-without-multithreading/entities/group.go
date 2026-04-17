package entities

// Group structurally consolidates a collection of Users tracking sequential Expenses together.
type Group struct {
	ID       string
	Name     string
	Members  []*User
	Expenses []*Expense
}

func NewGroup(id, name string) *Group {
	return &Group{
		ID:       id,
		Name:     name,
		Members:  []*User{},
		Expenses: []*Expense{},
	}
}

func (g *Group) AddMember(u *User) {
	g.Members = append(g.Members, u)
}

func (g *Group) AddExpense(e *Expense) {
	g.Expenses = append(g.Expenses, e)
}
