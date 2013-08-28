package libfun

import (
	"fmt"
)

type Make int

var makeStrings []string = []string{"Honda", "Nissan", "Tesla"}

// This is better:
const (
	Honda = iota
	Nissan
	Tesla
)

func (m Make) String() string {
	return makeStrings[m]
}

type Car struct {
	id       int
	name     string
	makeType Make
}

type Truck struct {
	id     int
	name   string
	weight int
}

type Vehicle interface {
	PeddleToTheMetal()
	DisplayName() string
}

// -- <implementation of="Vehicle">
func (car *Car) PeddleToTheMetal() {
	fmt.Printf("Driving my %s as fast as I can!!!\n", car.makeType)
}

func (truck *Truck) PeddleToTheMetal() {
	fmt.Printf("Truck %d is so slow\n", truck.id)
}

func (car *Car) DisplayName() string {
	return fmt.Sprintf("CarId: %d", car.id)
}

func (truck *Truck) DisplayName() string {
	return fmt.Sprintf("TruckId: %d, Weight: %d", truck.id, truck.weight)
}

// -- </implementation>

func NewCar() *Car {
	car := &Car{1, "Mark's Dream Car", Tesla}
	return car
}

func NewTruck() *Truck {
	truck := &Truck{2, "Big Ass Truck", 1000}
	return truck
}

func PunchIt(v Vehicle) {
	fmt.Printf("I've got vehicle ID %s... time to punch it!\n", v.DisplayName())
	v.PeddleToTheMetal()
}
