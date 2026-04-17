package strategy.parking;
import entities.Level;
import entities.ParkingSpot;
import entities.Vehicle;
import java.util.List;
import java.util.Optional;

public interface ParkingStrategy {
    Optional<ParkingSpot> findSpot(List<Level> levels, Vehicle vehicle);
}
