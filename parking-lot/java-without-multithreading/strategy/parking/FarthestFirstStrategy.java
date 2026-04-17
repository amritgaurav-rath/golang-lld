package strategy.parking;
import entities.Level;
import entities.ParkingSpot;
import entities.Vehicle;
import java.util.List;
import java.util.Optional;

public class FarthestFirstStrategy implements ParkingStrategy {
    @Override
    public Optional<ParkingSpot> findSpot(List<Level> levels, Vehicle vehicle) {
        // Reverse array traversal natively!
        for (int i = levels.size() - 1; i >= 0; i--) {
            Level level = levels.get(i);
            if (level.getAvailableSpots().get(vehicle.getType()) > 0) {
                List<ParkingSpot> spots = level.getSpots();
                for (int j = spots.size() - 1; j >= 0; j--) {
                    ParkingSpot spot = spots.get(j);
                    if (!spot.isOccupied() && spot.getType() == vehicle.getType()) {
                        return Optional.of(spot);
                    }
                }
            }
        }
        return Optional.empty();
    }
}
