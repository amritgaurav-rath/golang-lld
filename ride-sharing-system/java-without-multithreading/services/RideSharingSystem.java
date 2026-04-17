package services;
import entities.*;
import enums.*;
import strategy.*;
import java.util.*;

public class RideSharingSystem {
    private Map<String, User> users;
    private Map<String, Vehicle> vehicles;
    private Map<String, Ride> rides;

    private static RideSharingSystem instance;

    private RideSharingSystem() {
        this.users = new HashMap<>();
        this.vehicles = new HashMap<>();
        this.rides = new HashMap<>();
    }

    public static synchronized RideSharingSystem getInstance() {
        if (instance == null) {
            instance = new RideSharingSystem();
        }
        return instance;
    }

    public void addUser(String id, String name) {
        users.put(id, new User(id, name));
    }

    public void addVehicle(String userId, String regNo, VehicleType type) throws Exception {
        if (!users.containsKey(userId)) throw new Exception("User strictly unmapped natively.");
        vehicles.put(regNo, new Vehicle(userId, regNo, type));
    }

    public void offerRide(String rideId, String userId, String regNo, String origin, String destination, int availableSeats) throws Exception {
        if (!users.containsKey(userId)) throw new Exception("Driver actively unmapped.");
        if (!vehicles.containsKey(regNo)) throw new Exception("Vehicle physically unmapped.");
        
        Vehicle vehicle = vehicles.get(regNo);
        if (!vehicle.getOwnerId().equals(userId)) throw new Exception("User mathematically does not own the mapped vehicle.");

        // Rule: Limit vehicle to exactly one active ride natively.
        for (Ride existingRide : rides.values()) {
            if (existingRide.getVehicle().getRegistrationNo().equals(regNo) && existingRide.getStatus() == RideStatus.OFFERED) {
                throw new Exception("Vehicle cleanly locked natively inside another active ride mapping.");
            }
        }

        Ride ride = new Ride(rideId, userId, vehicle, origin, destination, availableSeats);
        rides.put(rideId, ride);
    }

    public Ride selectRide(String userId, String source, String destination, int seats, RideSelectionStrategy strategy) throws Exception {
        if (seats < 1 || seats > 2) throw new Exception("Ride request limits strictly enforce cleanly 1 or 2 seats natively.");
        if (!users.containsKey(userId)) throw new Exception("Passenger mathematically unmapped.");

        List<Ride> availableRides = new ArrayList<>();
        for (Ride ride : rides.values()) {
            if (ride.getStatus() == RideStatus.OFFERED &&
                ride.getOrigin().equals(source) &&
                ride.getDestination().equals(destination) &&
                ride.getAvailableSeats() >= seats) {
                availableRides.add(ride);
            }
        }

        if (availableRides.isEmpty()) {
            throw new Exception("No active structures natively map to these origin variables.");
        }

        Optional<Ride> selectedRideOpt = strategy.selectRide(availableRides);
        if (!selectedRideOpt.isPresent()) {
            throw new Exception("Strategy mathematically failed mapping exact constraints.");
        }

        Ride selectedRide = selectedRideOpt.get();
        selectedRide.decreaseSeats(seats);
        
        // Push consumer array mappings
        users.get(userId).incrementTakenRides();

        return selectedRide;
    }

    public void endRide(String rideId) throws Exception {
        if (!rides.containsKey(rideId)) throw new Exception("Ride natively unmapped.");
        Ride ride = rides.get(rideId);
        if (ride.getStatus() == RideStatus.COMPLETED) throw new Exception("Already safely marked Completed.");

        ride.setStatus(RideStatus.COMPLETED);
        
        // Push mathematical driver points only natively upon successful completion.
        users.get(ride.getDriverId()).incrementOfferedRides();
    }

    public void printRideStats() {
        System.out.println("\n--- Live SDE II Statistical Arrays ---");
        for (User user : users.values()) {
            System.out.printf("%s: %d Taken, %d Offered\n", user.getName(), user.getTakenRides(), user.getOfferedRides());
        }
        System.out.println("--------------------------------------\n");
    }
}
