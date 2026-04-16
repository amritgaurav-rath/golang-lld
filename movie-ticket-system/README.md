# Movie Ticket Booking System (Low Level Design)

This repository focuses on demonstrating a scalable architecture for a concurrent Movie Ticket System. It is responsible for globally matching users, theaters, movies, available show schedules, and explicit seat reservations.

## Requirements

1. **Browse Shows**: Users can query currently playing movies, viewing a collection of participating theaters and showtimes.
2. **Seat Selection**: Users can selectively pick and lock individual seats out of an explicit spatial seating arrangement within a specific show map.
3. **Seat Categories**: Seats come in structured tiers (`NORMAL`, `PREMIUM`) dictating unique pricing configurations.
4. **Ticket Booking**: Process of securely reserving and purchasing seats with strict lifecycle tracking (`PENDING`, `CONFIRMED`, `CANCELLED`).
5. **Concurrency Safety**: Ensures perfectly thread-safe transactions, actively avoiding double-booking the exact same seat within the exact same `Show` instance.

*Note: Integrated external payment processing APIs and authentication token layers are treated as abstract dependencies outside of the core Low Level Design scope.*

## Core Components

- **User**: The client entity interacting with the platform.
- **Movie**: Catalog entity containing playback duration and general metadata.
- **Theater**: Physical projector locations bridging local identity to multiple daily `Show` schedules.
- **Show**: Represents a specific runtime instance for a Movie at a Theater, encapsulating its own `Seat` availability index map. Contains a dedicated internal mutex-lock to process atomic multi-seat transactions.
- **Seat**: Details precise physical placements (`Row`, `Column`) along with dynamic booking availability states.
- **Booking**: The confirmation envelope coupling users to their locked seats, enabling end-to-end receipt validation tracking.

## Implementations
- `Go/`: The complete thread-safe Golang implementation encapsulating concurrent mutex-locking for strict booking collision prevention.
- `go-without-multithreading/`: A synchronous, stripped-down port evaluating exact business logic routines perfectly independent of mutex locking overhead.
