# Elevator System (Low Level Design)

This repository contains an implementation of a highly scalable Elevator Control System. The system manages multiple elevators across multiple floors, efficiently processing passenger requests while optimizing for wait times and optimal routing.

## Requirements

1. **Multiple Elevators & Floors**: The system handles a fleet of elevators serving multiple floors (e.g., 3 elevators serving 10 floors).
2. **Capacity Management**: Each elevator has a specific capacity limit that must not be exceeded during passenger pickup.
3. **Dynamic User Requests**: Users can request an elevator from any floor (hall call) and then define their destination.
4. **Time-Step Simulation**: The simulation advances in discrete time steps (via a `Tick()` loop) to simulate physical movement and deterministic execution.
5. **Optimized Dispatching**: The system minimizes wait times by routing requests dynamically based on the current proximity and movement direction of individual elevators. 
6. **Concurrent Requests**: The system handles simultaneous requests effectively across different floors using a SCAN-like traversal and heuristic-based assignment strategy. 
7. **Thread Safety**: Operations and request queues are designed to be safe representing proper object-oriented concurrency structures when active.

*Note: Certain strict physical mechanics (like exact door open/close animation states, dynamic weight sensors, or emergency stops) are out of scope for the logic of this core architecture.*

## Core Components

- **Direction**: Enumeration mapping the movement states (`UP`, `DOWN`, `IDLE`).
- **Request**: Describes a passenger's request parameters (source floor, destination floor, and passenger count).
- **Elevator**: Represents a single physical elevator car. It tracks `CurrentFloor`, directional state, and requested stops (`UpStops` and `DownStops`). It encapsulates the logic for computing assignment costs and stepping its own motor (`Tick()`).
- **ElevatorController**: The central dispatcher linking the fleet. Uses distance/direction heuristics to evaluate all elevators and route an incoming `Request` to the most optimal car.

## Dispatch Strategy
When a user requests an elevator, the Dispatcher (`ElevatorController`) computes an *assignment cost* for each active elevator based on:
1. **Directional Alignment**: Elevators moving toward the request in the correct direction are prioritized. An elevator moving opposite the passenger's desired direction faces an inflated heuristic cost.
2. **Proximity**: The absolute distance between the elevator's current location and the request origin.
3. **Capacity Constraints**: Any elevator that mathematically exceeds its capacity limit by servicing the request is temporarily excluded from routing. 

## Implementations
- `go-without-multithreading/`: A Golang synchronous implementation utilizing an explicit, single-threaded "game loop" to trace time steps logically.
- `java/`: A Java port mirroring the synchronous Go engine, showcasing standard OOP class designs avoiding thread-locking complexity.
