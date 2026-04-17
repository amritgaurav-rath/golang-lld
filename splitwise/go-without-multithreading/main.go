package main

import (
	"app/splitwise/go-without-multithreading/entities"
	"app/splitwise/go-without-multithreading/strategy"
	"fmt"
)

func main() {
	fmt.Println("🚀 Starting Splitwise (Senior OCP Golang Subfolders)")

	service := GetSplitwiseService()

	u1 := entities.NewUser("1", "John", "john@example.com")
	u2 := entities.NewUser("2", "Jane", "jane@example.com")
	u3 := entities.NewUser("3", "Bob", "bob@example.com")

	service.AddUser(u1)
	service.AddUser(u2)
	service.AddUser(u3)

	group := entities.NewGroup("1", "Trip to Paris")
	group.AddMember(u1)
	group.AddMember(u2)
	group.AddMember(u3)
	service.AddGroup(group)

	participants := []*entities.User{u1, u2, u3}

	// 1. Equitably slice math utilizing pure Strategy package engines
	var equalStrategy strategy.SplitStrategy = &strategy.EqualSplitStrategy{}
	splits1, _ := equalStrategy.CalculateSplits(300, participants, nil)
	
	// 2. Factory creation decoupling Cyclic Dependencies naturally
	e1 := entities.NewExpense("1", 300, "Dinner", u1, splits1)
	service.AddExpense(group.ID, e1)

	var percentStrategy strategy.SplitStrategy = &strategy.PercentSplitStrategy{}
	splits2, err := percentStrategy.CalculateSplits(100, participants, []float64{20, 50, 30})
	if err == nil {
		e2 := entities.NewExpense("2", 100, "Breakfast", u2, splits2)
		service.AddExpense(group.ID, e2)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	service.SettleBalance("1", "2") // Settle John and Jane
	service.SettleBalance("1", "3") // Settle John and Bob
}
