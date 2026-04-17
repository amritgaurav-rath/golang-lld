package strategy;
import entities.Restaurant;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public class MaxCapacityStrategy implements RestaurantSelectionStrategy {
    @Override
    public Optional<Restaurant> selectRestaurant(List<Restaurant> eligibleRestaurants, Map<String, Integer> requestedItems) {
        Restaurant bestRestaurant = null;
        int maxRemaining = -1;

        for (Restaurant r : eligibleRestaurants) {
            int remaining = r.getMaxOrders() - r.getCurrentOrders();
            if (remaining > maxRemaining) {
                maxRemaining = remaining;
                bestRestaurant = r;
            }
        }
        return Optional.ofNullable(bestRestaurant);
    }
}
