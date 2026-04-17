package strategy
import "app/foodOrderingSystem/go-without-multithreading/entities"

type HighestRatingStrategy struct{}

func (s *HighestRatingStrategy) SelectRestaurant(eligibleRestaurants []*entities.Restaurant, requestedItems map[string]int) *entities.Restaurant {
	var bestRestaurant *entities.Restaurant
	maxRating := -1.0

	for _, r := range eligibleRestaurants {
		if r.Rating > maxRating {
			maxRating = r.Rating
			bestRestaurant = r
		}
	}
	return bestRestaurant
}
