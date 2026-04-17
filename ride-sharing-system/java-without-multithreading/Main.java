import entities.*;
import enums.*;
import strategy.*;
import services.RideSharingSystem;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Initializing S.O.L.I.D. FAANG SDE II Ride-Sharing Platform...");

        RideSharingSystem sys = RideSharingSystem.getInstance();

        // 1. Onboard Nodes natively
        sys.addUser("U1", "Rohan");
        sys.addUser("U2", "Shashank");
        sys.addUser("U3", "Nandini");

        try {
            // 2. Add vehicles
            sys.addVehicle("U1", "KA-01-1234", VehicleType.POLO);
            sys.addVehicle("U2", "TS-05-6239", VehicleType.ACTIVA);
            sys.addVehicle("U3", "MH-12-876D", VehicleType.XUV);

            System.out.println("✅ Users and their Vehicles safely mapped.\n");

            // 3. Offer Rides
            sys.offerRide("R1", "U1", "KA-01-1234", "Hyderabad", "Bangalore", 1);
            sys.offerRide("R2", "U2", "TS-05-6239", "Bangalore", "Mysore", 1);
            sys.offerRide("R3", "U3", "MH-12-876D", "Hyderabad", "Bangalore", 4);

            System.out.println("✅ Rides flawlessly created globally.\n");

            // Rule check: Try offering duplicated vehicle actively
            try {
                sys.offerRide("R4", "U1", "KA-01-1234", "Bangalore", "Pune", 2);
            } catch (Exception e) {
                System.out.println("🔒 Expected Validation mathematical drop: " + e.getMessage());
            }

            // 4. Select Rides natively mapped utilizing dynamically executed strategies
            System.out.println("\n--- Invoking Strategy Routing Algorithms ---");
            
            // Nandini selects ride searching for an ACTIVA specifically (PreferredVehicleStrategy)
            try {
                Ride rideA = sys.selectRide("U3", "Bangalore", "Mysore", 1, new PreferredVehicleStrategy(VehicleType.ACTIVA));
                System.out.printf("✅ Nandini securely matched ACTIVA Ride '%s' via mathematical Strategy parameters!%n", rideA.getRideId());
                sys.endRide(rideA.getRideId());
            } catch (Exception e) {
                System.out.println("❌ Strategy Fallback: " + e.getMessage());
            }

            // Rohan dynamically loops for MostVacantStrategy natively!
            try {
                Ride rideB = sys.selectRide("U1", "Hyderabad", "Bangalore", 2, new MostVacantStrategy());
                System.out.printf("✅ Rohan mathematically locked largest available structural mapping: Ride '%s'! Capacity remaining natively: %d!%n", 
                    rideB.getRideId(), rideB.getAvailableSeats());
                sys.endRide(rideB.getRideId());
            } catch (Exception e) {
                System.out.println("❌ Match fail cleanly: " + e.getMessage());
            }

            // 5. Native Statistical dumps
            sys.printRideStats();

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
