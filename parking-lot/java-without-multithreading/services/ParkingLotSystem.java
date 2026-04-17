package services;
import entities.*;
import enums.*;
import strategy.parking.*;
import strategy.fee.*;
import java.util.*;

public class ParkingLotSystem {
    private List<Level> levels;
    private static ParkingLotSystem instance;
    private ParkingStrategy parkingStrategy;
    private FeeStrategy feeStrategy;

    private ParkingLotSystem() {
        this.levels = new ArrayList<>();
        // Default strategies
        this.parkingStrategy = new NearestFirstStrategy();
        this.feeStrategy = new FlatRateFeeStrategy();
    }

    public static synchronized ParkingLotSystem getInstance() {
        if (instance == null) {
            instance = new ParkingLotSystem();
        }
        return instance;
    }
    
    public void setLevels(List<Level> levels) { this.levels = levels; }
    public void setParkingStrategy(ParkingStrategy parkingStrategy) { this.parkingStrategy = parkingStrategy; }
    public void setFeeStrategy(FeeStrategy feeStrategy) { this.feeStrategy = feeStrategy; }

    public ParkingSpot parkVehicle(Vehicle vehicle) throws Exception {
        // Natively delegating to Strategy mathematical evaluator
        Optional<ParkingSpot> spotOpt = parkingStrategy.findSpot(levels, vehicle);
        
        if (spotOpt.isPresent()) {
            ParkingSpot spot = spotOpt.get();
            spot.park(vehicle);
            
            // Adjust mathematical metric mappings
            for(Level level : levels) {
                if(level.getSpots().contains(spot)) {
                    level.modifyAvailability(vehicle.getType(), -1);
                    break;
                }
            }
            return spot;
        }

        throw new Exception("Parking lot is fully occupied for " + vehicle.getType());
    }

    public Vehicle unparkVehicle(String spotId) throws Exception {
        for (Level level : levels) {
            for (ParkingSpot spot : level.getSpots()) {
                if (spot.getId().equals(spotId)) {
                    if (!spot.isOccupied()) {
                        throw new Exception("Spot is already empty.");
                    }
                    Vehicle v = spot.getVehicle();
                    spot.unpark();
                    level.modifyAvailability(v.getType(), 1);
                    
                    // Dynamic math checkout processing
                    double fee = feeStrategy.calculateFee(v);
                    System.out.printf("💳 Fee computed for %s: $%.2f\n", v.getLicensePlate(), fee);
                    
                    return v;
                }
            }
        }
        throw new Exception("Spot ID " + spotId + " not found mapped.");
    }

    public void printAvailability() {
        System.out.println("\n--- Current Parking Availability ---");
        for (Level l : levels) {
            System.out.printf("Level %d -> Motorcycles: %d, Cars: %d, Trucks: %d\n", 
                l.getId(), l.getAvailableSpots().get(VehicleType.MOTORCYCLE),
                l.getAvailableSpots().get(VehicleType.CAR), l.getAvailableSpots().get(VehicleType.TRUCK));
        }
        System.out.println("------------------------------------\n");
    }
}
