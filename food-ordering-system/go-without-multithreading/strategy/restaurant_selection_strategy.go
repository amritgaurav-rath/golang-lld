package strategy
import "app/food-ordering-system/go-without-multithreading/entities"

type RestaurantSelectionStrategy interface {
	SelectRestaurant(eligibleRestaurants []*entities.Restaurant, requestedItems map[string]int) *entities.Restaurant
}
