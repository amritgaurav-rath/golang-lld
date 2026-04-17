package entities;
import enums.OrderStatus;
import java.util.Map;

public class Order {
    private String orderId;
    private String customerName;
    private String restaurantName;
    private Map<String, Integer> items;
    private OrderStatus status;

    public Order(String orderId, String customerName, String restaurantName, Map<String, Integer> items) {
        this.orderId = orderId;
        this.customerName = customerName;
        this.restaurantName = restaurantName;
        this.items = items;
        this.status = OrderStatus.ACCEPTED;
    }

    public String getOrderId() { return orderId; }
    public String getCustomerName() { return customerName; }
    public String getRestaurantName() { return restaurantName; }
    public Map<String, Integer> getItems() { return items; }
    public OrderStatus getStatus() { return status; }

    public void setStatus(OrderStatus status) { this.status = status; }
}
