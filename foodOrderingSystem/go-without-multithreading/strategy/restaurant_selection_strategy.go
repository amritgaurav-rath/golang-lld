package strategy
import "app/foodOrderingSystem/go-without-multithreading/entities"

type RestaurantSelectionStrategy interface {
	SelectRestaurant(eligibleRestaurants []*entities.Restaurant, requestedItems map[string]int) *entities.Restaurant
}
