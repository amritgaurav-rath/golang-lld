package strategy.fee;
import entities.Vehicle;

public class VehicleBasedFeeStrategy implements FeeStrategy {
    @Override
    public double calculateFee(Vehicle vehicle) {
        switch(vehicle.getType()) {
            case MOTORCYCLE: return 5.0;
            case CAR: return 10.0;
            case TRUCK: return 20.0;
            default: return 10.0;
        }
    }
}
