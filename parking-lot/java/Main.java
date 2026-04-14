import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class Main {
    private static final Object printLock = new Object();

    public static void printMessage(String msg) {
        synchronized (printLock) {
            System.out.println(msg);
        }
    }

    public static void main(String[] args) {
        printMessage("🚀 Initializing Parking Lot Simulation...");

        // Initialize 2 levels for the parking lot
        // Level 1: 5 Motorcycle, 5 Car, 2 Truck spots
        Level level1 = new Level(1, 5, 5, 2);
        // Level 2: 3 Motorcycle, 4 Car, 1 Truck spots
        Level level2 = new Level(2, 3, 4, 1);

        List<Level> levels = new ArrayList<>();
        levels.add(level1);
        levels.add(level2);

        ParkingLot lot = new ParkingLot(levels);

        lot.printAvailability();

        int numVehicles = 30;
        printMessage("Simulating " + numVehicles + " vehicles arriving concurrently...\n");

        List<Thread> threadPool = new ArrayList<>();

        for (int i = 1; i <= numVehicles; i++) {
            final int vehicleId = i;
            Thread t = new Thread(() -> {
                Random rand = new Random();
                
                // Randomly pick a vehicle type
                VehicleType[] vTypes = {VehicleType.MOTORCYCLE, VehicleType.CAR, VehicleType.TRUCK};
                VehicleType vType = vTypes[rand.nextInt(vTypes.length)];
                
                // Format ID nicely (e.g. VEH-001)
                String idStr = String.format("VEH-%03d", vehicleId);
                
                Vehicle v = new Vehicle(idStr, vType);

                try {
                    // Add a slight jitter to simulate staggered parallel arrivals
                    Thread.sleep(rand.nextInt(11));

                    ParkingSpot spot = lot.parkVehicle(v);
                    if (spot == null) {
                        printMessage("❌ " + v.getLicensePlate() + " (" + v.getType() + ") rejected: parking lot is full");
                        return;
                    }

                    printMessage("✅ " + v.getLicensePlate() + " (" + v.getType() + ") PARKED in spot " + spot.getId() + ".");

                    // Simulate time spent parked (10 to 50 milliseconds)
                    Thread.sleep(10 + rand.nextInt(41));

                    // Some vehicles decide to unpark (60% chance to leave)
                    if (rand.nextInt(10) < 6) {
                        Vehicle unparkedV = lot.unparkVehicle(spot.getId());
                        if (unparkedV == null) {
                            printMessage("⚠️  Error unparking " + v.getLicensePlate());
                        } else {
                            printMessage("👋 " + unparkedV.getLicensePlate() + " (" + unparkedV.getType() + ") LEFT spot " + spot.getId() + ".");
                        }
                    }

                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            });
            threadPool.add(t);
            t.start();
        }

        // Wait for all threads to finish
        for (Thread t : threadPool) {
            try {
                t.join();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }

        printMessage("\n🏁 Simulation completed.");
        lot.printAvailability();
    }
}
