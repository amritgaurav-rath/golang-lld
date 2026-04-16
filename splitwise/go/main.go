package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("🚀 Starting Splitwise (Multi-Threaded / Concurrent)")

	service := GetSplitwiseService()

	alice := &User{ID: "U1", Name: "Alice", Email: "alice@example.com"}
	bob := &User{ID: "U2", Name: "Bob", Email: "bob@example.com"}
	charlie := &User{ID: "U3", Name: "Charlie", Email: "charlie@example.com"}

	service.AddUser(alice)
	service.AddUser(bob)
	service.AddUser(charlie)

	group := &Group{
		ID:      "G1",
		Name:    "Vacation",
		Members: []*User{alice, bob, charlie},
	}
	service.AddGroup(group)

	var wg sync.WaitGroup

	// Dispatch multiple parallel expenses across routines
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp := &Expense{
			ID:          "EXP1",
			Amount:      300.0,
			Description: "Hotel",
			PaidBy:      alice,
			Splits: []Split{
				NewEqualSplit(alice),
				NewEqualSplit(bob),
				NewEqualSplit(charlie),
			},
		}
		service.AddExpense(group.ID, exp)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		exp := &Expense{
			ID:          "EXP2",
			Amount:      100.0,
			Description: "Dinner",
			PaidBy:      bob,
			Splits: []Split{
				NewPercentSplit(alice, 20),
				NewPercentSplit(bob, 50),
				NewPercentSplit(charlie, 30),
			},
		}
		service.AddExpense(group.ID, exp)
	}()

	// Wait for concurrent expenses to conclude
	wg.Wait()

	service.PrintBalances()

	// Perform settlement securely
	fmt.Println("Attempting to settle Charlie and Bob natively...")
	transaction, err := service.SettleBalance(charlie.ID, bob.ID)
	if err != nil {
		fmt.Println("Error settling:", err)
	} else {
		fmt.Printf("✅ Settled! %s paid %s $%.2f\n", transaction.Sender.Name, transaction.Receiver.Name, transaction.Amount)
	}

	service.PrintBalances()
}
