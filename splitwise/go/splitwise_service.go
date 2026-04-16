package main

import (
	"fmt"
	"math"
	"sync"
)

// SplitwiseService is a singleton managing the system concurrently
type SplitwiseService struct {
	mu       sync.RWMutex // Protects concurrent access to maps
	Users    map[string]*User
	Groups   map[string]*Group
	// Balances maps [UserA_ID][UserB_ID] = Amount that User A owes to User B.
	Balances map[string]map[string]float64
}

var (
	instance *SplitwiseService
	once     sync.Once
)

// GetSplitwiseService securely initializes and returns the singleton instance
func GetSplitwiseService() *SplitwiseService {
	once.Do(func() {
		instance = &SplitwiseService{
			Users:    make(map[string]*User),
			Groups:   make(map[string]*Group),
			Balances: make(map[string]map[string]float64),
		}
	})
	return instance
}

// AddUser securely registers a new user in the system
func (s *SplitwiseService) AddUser(u *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Users[u.ID] = u
	s.Balances[u.ID] = make(map[string]float64)
}

// AddGroup securely registers a new group
func (s *SplitwiseService) AddGroup(g *Group) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Groups[g.ID] = g
}

// AddExpense concurrently processes and splits a group expense
func (s *SplitwiseService) AddExpense(groupID string, expense *Expense) error {
	s.mu.Lock()
	group, exists := s.Groups[groupID]
	if !exists {
		s.mu.Unlock()
		return fmt.Errorf("group %s does not exist", groupID)
	}
	s.mu.Unlock()

	// Route the specific calculation based on split types natively before locking
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
		total := 0.0
		for _, split := range expense.Splits {
			total += split.GetAmount()
		}
		if total != expense.Amount {
			return fmt.Errorf("exact splits total does not match expense amount")
		}
	}

	// Update Balances globally locking the system graph
	s.mu.Lock()
	defer s.mu.Unlock()

	group.Expenses = append(group.Expenses, expense)
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

// SettleBalance securely zeros out mutual debits
func (s *SplitwiseService) SettleBalance(userA, userB string) (*Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	amountOwed := s.Balances[userA][userB]
	if amountOwed == 0 {
		return nil, fmt.Errorf("no balance to settle between %s and %s", userA, userB)
	}

	var sender, receiver string
	var amount float64

	if amountOwed > 0 {
		sender, receiver, amount = userA, userB, amountOwed
	} else {
		sender, receiver, amount = userB, userA, -amountOwed
	}

	// Reset balances iteratively
	s.Balances[userA][userB] = 0
	s.Balances[userB][userA] = 0

	return &Transaction{
		ID:       fmt.Sprintf("TRX-%s-%s", sender, receiver),
		Sender:   s.Users[sender],
		Receiver: s.Users[receiver],
		Amount:   amount,
	}, nil
}

// PrintBalances securely traces all nodes
func (s *SplitwiseService) PrintBalances() {
	s.mu.RLock()
	defer s.mu.RUnlock()

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
