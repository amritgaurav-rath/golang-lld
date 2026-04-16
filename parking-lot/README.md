# Parking Lot System (Low Level Design)

This repository contains an implementation of a highly scalable, concurrent Parking Lot System. The system efficiently assigns parking spots across multiple levels, handling various vehicle types, dynamic availability, and thread-safe entry/exit ticketing.

## Requirements

1. **Multi-Level Layout**: The parking lot contains multiple levels, each housing a defined number of spot capacities.
2. **Vehicle Diversity**: Supports different categories of vehicles such as `Car`, `Motorcycle`, and `Truck`.
3. **Spot Compatibility**: Each parking spot guarantees support for specific vehicle types.
4. **Dynamic Assignment**: The system seamlessly assigns optimal parking spots when a vehicle enters and processes releases when the vehicle exits.
5. **Real-time Tracking**: Complete real-time awareness and polling of available spots to provide inventory checks.
6. **Concurrency Integrity**: Emphasizes avoiding race conditions. Designed properly to accommodate multiple entry and exit panels accessed concurrently.

## Core Components

- **Vehicle**: Represented dynamically based on type (`Car`, `Motorcycle`, `Truck`), inheriting basic vehicle properties like a License Plate.
- **ParkingSpot**: Tracks its assigned vehicle type, geographical structural ID, and current availability occupancy status.
- **Level**: A single floor containing multiple parking spots, providing internal indexing to find open, matching spots quickly.
- **ParkingLot (Dispatcher)**: The singleton/central controller linking multiple levels, routing incoming cars to the closest available compatible spot.

## Implementations
- `cpp/`: Native thread-safe C++ implementation modeling classic OOP hierarchies.
- `java/`: Scalable Java object-oriented build.
- `golang/` & `go-without-multithreading/`: Golang architectural codebases specifically implemented to highlight standard synchronization vs un-synchronized deterministic simulation.