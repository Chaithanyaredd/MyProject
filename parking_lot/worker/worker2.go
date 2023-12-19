package worker

import (
	"fmt"
	"parking_lot/parkinglot"
)

type WorkerB struct {
	parkingLot *parkinglot.ParkingLot
}

func NewWorkerB(parkingLot *parkinglot.ParkingLot) *WorkerB {
	return &WorkerB{
		parkingLot: parkingLot,
	}
}

func (w *WorkerB) PerformTask() {
	w.parkingLot.Mu.Lock()
	defer w.parkingLot.Mu.Unlock()

	fmt.Printf("Worker B: Performing task\n")
	// Perform worker B specific task (example: access certain information)
}
