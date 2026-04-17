package strategy;
import entities.Ride;
import java.util.List;
import java.util.Optional;

public class MostVacantStrategy implements RideSelectionStrategy {
    @Override
    public Optional<Ride> selectRide(List<Ride> availableRides) {
        Ride mostVacant = null;
        int maxSeats = -1;

        for (Ride ride : availableRides) {
            if (ride.getAvailableSeats() > maxSeats) {
                maxSeats = ride.getAvailableSeats();
                mostVacant = ride;
            }
        }
        
        return Optional.ofNullable(mostVacant);
    }
}
