package worker

import (
	"fmt"
	"parking_lot/parkinglot"
)

type WorkerA struct {
	parkingLot *parkinglot.ParkingLot
}

func NewWorkerA(parkingLot *parkinglot.ParkingLot) *WorkerA {
	return &WorkerA{
		parkingLot: parkingLot,
	}
}

func (w *WorkerA) PerformTask() {
	w.parkingLot.Mu.Lock()
	defer w.parkingLot.Mu.Unlock()

	fmt.Printf("Worker A: Performing task\n")
	// Perform worker A specific task (example: access certain information)
}
