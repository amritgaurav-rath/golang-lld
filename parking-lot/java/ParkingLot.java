import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ParkingLot {
    private final List<Level> levels;

    public ParkingLot(List<Level> levels) {
        this.levels = levels;
    }

    public ParkingSpot parkVehicle(Vehicle v) {
        for (Level level : levels) {
            Map<VehicleType, Integer> avail = level.getAvailability();
            if (avail.getOrDefault(v.getType(), 0) > 0) {
                ParkingSpot spot = level.parkVehicle(v);
                if (spot != null) {
                    return spot;
                }
            }
        }
        return null;
    }

    public Vehicle unparkVehicle(String spotId) {
        for (Level level : levels) {
            Vehicle v = level.unparkVehicle(spotId);
            if (v != null) {
                return v;
            }
        }
        return null;
    }

    public Map<VehicleType, Integer> getTotalAvailability() {
        Map<VehicleType, Integer> totalAvail = new ConcurrentHashMap<>();
        for (Level level : levels) {
            Map<VehicleType, Integer> avail = level.getAvailability();
            for (Map.Entry<VehicleType, Integer> entry : avail.entrySet()) {
                totalAvail.put(entry.getKey(), totalAvail.getOrDefault(entry.getKey(), 0) + entry.getValue());
            }
        }
        return totalAvail;
    }

    public void printAvailability() {
        System.out.println("--- Current Parking Availability ---");
        Map<VehicleType, Integer> avail = getTotalAvailability();
        for (Map.Entry<VehicleType, Integer> entry : avail.entrySet()) {
            if (entry.getValue() > 0 || entry.getKey() != VehicleType.UNKNOWN) {
                System.out.println(entry.getKey() + " Spots: " + entry.getValue());
            }
        }
        System.out.println("------------------------------------");
    }
}
