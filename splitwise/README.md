# Splitwise System (Low Level Design)

This repository contains deeply scalable implementations of a Splitwise-style cost-sharing architecture. The system efficiently supports global users, friend groups, detailed expense ledgers, sophisticated balancing mathematics, and concurrent settlement execution safely.

## Requirements

1. **User Base & Identities**: Platform securely retains independent user accounts mapped globally.
2. **Groups**: Supports collaborative ledgers natively tracking grouped users and isolated chronological histories of joint expenses.
3. **Expenses & Receipt Context**: Expenses define total amounts, descriptive reasons, and directly map abstract Split interfaces specifying granular distributions.
4. **Adaptive Algorithms**: Features explicit math interfaces (`EqualSplit`, `ExactSplit`, `PercentSplit`) validating correct distributions intrinsically.
5. **Ledger Mathematics**: Instead of logging raw totals, cleanly processes a topological mapping natively representing peer-to-peer balance matrices (*"Who owes who exactly what"*). 
6. **Thread Safety & Consistency**: Under parallel load requests, the system guarantees accurate multi-transaction sums without triggering double spending or internal state corruption.

*Note: Distributed ID databases or physical networking implementations fall natively outside the specific abstract classes.*

## Core Components

- **User**: Base identity model storing essential constraints like emails/names.
- **Group**: Structural aggregator for multi-member instances natively indexing an `Expense` log.
- **Expense**: Contains raw totals, paying references, description, and slice arrays mapping raw `Split` distributions across involved participants.
- **Split**: Interface resolving precise distribution mathematics natively into flat float balances per member (`EqualSplit`, `PercentSplit`, etc).
- **Transaction**: Settlement ledger proving zero outs between mutually indebted peers.
- **SplitwiseService (Singleton)**: Controller logic computing total expense parsing logic prior to adjusting the global debt matrices dynamically. 

## Implementations
- `go-without-multithreading/`: The synchronous pure logic model verifying accurate OOP data cascades strictly unhampered by thread locks.
- `go/`: Thread-safe, enterprise Golang implementation wrapping global mutators (`sync.RWMutex`) to gracefully navigate and absorb parallel transaction bursts simultaneously.

## Setup & Running
Navigate into either internal folder and fire:
```bash
go run .
```
