package strategy;
import entities.Restaurant;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public class LowestCostStrategy implements RestaurantSelectionStrategy {
    @Override
    public Optional<Restaurant> selectRestaurant(List<Restaurant> eligibleRestaurants, Map<String, Integer> requestedItems) {
        Restaurant bestRestaurant = null;
        double minCost = Double.MAX_VALUE;

        // Decoupled dynamic loop evaluating variable item amounts correctly mapped natively!
        for (Restaurant r : eligibleRestaurants) {
            double currentCost = 0;
            for (Map.Entry<String, Integer> entry : requestedItems.entrySet()) {
                currentCost += (r.getMenu().get(entry.getKey()) * entry.getValue());
            }

            if (currentCost < minCost) {
                minCost = currentCost;
                bestRestaurant = r;
            }
        }
        return Optional.ofNullable(bestRestaurant);
    }
}
