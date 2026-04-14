# 🚗 Parking Lot System (LLD)

## 📌 Problem Statement
Design a Parking Lot System that efficiently manages vehicle parking across multiple levels with support for different vehicle types and concurrent operations. This problem statement is solved in three different languages, i.e., C++, Java & Golang.

---

## 🧾 Requirements

### Functional Requirements
- The parking lot should have multiple levels, each with a fixed number of parking spots.
- The system should support different types of vehicles:
  - Car
  - Motorcycle
  - Truck
- Each parking spot should support a specific vehicle type.
- The system should:
  - Assign a parking spot on vehicle entry
  - Release the spot on exit
- Track availability of parking spots in real-time.
- Provide availability information to customers.

---

### Non-Functional Requirements
- The system should support:
  - Multiple entry and exit points
  - Concurrent access (thread-safe operations)
- Should be scalable for large parking lots.