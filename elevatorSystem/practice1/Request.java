package elevatorSystem.practice1;

public class Request {
    private int floor;
    private int destination;
    private RequestType type;

    public Request(int floor, int dest) {
        this.floor = floor;
        this.destination = dest;
        this.type = dest > floor ? RequestType.PICKUP_UP : RequestType.PICKUP_DOWN;
    }

    public int getFloor() {
        return this.floor;
    }

    public RequestType getType() {
        return this.type;
    }

    public void setFloorAsDestination() {
        this.floor = this.destination;
        this.type = RequestType.DESTINATION;
    }
}
