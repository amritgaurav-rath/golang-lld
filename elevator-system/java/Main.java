import java.util.Random;

/**
 * Entry point for testing the multi-threaded concurrent elevator system simulation.
 */
public class Main {
    public static void main(String[] args) {
        System.out.println("\uD83D\uDE80 Initializing Multi-Threaded Elevator System in Java...");

        // 3 Elevators with capacity of 10 passengers each.
        // It immediately boots up threads for each motor.
        ElevatorController controller = new ElevatorController(3, 10);

        System.out.println("Starting concurrent request generator...\n");

        // Start a thread to act as users randomly pushing buttons
        Thread requestGenerator = new Thread(() -> {
            int[][] reqsData = {
                {0, 5, 2},
                {2, 8, 4},
                {8, 1, 3},
                {3, 6, 5},
                {6, 0, 2},
                {0, 10, 9}
            };
            
            Random rand = new Random();

            for (int i = 0; i < reqsData.length; i++) {
                try {
                    int[] data = reqsData[i];
                    Request req = new Request(data[0], data[1], data[2]);
                    controller.requestElevator(req);
                    
                    // Simulate random arrival times
                    Thread.sleep(rand.nextInt(1500) + 500); 
                } catch (Exception e) {
                    System.out.printf("❌ Request Failed: %s\n", e.getMessage());
                }
            }
            System.out.println("\n✅ All user requests have been generated.");
        });

        requestGenerator.start();

        // Let the simulation run for a bit so elevators can fulfill all queues
        try {
            requestGenerator.join(); // wait for all requests to be generated
            System.out.println("⏳ Waiting for elevators to finish processing active queues...\n");
            
            // Wait to allow real-time motion and fulfillment
            Thread.sleep(12000); 
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }

        System.out.println("\n\uD83D\uDED1 Shutting down Elevators...");
        controller.shutdown();
        
        // Final sleep to let threads cleanly print their exit message
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }

        System.out.println("\n\uD83C\uDFC1 Multi-Threaded Elevator Simulation Completed.");
    }
}
