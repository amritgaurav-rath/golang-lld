package entities;
import java.util.List;
import enums.BookingStatus;

public class Booking {
    private String id;
    private User user;
    private Show show;
    private List<Seat> seats;
    private double totalPrice;
    private BookingStatus status;

    public Booking(String id, User user, Show show, List<Seat> seats, double totalPrice, BookingStatus status) {
        this.id = id;
        this.user = user;
        this.show = show;
        this.seats = seats;
        this.totalPrice = totalPrice;
        this.status = status;
    }

    public String getId() { return id; }
    public BookingStatus getStatus() { return status; }
    public void setStatus(BookingStatus status) { this.status = status; }
    public double getTotalPrice() { return totalPrice; }
    public List<Seat> getSeats() { return seats; }
}
