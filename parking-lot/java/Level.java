import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Level {
    private final int id;
    private final List<ParkingSpot> spots;
    private final Map<VehicleType, Integer> availableSpots;
    private final Lock lock;

    public Level(int id, int numMotorcycleSpots, int numCarSpots, int numTruckSpots) {
        this.id = id;
        this.spots = new ArrayList<>();
        this.availableSpots = new ConcurrentHashMap<>();
        this.lock = new ReentrantLock();

        availableSpots.put(VehicleType.MOTORCYCLE, numMotorcycleSpots);
        availableSpots.put(VehicleType.CAR, numCarSpots);
        availableSpots.put(VehicleType.TRUCK, numTruckSpots);

        int spotNum = 1;
        for (int i = 0; i < numMotorcycleSpots; i++) {
            spots.add(new ParkingSpot("L" + id + "-M" + (spotNum++), VehicleType.MOTORCYCLE, id));
        }

        spotNum = 1;
        for (int i = 0; i < numCarSpots; i++) {
            spots.add(new ParkingSpot("L" + id + "-C" + (spotNum++), VehicleType.CAR, id));
        }

        spotNum = 1;
        for (int i = 0; i < numTruckSpots; i++) {
            spots.add(new ParkingSpot("L" + id + "-T" + (spotNum++), VehicleType.TRUCK, id));
        }
    }

    public ParkingSpot parkVehicle(Vehicle v) {
        lock.lock();
        try {
            if (availableSpots.getOrDefault(v.getType(), 0) <= 0) {
                return null;
            }

            for (ParkingSpot spot : spots) {
                if (!spot.isOccupied() && spot.getType() == v.getType()) {
                    if (spot.park(v)) {
                        availableSpots.put(v.getType(), availableSpots.get(v.getType()) - 1);
                        return spot;
                    }
                }
            }
            return null;
        } finally {
            lock.unlock();
        }
    }

    public Vehicle unparkVehicle(String spotId) {
        lock.lock();
        try {
            for (ParkingSpot spot : spots) {
                if (spot.getId().equals(spotId)) {
                    if (!spot.isOccupied()) {
                        return null;
                    }
                    Vehicle v = spot.getParkedVehicle();
                    spot.unpark();
                    availableSpots.put(v.getType(), availableSpots.get(v.getType()) + 1);
                    return v;
                }
            }
            return null;
        } finally {
            lock.unlock();
        }
    }

    public Map<VehicleType, Integer> getAvailability() {
        // Return a copy to ensure thread safety
        return new ConcurrentHashMap<>(availableSpots);
    }

    public int getId() {
        return id;
    }
}
