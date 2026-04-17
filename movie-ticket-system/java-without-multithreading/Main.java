import entities.*;
import enums.*;
import observer.UserObserver;
import strategy.WeekdayPricingStrategy;
import strategy.WeekendPricingStrategy;
import services.MovieTicketBookingSystem;
import java.util.*;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Initializing Single-Threaded Movie Ticket System with Explicit Design Patterns");
        MovieTicketBookingSystem sys = MovieTicketBookingSystem.getInstance();

        User alice = new User("U1", "Alice", "alice@test.com");
        User bob = new User("U2", "Bob", "bob@test.com");
        User charlie = new User("U3", "Charlie", "charlie@test.com");

        Movie movie = new Movie("M1", "Inception", "Sci-Fi", 148);
        
        // 1. Observer Explicit Binding
        movie.addObserver(new UserObserver(alice));
        movie.addObserver(new UserObserver(bob));
        System.out.println("\n--- Triggering Observer Push Notifications ---");
        sys.addMovie(movie); // Notification triggers natively when catalog expands

        Theater theater = new Theater("T1", "PVR Cinemas", "Mumbai");
        sys.addTheater(theater);

        Show show = new Show("SH1", movie, theater, System.currentTimeMillis(), System.currentTimeMillis() + 7200000);
        
        for (int i = 1; i <= 5; i++) {
            show.addSeat(new Seat("A" + i, 1, i, SeatType.PREMIUM, 25.0, SeatStatus.AVAILABLE));
            show.addSeat(new Seat("B" + i, 2, i, SeatType.NORMAL, 10.0, SeatStatus.AVAILABLE));
        }
        sys.addShow(show);

        System.out.println("\n--- Explicit Strategy Pattern Implementation ---");
        System.out.println("Alice attempts to dynamically book A1, A2 using WEEKEND Pricing Strategy (+$10 / seat surge).");
        
        try {
            // Evaluates mathematical pricing natively on the fly via Strategy Interfaces!
            Booking booking1 = sys.bookTickets(alice, show, Arrays.asList("A1", "A2"), new WeekendPricingStrategy());
            System.out.printf("✅ [Alice] Booking Reserved (PENDING)! Booking ID: %s, Real-Time Amount: $%.2f\n", booking1.getId(), booking1.getTotalPrice());
            sys.confirmBooking(booking1.getId());
        } catch (Exception e) {
            System.out.println("❌ Booking failed: " + e.getMessage());
        }

        System.out.println("\nBob attempts to book overlapping Seats A1, A2 via WEEKDAY Pricing Strategy (Standard)...");
        try {
            sys.bookTickets(bob, show, Arrays.asList("A1", "A2"), new WeekdayPricingStrategy());
        } catch (Exception e) {
            System.out.println("❌ [Bob] Booking Failed accurately (already natively taken): " + e.getMessage());
        }

        System.out.println("\nCharlie bookings B1, B2 safely on standard Weekday Pricing...");
        try {
            Booking booking3 = sys.bookTickets(charlie, show, Arrays.asList("B1", "B2"), new WeekdayPricingStrategy());
            System.out.printf("✅ [Charlie] Reserved B1, B2 dynamically. Total: $%.2f. Status: PENDING.\n", booking3.getTotalPrice());
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }
}
