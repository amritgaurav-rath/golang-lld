package strategy
import (
	"app/food-ordering-system/go-without-multithreading/entities"
	"math"
)

type LowestCostStrategy struct{}

func (s *LowestCostStrategy) SelectRestaurant(eligibleRestaurants []*entities.Restaurant, requestedItems map[string]int) *entities.Restaurant {
	var bestRestaurant *entities.Restaurant
	minCost := math.MaxFloat64

	for _, r := range eligibleRestaurants {
		currentCost := 0.0
		for item, qty := range requestedItems {
			currentCost += (r.Menu[item] * float64(qty))
		}
		if currentCost < minCost {
			minCost = currentCost
			bestRestaurant = r
		}
	}
	return bestRestaurant
}
