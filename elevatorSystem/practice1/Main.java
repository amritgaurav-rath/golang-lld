/*
    1. Classes:
    Elevator: id, capacity, currentFloor, direction, Set<Request> requests
    Request: id, reqDirection, destFloor
    ElevatorController: 
        - List<Elevator>, 
        + requestElevator(floor, direction),
        + 
        Constructor: ElevatorController
    Direction:
        UP
        DOWN
        IDLE
 */

package elevatorSystem.practice1;

public class Main {
    public static void main(String args[]) {
        ElevatorController controller = new ElevatorController(3,10);

        controller.requestElevator(3, 5);
        controller.requestElevator(7, 9);
        controller.requestElevator(3, 0);
        controller.requestElevator(9, 2);

        System.out.println("Starting Simulation...");
        for(int i = 0; i < 20; i++) {
            System.out.println("--- TICK " + i + " ---");
            controller.step();
        }
    }
}