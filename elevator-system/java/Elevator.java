import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Elevator {
    private final String id;
    private int currentFloor;
    private Direction currentDirection;
    private final int capacity;
    private int currentLoad;

    private final Map<Integer, Boolean> upStops;
    private final Map<Integer, Boolean> downStops;
    private List<Request> requests;

    public Elevator(String id, int capacity) {
        this.id = id;
        this.currentFloor = 0;
        this.currentDirection = Direction.IDLE;
        this.capacity = capacity;
        this.currentLoad = 0;
        this.upStops = new HashMap<>();
        this.downStops = new HashMap<>();
        this.requests = new ArrayList<>();
    }

    public String getId() {
        return id;
    }

    public Direction getCurrentDirection() {
        return currentDirection;
    }

    public void addRequest(Request req) throws Exception {
        if (currentLoad + req.getPassengers() > capacity) {
            throw new Exception("elevator " + id + " capacity exceeded");
        }

        requests.add(req);
        currentLoad += req.getPassengers();

        if (req.getSourceFloor() > currentFloor) {
            upStops.put(req.getSourceFloor(), true);
        } else if (req.getSourceFloor() < currentFloor) {
            downStops.put(req.getSourceFloor(), true);
        } else {
            if (req.getDestinationFloor() > currentFloor) {
                upStops.put(req.getDestinationFloor(), true);
            } else {
                downStops.put(req.getDestinationFloor(), true);
            }
        }

        if (req.getDestinationFloor() > req.getSourceFloor()) {
            upStops.put(req.getDestinationFloor(), true);
        } else if (req.getDestinationFloor() < req.getSourceFloor()) {
            downStops.put(req.getDestinationFloor(), true);
        }

        if (currentDirection == Direction.IDLE) {
            if (req.getSourceFloor() > currentFloor || req.getDestinationFloor() > currentFloor) {
                currentDirection = Direction.UP;
            } else if (req.getSourceFloor() < currentFloor || req.getDestinationFloor() < currentFloor) {
                currentDirection = Direction.DOWN;
            }
        }
    }

    public int getDistanceIfAssigned(Request req) {
        if (currentLoad + req.getPassengers() > capacity) {
            return Integer.MAX_VALUE;
        }

        int dist = Math.abs(currentFloor - req.getSourceFloor());

        if (currentDirection == Direction.IDLE) {
            return dist;
        }
        if (currentDirection == Direction.UP && req.getSourceFloor() >= currentFloor) {
            return dist;
        }
        if (currentDirection == Direction.DOWN && req.getSourceFloor() <= currentFloor) {
            return dist;
        }

        return dist + 1000;
    }

    private void processCurrentFloor() {
        if (currentDirection == Direction.UP) {
            if (upStops.containsKey(currentFloor) && upStops.get(currentFloor)) {
                System.out.printf("   -> Elevator %s \uD83D\uDED1 STOPPED at floor %d (UP)\n", id, currentFloor);
                upStops.remove(currentFloor);
                dropOffPassengers();
            }
        } else if (currentDirection == Direction.DOWN) {
            if (downStops.containsKey(currentFloor) && downStops.get(currentFloor)) {
                System.out.printf("   -> Elevator %s \uD83D\uDED1 STOPPED at floor %d (DOWN)\n", id, currentFloor);
                downStops.remove(currentFloor);
                dropOffPassengers();
            }
        }
    }

    private void dropOffPassengers() {
        List<Request> remaining = new ArrayList<>();
        for (Request req : requests) {
            if (req.getDestinationFloor() == currentFloor) {
                currentLoad -= req.getPassengers();
                System.out.printf("      -> %d passengers ALIGHTED at floor %d from elevator %s\n", 
                        req.getPassengers(), currentFloor, id);
            } else {
                remaining.add(req);
            }
        }
        requests = remaining;
    }

    private Direction getNextDirection() {
        if (currentDirection == Direction.UP) {
            for (int floor : upStops.keySet()) {
                if (floor > currentFloor) {
                    return Direction.UP;
                }
            }
            if (!downStops.isEmpty()) {
                return Direction.DOWN;
            }
        } else if (currentDirection == Direction.DOWN) {
            for (int floor : downStops.keySet()) {
                if (floor < currentFloor) {
                    return Direction.DOWN;
                }
            }
            if (!upStops.isEmpty()) {
                return Direction.UP;
            }
        }

        if (!upStops.isEmpty()) {
            return Direction.UP;
        } else if (!downStops.isEmpty()) {
            return Direction.DOWN;
        }

        return Direction.IDLE;
    }

    public void tick() {
        Direction nextDir = getNextDirection();
        currentDirection = nextDir;

        if (currentDirection == Direction.IDLE) {
            return;
        }

        if (currentDirection == Direction.UP) {
            currentFloor++;
        } else if (currentDirection == Direction.DOWN) {
            currentFloor--;
        }

        processCurrentFloor();
    }
}
