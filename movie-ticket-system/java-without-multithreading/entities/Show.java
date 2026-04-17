package entities;
import java.util.Map;
import java.util.HashMap;

public class Show {
    private String id;
    private Movie movie;
    private Theater theater;
    private long startTime;
    private long endTime;
    private Map<String, Seat> seats;

    public Show(String id, Movie movie, Theater theater, long startTime, long endTime) {
        this.id = id;
        this.movie = movie;
        this.theater = theater;
        this.startTime = startTime;
        this.endTime = endTime;
        this.seats = new HashMap<>();
    }

    public String getId() { return id; }
    public Map<String, Seat> getSeats() { return seats; }
    public void addSeat(Seat seat) { this.seats.put(seat.getId(), seat); }
}
