import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Represents a single concurrent elevator, managing its state and processing requests.
 */
public class Elevator implements Runnable {
    private final String id;
    private int currentFloor;
    private Direction currentDirection;
    private final int capacity;
    private int currentLoad;

    private final Map<Integer, Boolean> upStops;
    private final Map<Integer, Boolean> downStops;
    private List<Request> requests;
    
    private volatile boolean running;

    public Elevator(String id, int capacity) {
        this.id = id;
        this.currentFloor = 0;
        this.currentDirection = Direction.IDLE;
        this.capacity = capacity;
        this.currentLoad = 0;
        this.upStops = new HashMap<>();
        this.downStops = new HashMap<>();
        this.requests = new ArrayList<>();
        this.running = true;
    }

    public String getId() {
        return id;
    }

    public synchronized Direction getCurrentDirection() {
        return currentDirection;
    }
    
    public void stopElevator() {
        running = false;
    }

    /**
     * Synchronized evaluation to determine if this elevator is suitable for a request concurrently.
     * By making this method synchronized, we ensure the dispatcher reads an accurate snapshot
     * of currentLoad, currentFloor, and currentDirection without encountering race conditions.
     */
    public synchronized int getDistanceIfAssigned(Request req) {
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

    /**
     * Adds a request to the elevator's internal queue in a thread-safe manner.
     * After updating the required stops, it uses notifyAll() to wake up the elevator's run()
     * loop if it was currently waiting (sleeping) in an IDLE state.
     */
    public synchronized void addRequest(Request req) throws Exception {
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
        
        // Notify the elevator thread that new work arrived
        notifyAll();
    }

    /**
     * The main execution loop for the Elevator thread, running asynchronously.
     */
    @Override
    public void run() {
        while (running) {
            synchronized (this) {
                // If there's no work, pause the thread execution using wait() to save CPU cycles.
                // The thread will remain asleep until addRequest() invokes notifyAll().
                if (currentDirection == Direction.IDLE && upStops.isEmpty() && downStops.isEmpty()) {
                    try {
                        wait(500); // Wait for new requests to arrive, timeout to check 'running' flag periodically
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        break;
                    }
                    continue; // Re-evaluate after wake up
                }

                Direction nextDir = getNextDirection();
                currentDirection = nextDir;

                if (currentDirection != Direction.IDLE) {
                    if (currentDirection == Direction.UP) {
                        currentFloor++;
                    } else if (currentDirection == Direction.DOWN) {
                        currentFloor--;
                    }
                }
                
                processCurrentFloor();
            }

            // Simulate the physical time it takes to move between floors outside the synchronized block
            // to allow dispatchers to add requests while it moves.
            try {
                Thread.sleep(500);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
        System.out.printf("   [SYSTEM] Elevator %s has safely shut down.\n", id);
    }

    private synchronized void processCurrentFloor() {
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

    private synchronized void dropOffPassengers() {
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

    private synchronized Direction getNextDirection() {
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
}
