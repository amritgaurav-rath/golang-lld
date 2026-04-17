package strategy;
import entities.Restaurant;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public interface RestaurantSelectionStrategy {
    Optional<Restaurant> selectRestaurant(List<Restaurant> eligibleRestaurants, Map<String, Integer> requestedItems);
}
