package elevatorSystem.practice1;
import java.util.*;

public class ElevatorController {
    private List<Elevator> elevators;

    public ElevatorController(int numElevators, int floors) {
        this.elevators = new ArrayList<>(); // You MUST initialize the list first!
        
        for(int i=0; i<numElevators; i++) {
            this.elevators.add(new Elevator(i)); // Calling the constructor safely!
        }
    }

    public boolean requestElevator(int floor, int destFloor) {
        if(floor<0 || floor >9 || destFloor < 0 || destFloor > 9) {
            return false;
        }

        Request request = new Request(floor, destFloor);

        Elevator bestElevator = findBestElevator(request);        

        // Force direction update before step() to allocate sequentially fairly!
        if (bestElevator.getDirection() == Direction.IDLE) {
            if (floor > bestElevator.getFloor()) {
                bestElevator.setDirection(Direction.UP);
            } else if (floor < bestElevator.getFloor()) {
                bestElevator.setDirection(Direction.DOWN);
            }
        }

        return bestElevator.addRequest(request);
    }

    private Elevator findBestElevator(Request request) {
        Elevator bestElevator = findMovingTowards(request);

        if (bestElevator != null) {
            return bestElevator;
        }

        bestElevator = findNearestIdleElevator(request);
        if (bestElevator != null) {
            return bestElevator;
        }

        return findNearestElevator(request);
    }

    private Elevator findMovingTowards(Request request) {
        Elevator nearest = null;
        int minDist = Integer.MAX_VALUE;

        for (Elevator elevator:elevators) {
            if(request.getType() == RequestType.PICKUP_UP && elevator.getDirection() == Direction.UP) {
                int dist = request.getFloor()-elevator.getFloor();

                if(dist >= 0 && dist<minDist) {
                    minDist = dist;
                    nearest = elevator;
                }
            } else if(request.getType() == RequestType.PICKUP_DOWN && elevator.getDirection() == Direction.DOWN) {
                int dist = elevator.getFloor()-request.getFloor();

                if(dist >= 0 && dist<minDist) {
                    minDist = dist;
                    nearest = elevator;
                }
            }
        }

        return nearest;
    }

    private Elevator findNearestIdleElevator(Request request) {
        Elevator nearest = null;
        int minDist = Integer.MAX_VALUE;

        for (Elevator elevator:elevators) {
            if(elevator.getDirection() == Direction.IDLE) {
                int dist = Math.abs(request.getFloor()-elevator.getFloor());

                if(dist >= 0 && dist<minDist) {
                    minDist = dist;
                    nearest = elevator;
                }
            }
        }

        return nearest;
    }

    private Elevator findNearestElevator(Request request) {
        int minDist = Integer.MAX_VALUE;
        Elevator nearest = null;

        for (Elevator elevator:elevators) {
            int dist = Math.abs(request.getFloor()-elevator.getFloor());

            if(dist >= 0 && dist<minDist) {
                minDist = dist;
                nearest = elevator;
            }
        }

        return nearest;
    }

    public void step() {
        for(Elevator elevator:elevators) {
            elevator.step();
        }
    }
}
