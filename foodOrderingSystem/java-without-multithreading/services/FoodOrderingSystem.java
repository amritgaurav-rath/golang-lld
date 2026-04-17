package services;
import entities.Order;
import entities.Restaurant;
import strategy.RestaurantSelectionStrategy;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public class FoodOrderingSystem {
    private Map<String, Restaurant> restaurants;
    private Map<String, Order> orders;
    private static FoodOrderingSystem instance;

    private FoodOrderingSystem() {
        this.restaurants = new ConcurrentHashMap<>();
        this.orders = new ConcurrentHashMap<>();
    }

    public static synchronized FoodOrderingSystem getInstance() {
        if (instance == null) {
            instance = new FoodOrderingSystem();
        }
        return instance;
    }

    public void onboardRestaurant(String name, int maxOrders, double rating, Map<String, Double> menu) {
        Restaurant r = new Restaurant(name, maxOrders, rating);
        for(Map.Entry<String, Double> entry : menu.entrySet()) {
            r.updateMenu(entry.getKey(), entry.getValue());
        }
        restaurants.put(name, r);
    }

    public void updateMenu(String name, String action, String itemName, Double price) throws Exception {
        if (!restaurants.containsKey(name)) throw new Exception("Restaurant unmapped natively.");
        if (action.equalsIgnoreCase("add") || action.equalsIgnoreCase("update")) {
            restaurants.get(name).updateMenu(itemName, price);
        }
    }

    public void updateCapacity(String name, int capacity) throws Exception {
        if (!restaurants.containsKey(name)) throw new Exception("Restaurant computationally unmapped.");
        restaurants.get(name).updateCapacity(capacity);
    }

    public Order placeOrder(String orderId, String user, Map<String, Integer> items, RestaurantSelectionStrategy strategy) throws Exception {
        List<Restaurant> eligibleRestaurants = new ArrayList<>();
        
        // Validation boundary: filter exclusively restaurants physically serving all requested mapping items AND having safe capacity runtime limits!
        for (Restaurant r : restaurants.values()) {
            if (r.canFulfill(items)) {
                eligibleRestaurants.add(r);
            }
        }

        if (eligibleRestaurants.isEmpty()) {
            throw new Exception("Output: Order can't be fulfilled");
        }

        Optional<Restaurant> selectedOpt = strategy.selectRestaurant(eligibleRestaurants, items);
        if (!selectedOpt.isPresent()) {
            throw new Exception("Strategy isolated mathematically matching zero natively.");
        }

        Restaurant selected = selectedOpt.get();
        selected.incrementCurrentOrders();
        
        Order order = new Order(orderId, user, selected.getName(), items);
        orders.put(orderId, order);
        
        System.out.printf("Output: Order assigned to %s\n", selected.getName());
        return order;
    }

    public void updateOrderStatus(String orderId, enums.OrderStatus status) throws Exception {
        if (!orders.containsKey(orderId)) throw new Exception("Order strictly unmapped natively.");
        Order o = orders.get(orderId);
        
        if (o.getStatus() == enums.OrderStatus.COMPLETED) {
            throw new Exception("Order structurally already explicitly processed.");
        }

        if (status == enums.OrderStatus.COMPLETED) {
            o.setStatus(status);
            restaurants.get(o.getRestaurantName()).markOrderCompleted();
        }
    }
}
