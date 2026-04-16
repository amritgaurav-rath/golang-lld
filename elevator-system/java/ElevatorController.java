import java.util.ArrayList;
import java.util.List;

/**
 * Manages a fleet of concurrent elevators, distributing requests optimally.
 */
public class ElevatorController {
    private final List<Elevator> elevators;

    public ElevatorController(int numElevators, int capacity) {
        elevators = new ArrayList<>();
        for (int i = 0; i < numElevators; i++) {
            Elevator el = new Elevator("E" + (i + 1), capacity);
            elevators.add(el);
            // Spin up the background thread for this elevator immediately
            new Thread(el, "Thread-" + el.getId()).start();
        }
    }

    /**
     * Dispatcher method to find the closest suitable elevator sequence concurrently.
     * Thanks to the synchronized methods within the Elevator class, polling for
     * getDistanceIfAssigned() guarantees atomic, accurate state retrieval.
     */
    public synchronized void requestElevator(Request req) throws Exception {
        Elevator optimalElevator = null;
        int minDistance = Integer.MAX_VALUE;

        for (Elevator el : elevators) {
            int dist = el.getDistanceIfAssigned(req);
            if (dist < minDistance) {
                minDistance = dist;
                optimalElevator = el;
            }
        }

        if (optimalElevator == null) {
            throw new Exception("all elevators are currently over capacity to handle " + req.getPassengers() + " passengers");
        }

        System.out.printf("[Dispatcher] \uD83D\uDFE2 Routing (Src: %d -> Dest: %d, Pax: %d) into Optimal => %s\n",
                req.getSourceFloor(), req.getDestinationFloor(), req.getPassengers(), optimalElevator.getId());
        
        optimalElevator.addRequest(req);
    }
    
    /**
     * Initiates a graceful shutdown of all elevator threads.
     */
    public void shutdown() {
        for (Elevator el : elevators) {
            el.stopElevator();
        }
    }
}
