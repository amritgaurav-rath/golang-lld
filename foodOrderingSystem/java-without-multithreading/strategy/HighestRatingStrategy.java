package strategy;
import entities.Restaurant;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public class HighestRatingStrategy implements RestaurantSelectionStrategy {
    @Override
    public Optional<Restaurant> selectRestaurant(List<Restaurant> eligibleRestaurants, Map<String, Integer> requestedItems) {
        Restaurant bestRestaurant = null;
        double maxRating = -1.0;

        for (Restaurant r : eligibleRestaurants) {
            if (r.getRating() > maxRating) {
                maxRating = r.getRating();
                bestRestaurant = r;
            }
        }
        return Optional.ofNullable(bestRestaurant);
    }
}
