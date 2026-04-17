package entities;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class Restaurant {
    private String name;
    private double rating;
    private int maxOrders;
    private int currentOrders;
    private Map<String, Double> menu;

    public Restaurant(String name, int maxOrders, double rating) {
        this.name = name;
        this.maxOrders = maxOrders;
        this.rating = rating;
        this.menu = new ConcurrentHashMap<>();
        this.currentOrders = 0;
    }

    public String getName() { return name; }
    public double getRating() { return rating; }
    public int getMaxOrders() { return maxOrders; }
    public int getCurrentOrders() { return currentOrders; }
    public Map<String, Double> getMenu() { return menu; }

    public void updateMenu(String itemName, Double price) {
        this.menu.put(itemName, price);
    }

    public void updateCapacity(int capacity) {
        this.maxOrders = capacity;
    }

    public boolean canFulfill(Map<String, Integer> requestedItems) {
        // Natively blocks mathematical routing if current operational grid is fully maxed exactly
        if (currentOrders >= maxOrders) return false;
        
        // Loop structural dependencies natively verifying absolute fulfillment dynamically
        for (String item : requestedItems.keySet()) {
            if (!menu.containsKey(item)) return false;
        }
        return true;
    }

    public void incrementCurrentOrders() { this.currentOrders++; }
    
    public void markOrderCompleted() {
        if(currentOrders > 0) this.currentOrders--;
    }
}
