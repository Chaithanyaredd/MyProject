package parkinglot

import ("sync"
		"fmt"
)

const (
	TotalSlots = 100
	Big        = "big"
	Medium     = "medium"
	Small      = "small"
	Bike       = "bike"
)

type ParkingLot struct {
	bigSlots    int
	mediumSlots int
	smallSlots  int
	bikeSlots   int
	ParkedVehicle  map[string]string // Map vehicle names to their respective sizes
	Mu          sync.Mutex        // Mutex to synchronize access to shared data
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		bigSlots:    TotalSlots / 4,
		mediumSlots: TotalSlots / 4,
		smallSlots:  TotalSlots / 4,
		bikeSlots:   TotalSlots / 4,
		ParkedVehicle:  make(map[string]string),
	}
}

func (p *ParkingLot) ParkVehicle(name, size string) {
	switch size {
	case Big, Medium, Small:
		p.Park(name, size)
	case Bike:
		p.ParkBike(name)
	default:
		fmt.Println("Invalid vehicle size. Please enter 'big', 'medium', 'small', or 'bike'.")
	}
}

func (p *ParkingLot) RemoveVehicle(name, size string) {
	if size == Bike {
		p.RemoveBike(name)
	} else {
		p.Remove(name, size)
	}
}

func (p *ParkingLot) Park(name, size string) {
	slotCount := p.GetSlotCount(size)
	if slotCount > 0 {
		p.ParkedVehicle[name] = size
		p.DecrementSlotCount(size)
		fmt.Printf("%s '%s' parked in %s slot\n", size, name, size)
	} else {
		fmt.Printf("No available slots for %s '%s'\n", size, name)
	}
}

func (p *ParkingLot) ParkBike(name string) {
	if p.bikeSlots > 0 {
		p.ParkedVehicle[name] = Bike
		p.bikeSlots--
		fmt.Printf("Bike '%s' parked\n", name)
	} else {
		fmt.Printf("No available slots for bike '%s'\n", name)
	}
}

func (p *ParkingLot) Remove(name, size string) {
	if size == Bike {
		p.RemoveBike(name)
	} else {
		if size, exists := p.ParkedVehicle[name]; exists {
			delete(p.ParkedVehicle, name)
			p.IncrementSlotCount(size)
			fmt.Printf("%s '%s' removed from %s slot\n", size, name, size)
		} else {
			fmt.Printf("Vehicle '%s' is not parked in the parking lot\n", name)
		}
	}
}

func (p *ParkingLot) RemoveBike(name string) {
	if size, exists := p.ParkedVehicle[name]; exists && size == Bike {
		delete(p.ParkedVehicle, name)
		p.bikeSlots++
		fmt.Printf("Bike '%s' removed\n", name)
	} else {
		fmt.Printf("Bike '%s' is not parked in the parking lot\n", name)
	}
}

func (p *ParkingLot) GetSlotCount(size string) int {
	switch size {
	case Big:
		return p.bigSlots
	case Medium:
		return p.mediumSlots
	case Small:	
		return p.smallSlots
	case Bike:
		return p.bikeSlots
	default:
		return 0
	}
}

func (p *ParkingLot) IncrementSlotCount(size string) {
	switch size {
	case Big:
		p.bigSlots++
	case Medium:
		p.mediumSlots++
	case Small:
		p.smallSlots++
	case Bike:
		p.bikeSlots++
	}
}

func (p *ParkingLot) DecrementSlotCount(size string) {
	switch size {
	case Big:
		p.bigSlots--
	case Medium:
		p.mediumSlots--
	case Small:
		p.smallSlots--
	case Bike:
		p.bikeSlots--
	}
}

func (p *ParkingLot) PrintRemainingSlots() {
	fmt.Printf("Remaining big slots: %d\n", p.bigSlots)
	fmt.Printf("Remaining medium slots: %d\n", p.mediumSlots)
	fmt.Printf("Remaining small slots: %d\n", p.smallSlots)
	fmt.Printf("Remaining bike slots: %d\n", p.bikeSlots)
}