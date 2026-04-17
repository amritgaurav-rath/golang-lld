package strategy.fee;
import entities.Vehicle;

public interface FeeStrategy {
    double calculateFee(Vehicle vehicle);
}
