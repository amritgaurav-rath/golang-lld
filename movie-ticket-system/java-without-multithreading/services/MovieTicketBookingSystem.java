package services;
import entities.*;
import enums.*;
import strategy.PricingStrategy;
import java.util.*;

public class MovieTicketBookingSystem {
    private Map<String, Movie> movies;
    private Map<String, Theater> theaters;
    private Map<String, Show> shows;
    private Map<String, Booking> bookings;
    
    private static MovieTicketBookingSystem instance;

    private MovieTicketBookingSystem() {
        movies = new HashMap<>();
        theaters = new HashMap<>();
        shows = new HashMap<>();
        bookings = new HashMap<>();
    }

    public static synchronized MovieTicketBookingSystem getInstance() {
        if (instance == null) {
            instance = new MovieTicketBookingSystem();
        }
        return instance;
    }

    public void addMovie(Movie movie) { 
        movies.put(movie.getId(), movie); 
        movie.notifyObservers(movie); // Explictly flush observers via Push
    }
    
    public void addTheater(Theater theater) { theaters.put(theater.getId(), theater); }
    public void addShow(Show show) {
        shows.put(show.getId(), show);
    }

    public Booking bookTickets(User user, Show show, List<String> seatIds, PricingStrategy pricingModel) throws Exception {
        List<Seat> selectedSeats = new ArrayList<>();
        double totalPrice = 0.0;

        for (String seatId : seatIds) {
            Seat seat = show.getSeats().get(seatId);
            if (seat == null) throw new Exception("Seat " + seatId + " does not exist");
            if (seat.getStatus() != SeatStatus.AVAILABLE) throw new Exception("Seat " + seatId + " is not available");
            
            selectedSeats.add(seat);
            // Dynamic Pricing Strategy Evaluated Real Time!
            totalPrice += pricingModel.calculatePrice(seat.getPrice());
        }

        for (Seat seat : selectedSeats) {
            seat.setStatus(SeatStatus.BOOKED);
        }

        String bookingId = "BKG-" + System.currentTimeMillis();
        Booking booking = new Booking(bookingId, user, show, selectedSeats, totalPrice, BookingStatus.PENDING);
        bookings.put(bookingId, booking);
        
        return booking;
    }

    public void confirmBooking(String bookingId) throws Exception {
        Booking booking = bookings.get(bookingId);
        if (booking == null) throw new Exception("Booking not found");
        if (booking.getStatus() == BookingStatus.PENDING) {
            booking.setStatus(BookingStatus.CONFIRMED);
        } else {
            throw new Exception("Booking is not pending");
        }
    }

    public void cancelBooking(String bookingId) throws Exception {
        Booking booking = bookings.get(bookingId);
        if (booking == null) throw new Exception("Booking not found");
        if (booking.getStatus() == BookingStatus.CANCELLED) throw new Exception("Already cancelled");

        for (Seat seat : booking.getSeats()) {
            seat.setStatus(SeatStatus.AVAILABLE);
        }
        booking.setStatus(BookingStatus.CANCELLED);
    }
}
