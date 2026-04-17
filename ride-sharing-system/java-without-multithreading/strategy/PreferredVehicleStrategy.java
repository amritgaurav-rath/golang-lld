package strategy;
import entities.Ride;
import enums.VehicleType;
import java.util.List;
import java.util.Optional;

public class PreferredVehicleStrategy implements RideSelectionStrategy {
    private VehicleType preferredType;

    public PreferredVehicleStrategy(VehicleType preferredType) {
        this.preferredType = preferredType;
    }

    @Override
    public Optional<Ride> selectRide(List<Ride> availableRides) {
        for (Ride ride : availableRides) {
            if (ride.getVehicle().getType() == preferredType) {
                return Optional.of(ride);
            }
        }
        return Optional.empty(); // Returns safely if mathematically isolated
    }
}
