package entities;
import enums.VehicleType;
import java.util.*;

public class Level {
    private int id;
    private List<ParkingSpot> spots;
    private Map<VehicleType, Integer> availableSpots;

    public Level(int id, int numMotorcycleSpots, int numCarSpots, int numTruckSpots) {
        this.id = id;
        this.spots = new ArrayList<>();
        this.availableSpots = new HashMap<>();

        availableSpots.put(VehicleType.MOTORCYCLE, numMotorcycleSpots);
        availableSpots.put(VehicleType.CAR, numCarSpots);
        availableSpots.put(VehicleType.TRUCK, numTruckSpots);

        int spotNum = 1;
        for(int i=0; i<numMotorcycleSpots; i++) { spots.add(new ParkingSpot(String.format("L%d-M%d", id, spotNum++), VehicleType.MOTORCYCLE)); }
        spotNum = 1;
        for(int i=0; i<numCarSpots; i++) { spots.add(new ParkingSpot(String.format("L%d-C%d", id, spotNum++), VehicleType.CAR)); }
        spotNum = 1;
        for(int i=0; i<numTruckSpots; i++) { spots.add(new ParkingSpot(String.format("L%d-T%d", id, spotNum++), VehicleType.TRUCK)); }
    }

    public int getId() { return id; }
    public List<ParkingSpot> getSpots() { return spots; }
    public Map<VehicleType, Integer> getAvailableSpots() { return availableSpots; }

    public void modifyAvailability(VehicleType type, int delta) {
        availableSpots.put(type, availableSpots.get(type) + delta);
    }
}
