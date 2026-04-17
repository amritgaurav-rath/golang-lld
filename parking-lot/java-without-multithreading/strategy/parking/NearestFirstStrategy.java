package strategy.parking;
import entities.Level;
import entities.ParkingSpot;
import entities.Vehicle;
import java.util.List;
import java.util.Optional;

public class NearestFirstStrategy implements ParkingStrategy {
    @Override
    public Optional<ParkingSpot> findSpot(List<Level> levels, Vehicle vehicle) {
        for (Level level : levels) {
            if (level.getAvailableSpots().get(vehicle.getType()) > 0) {
                for (ParkingSpot spot : level.getSpots()) {
                    if (!spot.isOccupied() && spot.getType() == vehicle.getType()) {
                        return Optional.of(spot);
                    }
                }
            }
        }
        return Optional.empty();
    }
}
