package worker

import (
	"fmt"
	"parking_lot/parkinglot"
)

type WorkerC struct {
	parkingLot *parkinglot.ParkingLot
}

func NewWorkerC(parkingLot *parkinglot.ParkingLot) *WorkerC {
	return &WorkerC{
		parkingLot: parkingLot,
	}
}

func (w *WorkerC) PerformTask() {
	w.parkingLot.Mu.Lock()
	defer w.parkingLot.Mu.Unlock()

	fmt.Printf("Worker C: Performing task\n")
	// Perform worker C specific task (example: access certain information)
}
