package entities;
import enums.VehicleType;

public class ParkingSpot {
    private String id;
    private VehicleType type;
    private boolean isOccupied;
    private Vehicle vehicle;

    public ParkingSpot(String id, VehicleType type) {
        this.id = id;
        this.type = type;
        this.isOccupied = false;
        this.vehicle = null;
    }

    public String getId() { return id; }
    public VehicleType getType() { return type; }
    public boolean isOccupied() { return isOccupied; }
    public Vehicle getVehicle() { return vehicle; }

    public void park(Vehicle vehicle) throws Exception {
        if (this.isOccupied) throw new Exception("Spot mapping internally compromised (Already Occupied).");
        this.vehicle = vehicle;
        this.isOccupied = true;
    }

    public void unpark() {
        this.vehicle = null;
        this.isOccupied = false;
    }
}
