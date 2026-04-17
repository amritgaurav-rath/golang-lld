package strategy
import "app/food-ordering-system/go-without-multithreading/entities"

type MaxCapacityStrategy struct{}

func (s *MaxCapacityStrategy) SelectRestaurant(eligibleRestaurants []*entities.Restaurant, requestedItems map[string]int) *entities.Restaurant {
	var bestRestaurant *entities.Restaurant
	maxRemaining := -1

	for _, r := range eligibleRestaurants {
		remaining := r.MaxOrders - r.CurrentOrders
		if remaining > maxRemaining {
			maxRemaining = remaining
			bestRestaurant = r
		}
	}
	return bestRestaurant
}
