import entities.*;
import enums.*;
import strategy.parking.*;
import strategy.fee.*;
import services.ParkingLotSystem;
import java.util.*;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Initializing FAANG Tier S.O.L.I.D. Parking Lot (Strategies Active)...");

        List<Level> levels = new ArrayList<>();
        levels.add(new Level(1, 5, 5, 2));
        levels.add(new Level(2, 3, 4, 1));
        
        ParkingLotSystem sys = ParkingLotSystem.getInstance();
        sys.setLevels(levels);
        
        // 1. Injecting dynamic FAANG pattern strategies
        sys.setParkingStrategy(new FarthestFirstStrategy()); // Inverts the generic search to fill backwards!
        sys.setFeeStrategy(new VehicleBasedFeeStrategy());

        sys.printAvailability();

        System.out.println("Alice sequence: Parks Car, Parks Motorcycle via FarthestFirstStrategy.");
        
        Vehicle aliceCar = new Vehicle("ALICE-01", VehicleType.CAR);
        Vehicle aliceMoto = new Vehicle("ALICE-02", VehicleType.MOTORCYCLE);
        
        try {
            ParkingSpot spot1 = sys.parkVehicle(aliceCar);
            System.out.println("✅ Alice successfully PARKED Car at: " + spot1.getId() + " (Expect Level 2!)");
            
            ParkingSpot spot2 = sys.parkVehicle(aliceMoto);
            System.out.println("✅ Alice successfully PARKED Motorcycle at: " + spot2.getId() + " (Expect Level 2!)");
            
            sys.printAvailability();

            System.out.println("Alice unparks Car...");
            Vehicle unparked = sys.unparkVehicle(spot1.getId());
            System.out.println("👋 Unparked safely: " + unparked.getLicensePlate());
            
            sys.printAvailability();

        } catch(Exception e) {
            System.out.println("❌ ERROR: " + e.getMessage());
        }
    }
}
