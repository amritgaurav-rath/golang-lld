package main

import (
	"app/food-ordering-system/go-without-multithreading/enums"
	"app/food-ordering-system/go-without-multithreading/services"
	"app/food-ordering-system/go-without-multithreading/strategy"
	"fmt"
)

func main() {
	fmt.Println("🚀 Initializing S.O.L.I.D. FAANG Feed.Me Food Ordering System (Golang)...")
	sys := services.GetInstance()

	// 1. Onboard
	sys.OnboardRestaurant("R1", 5, 4.5, map[string]float64{"Veg Biryani": 100.0, "Chicken Biryani": 150.0})
	sys.OnboardRestaurant("R2", 5, 4.0, map[string]float64{"Chicken Biryani": 175.0, "Idli": 10.0, "Dosa": 50.0, "Veg Biryani": 80.0})
	sys.OnboardRestaurant("R3", 1, 4.9, map[string]float64{"Gobi Manchurian": 150.0, "Idli": 15.0, "Chicken Biryani": 175.0, "Dosa": 30.0})

	// 2. Update Menu
	sys.UpdateMenu("R1", "add", "Chicken65", 250.0)
	sys.UpdateMenu("R2", "update", "Chicken Biryani", 150.0)

	// 3. Sequential Order Test Cases
	fmt.Println("\nOrder1: Ashwin [3Idli, 1Dosa] (Lowest Cost)")
	sys.PlaceOrder("Order1", "Ashwin", map[string]int{"Idli": 3, "Dosa": 1}, &strategy.LowestCostStrategy{}) // Expect R3

	fmt.Println("\nOrder2: Harish [3Idli, 1Dosa] (Lowest Cost)")
	sys.PlaceOrder("Order2", "Harish", map[string]int{"Idli": 3, "Dosa": 1}, &strategy.LowestCostStrategy{}) // Expect R2

	fmt.Println("\nOrder3: Shruthi [3Veg Biryani] (Highest Rating)")
	sys.PlaceOrder("Order3", "Shruthi", map[string]int{"Veg Biryani": 3}, &strategy.HighestRatingStrategy{}) // Expect R1

	fmt.Println("\n--- Updating Status: R3 marks Order1 COMPLETED ---")
	sys.UpdateOrderStatus("Order1", enums.Completed)

	fmt.Println("\nOrder4: Harish [3Idli, 1Dosa] (Lowest Cost)")
	sys.PlaceOrder("Order4", "Harish", map[string]int{"Idli": 3, "Dosa": 1}, &strategy.LowestCostStrategy{}) // Expect R3

	fmt.Println("\nOrder5: xyz [1Paneer Tikka, 1Idli] (Lowest Cost)")
	_, err := sys.PlaceOrder("Order5", "xyz", map[string]int{"Paneer Tikka": 1, "Idli": 1}, &strategy.LowestCostStrategy{})
	if err != nil {
		fmt.Println("Validation Error Captured:", err.Error())
	}

	fmt.Println("\nBONUS ORDER: BonusUser [1Veg Biryani] (Max Capacity Strategy)")
	sys.PlaceOrder("Order6", "BonusUser", map[string]int{"Veg Biryani": 1}, &strategy.MaxCapacityStrategy{})
}
