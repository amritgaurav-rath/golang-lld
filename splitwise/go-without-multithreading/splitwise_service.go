package main

import (
	"fmt"
	"math"
)

// SplitwiseService is a singleton managing the system natively without multithreading
type SplitwiseService struct {
	Users    map[string]*User
	Groups   map[string]*Group
	// Balances maps [UserA_ID][UserB_ID] = Amount that User A owes to User B.
	// Positive amount means A owes B. Negative means B owes A.
	Balances map[string]map[string]float64 
}

var instance *SplitwiseService

func GetSplitwiseService() *SplitwiseService {
	if instance == nil {
		instance = &SplitwiseService{
			Users:    make(map[string]*User),
			Groups:   make(map[string]*Group),
			Balances: make(map[string]map[string]float64),
		}
	}
	return instance
}

func (s *SplitwiseService) AddUser(u *User) {
	s.Users[u.ID] = u
	s.Balances[u.ID] = make(map[string]float64)
}

func (s *SplitwiseService) AddGroup(g *Group) {
	s.Groups[g.ID] = g
}

func (s *SplitwiseService) AddExpense(groupID string, expense *Expense) error {
	group, exists := s.Groups[groupID]
	if !exists {
		return fmt.Errorf("group %s does not exist", groupID)
	}

	// Route the specific calculation based on split types
	switch expense.Splits[0].(type) {
	case *EqualSplit:
		// Divide the total amount equally amongst all participants
		amountPerUser := expense.Amount / float64(len(expense.Splits))
		amountPerUser = math.Round(amountPerUser*100) / 100
		for _, split := range expense.Splits {
			split.SetAmount(amountPerUser)
		}
	case *PercentSplit:
		// Calculate precise monetary amounts based on percentage shares
		for _, rawSplit := range expense.Splits {
			percentSplit := rawSplit.(*PercentSplit)
			amount := (expense.Amount * percentSplit.Percent) / 100.0
			percentSplit.SetAmount(math.Round(amount*100) / 100)
		}
	case *ExactSplit:
		// Exact split already has amounts predefined
		total := 0.0
		for _, split := range expense.Splits {
			total += split.GetAmount()
		}
		if total != expense.Amount {
			return fmt.Errorf("exact splits total does not match expense amount")
		}
	}

	// Add expense to group
	group.Expenses = append(group.Expenses, expense)

	// Update Balances
	paidBy := expense.PaidBy.ID
	for _, split := range expense.Splits {
		splitUser := split.GetUser().ID
		if paidBy == splitUser {
			continue // Prevent users from owing themselves
		}
		// Adjust bi-directional debt edges
		s.Balances[splitUser][paidBy] += split.GetAmount()
		s.Balances[paidBy][splitUser] -= split.GetAmount()
	}

	return nil
}

func (s *SplitwiseService) SettleBalance(userA, userB string) (*Transaction, error) {
	amountOwed := s.Balances[userA][userB]
	if amountOwed == 0 {
		return nil, fmt.Errorf("no balance to settle between %s and %s", userA, userB)
	}

	var sender, receiver string
	var amount float64

	if amountOwed > 0 {
		// A owes B
		sender = userA
		receiver = userB
		amount = amountOwed
	} else {
		// B owes A
		sender = userB
		receiver = userA
		amount = -amountOwed
	}

	// Reset balances
	s.Balances[userA][userB] = 0
	s.Balances[userB][userA] = 0

	return &Transaction{
		ID:       fmt.Sprintf("TRX-%s-%s", sender, receiver),
		Sender:   s.Users[sender],
		Receiver: s.Users[receiver],
		Amount:   amount,
	}, nil
}

func (s *SplitwiseService) PrintBalances() {
	fmt.Println("--- System Balances ---")
	for userA, owesMap := range s.Balances {
		for userB, amount := range owesMap {
			if amount > 0 {
				fmt.Printf("%s owes %s: $%.2f\n", s.Users[userA].Name, s.Users[userB].Name, amount)
			}
		}
	}
	fmt.Println("-----------------------")
}
