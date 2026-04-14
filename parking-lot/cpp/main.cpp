#include <iostream>
#include <vector>
#include <thread>
#include <chrono>
#include <random>
#include <mutex>
#include <iomanip>
#include "ParkingLot.hpp"

std::mutex printMtx;

void PrintMessage(const std::string& msg) {
    std::lock_guard<std::mutex> lock(printMtx);
    std::cout << msg << std::endl;
}

int main() {
    PrintMessage("🚀 Initializing Parking Lot Simulation...");

    // Initialize 2 levels for the parking lot
    // Level 1: 5 Motorcycle, 5 Car, 2 Truck spots
    auto level1 = std::make_shared<Level>(1, 5, 5, 2);
    // Level 2: 3 Motorcycle, 4 Car, 1 Truck spots
    auto level2 = std::make_shared<Level>(2, 3, 4, 1);

    std::vector<std::shared_ptr<Level>> levels = {level1, level2};
    ParkingLot lot(levels);

    lot.PrintAvailability();

    int numVehicles = 30;
    PrintMessage("Simulating " + std::to_string(numVehicles) + " vehicles arriving concurrently...\n");

    std::vector<std::thread> threadPool;

    for (int i = 1; i <= numVehicles; ++i) {
        threadPool.emplace_back([&lot, i]() {
            std::random_device rd;
            std::mt19937 gen(rd());
            std::uniform_int_distribution<> typeDist(0, 2);
            
            // Randomly pick a vehicle type
            VehicleType vTypes[] = {VehicleType::Motorcycle, VehicleType::Car, VehicleType::Truck};
            VehicleType vType = vTypes[typeDist(gen)];
            
            // Format ID nicely (e.g. VEH-001)
            std::string idStr = "VEH-";
            if (i < 10) idStr += "00";
            else if (i < 100) idStr += "0";
            idStr += std::to_string(i);

            Vehicle* v = new Vehicle{idStr, vType};

            // Add a slight jitter to simulate staggered parallel arrivals
            std::uniform_int_distribution<> initialDelayDist(0, 10);
            std::this_thread::sleep_for(std::chrono::milliseconds(initialDelayDist(gen)));

            ParkingSpot* spot = lot.ParkVehicle(v);
            if (spot == nullptr) {
                PrintMessage("❌ " + v->LicensePlate + " (" + VehicleTypeToString(v->Type) + ") rejected: parking lot is full");
                delete v;
                return;
            }

            PrintMessage("✅ " + v->LicensePlate + " (" + VehicleTypeToString(v->Type) + ") PARKED in spot " + spot->ID + ".");

            // Simulate time spent parked (10 to 50 milliseconds)
            std::uniform_int_distribution<> parkTimeDist(10, 50);
            std::this_thread::sleep_for(std::chrono::milliseconds(parkTimeDist(gen)));

            // Some vehicles decide to unpark (60% chance to leave)
            std::uniform_int_distribution<> leaveChanceDist(0, 9);
            if (leaveChanceDist(gen) < 6) { 
                Vehicle* unparkedV = lot.UnparkVehicle(spot->ID);
                if (unparkedV == nullptr) {
                    PrintMessage("⚠️  Error unparking " + v->LicensePlate);
                } else {
                    PrintMessage("👋 " + unparkedV->LicensePlate + " (" + VehicleTypeToString(unparkedV->Type) + ") LEFT spot " + spot->ID + ".");
                    delete unparkedV;
                }
            }
        });
    }

    // Wait for all threads to finish
    for (auto& t : threadPool) {
        if (t.joinable()) {
            t.join();
        }
    }

    PrintMessage("\n🏁 Simulation completed.");
    lot.PrintAvailability();

    return 0;
}