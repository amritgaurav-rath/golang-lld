package strategy.fee;
import entities.Vehicle;

public class FlatRateFeeStrategy implements FeeStrategy {
    @Override
    public double calculateFee(Vehicle vehicle) {
        return 10.0;
    }
}
