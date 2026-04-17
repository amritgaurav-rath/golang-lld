package entities;
import enums.RideStatus;

public class Ride {
    private String rideId;
    private String driverId;
    private Vehicle vehicle;
    private String origin;
    private String destination;
    private int availableSeats;
    private RideStatus status;

    public Ride(String rideId, String driverId, Vehicle vehicle, String origin, String destination, int availableSeats) {
        this.rideId = rideId;
        this.driverId = driverId;
        this.vehicle = vehicle;
        this.origin = origin;
        this.destination = destination;
        this.availableSeats = availableSeats;
        this.status = RideStatus.OFFERED;
    }

    public String getRideId() { return rideId; }
    public String getDriverId() { return driverId; }
    public Vehicle getVehicle() { return vehicle; }
    public String getOrigin() { return origin; }
    public String getDestination() { return destination; }
    public int getAvailableSeats() { return availableSeats; }
    public RideStatus getStatus() { return status; }
    
    public void setStatus(RideStatus status) { this.status = status; }
    public void decreaseSeats(int offset) { this.availableSeats -= offset; }
}
