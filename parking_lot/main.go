package main

import (
	"fmt"
	"os"
	"parking_lot/parkinglot"
	"parking_lot/supervisor"
)
func main() {
	parkingLot := parkinglot.NewParkingLot()
	supervisor := supervisor.NewSupervisor(parkingLot)
	// workerA := worker.NewWorkerA(parkingLot)
	// workerB := worker.NewWorkerB(parkingLot)
	// workerC := worker.NewWorkerC(parkingLot)

	for {
		var action, vehicleName, vehicleSize string
		fmt.Print("Enter action ('park', 'remove', 'statistics', or 'exit'): ")
		fmt.Scan(&action)

		switch action {
		case "park":
			fmt.Print("Enter vehicle name: ")
			fmt.Scan(&vehicleName)
			fmt.Print("Enter vehicle size ('big', 'medium', 'small', 'bike'): ")
			fmt.Scan(&vehicleSize)
			parkingLot.ParkVehicle(vehicleName, vehicleSize)
			parkingLot.PrintRemainingSlots()
		case "remove":
			fmt.Print("Enter vehicle name to remove: ")
			fmt.Scan(&vehicleName)
			fmt.Print("Enter vehicle size ('big', 'medium', 'small', 'bike'): ")
			fmt.Scan(&vehicleSize)
			parkingLot.RemoveVehicle(vehicleName, vehicleSize)
			parkingLot.PrintRemainingSlots()
		case "statistics":
			supervisor.GetStatistics()
		case "exit":
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid action. Please enter 'park', 'remove', 'statistics', or 'exit'.")
		}
	}
}