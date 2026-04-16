import java.util.ArrayList;
import java.util.List;

public class ElevatorController {
    private final List<Elevator> elevators;

    public ElevatorController(int numElevators, int capacity) {
        elevators = new ArrayList<>();
        for (int i = 0; i < numElevators; i++) {
            elevators.add(new Elevator("E" + (i + 1), capacity));
        }
    }

    public void requestElevator(Request req) throws Exception {
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

    public boolean tickAll() {
        boolean isActive = false;
        for (Elevator el : elevators) {
            el.tick();
            if (el.getCurrentDirection() != Direction.IDLE) {
                isActive = true;
            }
        }
        return isActive;
    }
}
