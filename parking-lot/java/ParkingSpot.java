public class ParkingSpot {
    private final String id;
    private final VehicleType type;
    private final int levelId;
    private boolean isOccupied;
    private Vehicle parkedVehicle;

    public ParkingSpot(String id, VehicleType type, int levelId) {
        this.id = id;
        this.type = type;
        this.levelId = levelId;
        this.isOccupied = false;
        this.parkedVehicle = null;
    }

    public synchronized boolean park(Vehicle v) {
        if (isOccupied) return false;
        this.parkedVehicle = v;
        this.isOccupied = true;
        return true;
    }

    public synchronized void unpark() {
        this.parkedVehicle = null;
        this.isOccupied = false;
    }

    public String getId() {
        return id;
    }

    public VehicleType getType() {
        return type;
    }

    public int getLevelId() {
        return levelId;
    }

    public boolean isOccupied() {
        return isOccupied;
    }

    public Vehicle getParkedVehicle() {
        return parkedVehicle;
    }
}
