package elevatorSystem.practice1;
import java.util.*;

public class Elevator {
    private int id;
    private int currentFloor;
    private Direction direction;
    private Set<Request> requests;

    // defining constructor
    public Elevator(int id) {
        this.id = id;
        this.currentFloor = 0;
        this.direction = Direction.IDLE;
        this.requests = new HashSet<>();
    }

    public int getId() {
        return this.id;
    }

    public Direction getDirection() {
        return this.direction;
    }

    public void setDirection(Direction direction) {
        this.direction = direction;
    }

    public int getFloor() {
        return this.currentFloor;
    }

    public boolean addRequest(Request req) {
        this.requests.add(req);

        return true;
    }

    public void step() {
        if (this.requests.isEmpty()) {
            this.direction = Direction.IDLE;
            return;
        }
        if (this.direction == Direction.IDLE) {
            Request nearest = findNearestRequest();
            this.direction = nearest.getFloor() > this.getFloor() ? Direction.UP : Direction.DOWN;
        }

        processStops();

        if(!hasAnyRequestAhead(this.requests)) {
            this.direction = this.getDirection() == Direction.UP ? Direction.DOWN : Direction.UP;
            processStops(); // Re-evaluate stops for the brand new direction!
        }

        if (this.requests.isEmpty()) {
            this.direction = Direction.IDLE;
        } else {
            if (this.direction == Direction.UP) {
                this.currentFloor++;
            } 
            if (this.direction == Direction.DOWN) {
                this.currentFloor--;
            }
        }
    }

    private void processStops() {
        boolean stopped = false;
        List<Request> toRemove = new ArrayList<>(); // Your temporary trash can

        for(Request request : this.requests) {
            if(request.getFloor() == this.currentFloor && 
               ((request.getType()==RequestType.PICKUP_UP && this.direction == Direction.UP) || 
                (request.getType()==RequestType.PICKUP_DOWN && this.direction == Direction.DOWN))) {
                
                System.out.println("Elevator " + this.id + " Picked up Passenger at Floor " + this.currentFloor);
                request.setFloorAsDestination(); // Passenger is inside, morph this request into a DESTINATION
                stopped = true;
                
            } else if (request.getFloor() == this.currentFloor && request.getType() == RequestType.DESTINATION) {
                toRemove.add(request); // Put it in the trash can securely
                stopped = true;                
            }
        }
        
        this.requests.removeAll(toRemove); // Safely delete them all at once!

        if (stopped) {
            stop();
        }
    }

    private void stop() {
        System.out.println("Elevator " + this.id + " stopped at floor " + this.currentFloor);
    }

    private boolean hasAnyRequestAhead(Set<Request> requests) {
       for(Request req:requests) {
        if(req.getFloor() > this.getFloor() && this.getDirection() == Direction.UP) {
            return true;
        }
        if(req.getFloor() < this.getFloor() && this.getDirection() == Direction.DOWN) {
            return true;
        }
       }
       return false;
    }

    private Request findNearestRequest() {
        Request nearest = null;
        int minDist = Integer.MAX_VALUE;

        for(Request request:this.requests) {
            int dist = Math.abs(request.getFloor() - this.getFloor());
            if(dist<minDist) {
                nearest = request;
            }
        }

        return nearest;
    }
}
