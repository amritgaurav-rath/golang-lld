package main

import (
	"app/splitwise/go-without-multithreading/entities"
	"fmt"
	"sync"
)

// SplitwiseService operates as a Singleton Facade.
type SplitwiseService struct {
	users  map[string]*entities.User
	groups map[string]*entities.Group
	mu     sync.Mutex
}

var instance *SplitwiseService
var once sync.Once

func GetSplitwiseService() *SplitwiseService {
	once.Do(func() {
		instance = &SplitwiseService{
			users:  make(map[string]*entities.User),
			groups: make(map[string]*entities.Group),
		}
	})
	return instance
}

func (s *SplitwiseService) AddUser(user *entities.User) {
	s.users[user.ID] = user
}

func (s *SplitwiseService) AddGroup(group *entities.Group) {
	s.groups[group.ID] = group
}

func (s *SplitwiseService) AddExpense(groupID string, expense *entities.Expense) {
	group, exists := s.groups[groupID]
	if !exists {
		return
	}
	group.AddExpense(expense)

	s.updateBalances(expense)
	fmt.Printf("Expense '%s' of amount $%.2f safely stored.\n", expense.Description, expense.Amount)
}

func (s *SplitwiseService) updateBalances(expense *entities.Expense) {
	for _, split := range expense.Splits {
		paidBy := expense.PaidBy
		user := split.User
		amount := split.Amount
		if paidBy != user {
			s.updateBalance(paidBy, user, amount)
			s.updateBalance(user, paidBy, -amount)
		}
	}
}

func (s *SplitwiseService) updateBalance(user1, user2 *entities.User, amount float64) {
	key := user1.ID + ":" + user2.ID
	user1.Balances[key] += amount
}

func (s *SplitwiseService) SettleBalance(userID1, userID2 string) {
	user1, exists1 := s.users[userID1]
	user2, exists2 := s.users[userID2]
	if !exists1 || !exists2 {
		return
	}

	key := user1.ID + ":" + user2.ID
	balance := user1.Balances[key]

	if balance > 0 {
		s.createTransaction(user2, user1, balance)
		user1.Balances[key] = 0
		user2.Balances[user2.ID+":"+user1.ID] = 0
	} else if balance < 0 {
		s.createTransaction(user1, user2, -balance)
		user1.Balances[key] = 0
		user2.Balances[user2.ID+":"+user1.ID] = 0
	}
}

func (s *SplitwiseService) createTransaction(sender, receiver *entities.User, amount float64) {
	fmt.Printf("Transaction: %s securely pays %s an amount of $%.2f\n", sender.Name, receiver.Name, amount)
}
