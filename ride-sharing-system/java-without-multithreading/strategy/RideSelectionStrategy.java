package strategy;
import entities.Ride;
import java.util.List;
import java.util.Optional;

public interface RideSelectionStrategy {
    Optional<Ride> selectRide(List<Ride> availableRides);
}
