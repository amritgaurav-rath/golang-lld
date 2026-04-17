# Golang Single-Threaded Ride-Sharing SDE-II Architecture 🚖

This repository cleanly deploys a rigidly decoupled S.O.L.I.D. ride-sharing prototype answering standard FAANG architecture matrices.

## Architectural Deep Dive

### 1. `enums/`
Contains completely decoupled constants securely preventing raw strings natively.
- **`vehicle_type.go`**: Mathematically traces exactly `Activa`, `Polo`, and `XUV` types reliably!
- **`ride_status.go`**: Tracks mathematical lifecycle mappings avoiding boolean drift natively (`Offered`, `Completed`).

### 2. `entities/`
Contains 100% logic-free structural instances ensuring purely strict memory encapsulation natively.
- **`user.go`**
- **`vehicle.go`**
- **`ride.go`**

### 3. `strategy/`
Isolates algorithmic searches from application endpoints exactly matching the Java SDE-2 equivalent!
- **`preferred_vehicle_strategy.go`**: Automatically traces parameters extracting instances natively exactly matching constraints natively.
- **`most_vacant_strategy.go`**: Mathematically compares seating capacity tracking strictly finding largest active buffers seamlessly!

### 4. `services/ride_sharing_system.go`
Single-threaded facade safely executing explicit rules without locks or mutexes strictly matching the prompt's deterministic constraint bounds natively. Includes embedded rule validation logic correctly avoiding simultaneous driver-active arrays!
