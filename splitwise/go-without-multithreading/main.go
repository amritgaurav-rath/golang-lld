package main

import "fmt"

func main() {
	fmt.Println("🚀 Starting Splitwise (Synchronous)")

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

	// Equal Split Expense
	expense1 := &Expense{
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
	err := service.AddExpense(group.ID, expense1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Percent Split Expense
	expense2 := &Expense{
		ID:          "EXP2",
		Amount:      100.0,
		Description: "Dinner",
		PaidBy:      bob,
		Splits: []Split{
			NewPercentSplit(alice, 20), // 20
			NewPercentSplit(bob, 50),   // 50
			NewPercentSplit(charlie, 30),// 30
		},
	}
	err = service.AddExpense(group.ID, expense2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	service.PrintBalances()

	// Settle Bob and Charlie
	fmt.Println("Attempting to settle Charlie and Bob...")
	transaction, err := service.SettleBalance(charlie.ID, bob.ID)
	if err != nil {
		fmt.Println("Error settling:", err)
	} else {
		fmt.Printf("✅ Settled! %s paid %s $%.2f\n", transaction.Sender.Name, transaction.Receiver.Name, transaction.Amount)
	}

	service.PrintBalances()
}
