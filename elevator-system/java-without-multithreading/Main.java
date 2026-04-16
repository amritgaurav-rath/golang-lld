import java.util.ArrayList;
import java.util.List;

/**
 * Entry point for the synchronous elevator system simulation.
 */
public class Main {
    public static void main(String[] args) {
        System.out.println("\uD83D\uDE80 Initializing Synchronous (Single-Threaded) Elevator System in Java...");

        ElevatorController controller = new ElevatorController(3, 10);

        List<Request> requests = new ArrayList<>();
        requests.add(new Request(0, 5, 2));
        requests.add(new Request(2, 8, 4));
        requests.add(new Request(8, 1, 3));
        requests.add(new Request(3, 6, 5));
        requests.add(new Request(6, 0, 2));
        requests.add(new Request(0, 10, 9));

        System.out.println("Sequentially routing all user requests immediately...");

        for (int i = 0; i < requests.size(); i++) {
            try {
                controller.requestElevator(requests.get(i));
            } catch (Exception e) {
                System.out.printf("❌ Request %d Failed: %s\n", i + 1, e.getMessage());
            }
        }

        System.out.println("\nExecuting deterministic tick loop to spin Elevator motors physically...");

        int step = 1;
        while (true) {
            boolean active = controller.tickAll();
            if (!active) {
                break;
            }

            System.out.printf("--- Time Step %d ---\n", step);
            step++;
            
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }

        System.out.println("\n\uD83C\uDFC1 Synchronous Elevator Simulation Completed.");
    }
}
