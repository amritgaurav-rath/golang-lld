/**
 * Represents a passenger request for an elevator.
 */
public class Request {
    private final int sourceFloor;
    private final int destinationFloor;
    private final int passengers;

    public Request(int sourceFloor, int destinationFloor, int passengers) {
        this.sourceFloor = sourceFloor;
        this.destinationFloor = destinationFloor;
        this.passengers = passengers;
    }

    public int getSourceFloor() {
        return sourceFloor;
    }

    public int getDestinationFloor() {
        return destinationFloor;
    }

    public int getPassengers() {
        return passengers;
    }
}
