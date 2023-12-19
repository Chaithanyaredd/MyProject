package supervisor

import (
	"fmt"
	"parking_lot/parkinglot"
)

type Supervisor struct {
	parkingLot *parkinglot.ParkingLot
}

func NewSupervisor(parkingLot *parkinglot.ParkingLot) *Supervisor {
	return &Supervisor{
		parkingLot: parkingLot,
	}
}

func (s *Supervisor) GetStatistics() {
	s.parkingLot.Mu.Lock()
	defer s.parkingLot.Mu.Unlock()
	fmt.Printf("Total parked cars: %d\n", len(s.parkingLot.ParkedVehicle))
}
