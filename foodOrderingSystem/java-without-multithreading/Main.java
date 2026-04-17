import java.util.HashMap;
import java.util.Map;
import services.FoodOrderingSystem;
import strategy.LowestCostStrategy;
import strategy.HighestRatingStrategy;
import strategy.MaxCapacityStrategy;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Initializing S.O.L.I.D. FAANG Feed.Me Food Ordering System...");
        FoodOrderingSystem sys = FoodOrderingSystem.getInstance();

        // 1. Onboard Nodes natively
        Map<String, Double> m1 = new HashMap<>(); m1.put("Veg Biryani", 100.0); m1.put("Chicken Biryani", 150.0);
        sys.onboardRestaurant("R1", 5, 4.5, m1);

        Map<String, Double> m2 = new HashMap<>(); m2.put("Chicken Biryani", 175.0); m2.put("Idli", 10.0); m2.put("Dosa", 50.0); m2.put("Veg Biryani", 80.0);
        sys.onboardRestaurant("R2", 5, 4.0, m2);

        Map<String, Double> m3 = new HashMap<>(); m3.put("Gobi Manchurian", 150.0); m3.put("Idli", 15.0); m3.put("Chicken Biryani", 175.0); m3.put("Dosa", 30.0);
        sys.onboardRestaurant("R3", 1, 4.9, m3);

        try {
            // 2. Update Menu Parameters
            sys.updateMenu("R1", "add", "Chicken65", 250.0);
            sys.updateMenu("R2", "update", "Chicken Biryani", 150.0);

            // 3. Sequential Order Simulation
            System.out.println("\nOrder1: Ashwin [3Idli, 1Dosa] (Lowest Cost)");
            Map<String, Integer> items1 = new HashMap<>(); items1.put("Idli", 3); items1.put("Dosa", 1);
            sys.placeOrder("Order1", "Ashwin", items1, new LowestCostStrategy()); // Expect R3 (15*3+30=75)

            System.out.println("\nOrder2: Harish [3Idli, 1Dosa] (Lowest Cost)");
            Map<String, Integer> items2 = new HashMap<>(); items2.put("Idli", 3); items2.put("Dosa", 1);
            sys.placeOrder("Order2", "Harish", items2, new LowestCostStrategy()); // Expect R2 (R3 is maxed, R2=10*3+50=80)

            System.out.println("\nOrder3: Shruthi [3Veg Biryani] (Highest Rating)");
            Map<String, Integer> items3 = new HashMap<>(); items3.put("Veg Biryani", 3);
            sys.placeOrder("Order3", "Shruthi", items3, new HighestRatingStrategy()); // Expect R1 (Rating 4.5)

            System.out.println("\n--- Updating Status: R3 marks Order1 COMPLETED ---");
            sys.updateOrderStatus("Order1", enums.OrderStatus.COMPLETED); // Drops R3 back to 0 capacity
            
            System.out.println("\nOrder4: Harish [3Idli, 1Dosa] (Lowest Cost)");
            Map<String, Integer> items4 = new HashMap<>(); items4.put("Idli", 3); items4.put("Dosa", 1);
            sys.placeOrder("Order4", "Harish", items4, new LowestCostStrategy()); // Expect R3 (back down to 75 cost)

            System.out.println("\nOrder5: xyz [1Paneer Tikka, 1Idli] (Lowest Cost)");
            try {
                Map<String, Integer> items5 = new HashMap<>(); items5.put("Paneer Tikka", 1); items5.put("Idli", 1);
                sys.placeOrder("Order5", "xyz", items5, new LowestCostStrategy());
            } catch (Exception e) {
                System.out.println("Validation Error Captured: " + e.getMessage()); // Should fire Exception successfully!
            }

            System.out.println("\nBONUS ORDER: BonusUser [1Veg Biryani] (Max Capacity Strategy)");
            Map<String, Integer> items6 = new HashMap<>(); items6.put("Veg Biryani", 1);
            sys.placeOrder("Order6", "BonusUser", items6, new MaxCapacityStrategy()); // Successfully tests the bonus requirements natively.

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
