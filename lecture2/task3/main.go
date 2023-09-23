package main

import (
	"fmt"
)

type Vehicle interface {
	//	Insurance()
	BuildARoute()
	Start()
	Stop()
}

type Car struct {
	Name string
}

// func (c Car) Insurance() {
// 	var startDay string
// 	var endDay string
// 	fmt.Println("Enter the location from where you want to start your trip")
// 	fmt.Scan(&yourLocation)
// 	fmt.Println("Enter your destination")
// 	fmt.Scan(&destination)
// 	fmt.Printf("The route has been built for %s.\n", c.Name)
// }

func (c Car) BuildARoute() {
	var yourLocation string
	var destination string
	fmt.Println("Enter the location from where you want to start your trip")
	fmt.Scan(&yourLocation)
	fmt.Println("Enter your destination")
	fmt.Scan(&destination)
	fmt.Printf("The route has been built for %s.\n", c.Name)
}

func (c Car) Start() {
	fmt.Printf("%s starts the engine.\n", c.Name)
}

func (c Car) Stop() {
	fmt.Printf("%s stops the engine.\n", c.Name)
}

type Bike struct {
	Name string
}

func (b Bike) BuildARoute() {
	var yourLocation string
	var destination string
	fmt.Println("Enter the location from where you want to start your trip")
	fmt.Scan(&yourLocation)
	fmt.Println("Enter your destination")
	fmt.Scan(&destination)
	fmt.Printf("The route has been built for %s.\n", b.Name)
}

func (b Bike) Start() {
	fmt.Printf("%s starts pedaling.\n", b.Name)
}

func (b Bike) Stop() {
	fmt.Printf("%s stops pedaling.\n", b.Name)
}

type Garage struct {
	Vehicles []Vehicle
}

func (g Garage) OperateAll() {
	fmt.Println("Operating all vehicles in the garage:")
	for _, v := range g.Vehicles {
		fmt.Printf(" - %T\n", v)
		v.BuildARoute()
		v.Start()
		v.Stop()
		fmt.Println()
	}
}

func main() {
	car1 := Car{Name: "Car1"}
	car2 := Car{Name: "Car2"}
	bike1 := Bike{Name: "Bike1"}
	bike2 := Bike{Name: "Bike2"}

	garage := Garage{Vehicles: []Vehicle{car1, car2, bike1, bike2}}

	garage.OperateAll()
}
