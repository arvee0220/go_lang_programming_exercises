//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vehicle interface {
	GetModel() string
	GetType() string
}

type VehicleInfo struct {
	Model string
	Type VehicleType
}

type VehicleType int

const (
	MotorcycleType VehicleType = iota
	CarType
	TruckType
)

func (v VehicleInfo) GetModel() string {
	return v.Model
}

func (v VehicleInfo) GetType() string {
	switch v.Type {
	case MotorcycleType:
		return "Motorcycle"
	case CarType:
		return "Car"
	case TruckType:
		return "Truck"
	default:
		return "Unknown"
	}
}

func DirectToLift(v Vehicle) {
	switch v.GetType() {
	case "Motorcycle":
		fmt.Printf("Directing %s to small lift\n", v.GetModel())
	case "Car":
		fmt.Printf("Directing %s to standard lift\n", v.GetModel())
	case "Truck":
		fmt.Printf("Directing %s to large lift\n", v.GetModel())
	default:
		fmt.Printf("Unknown vehicle type for %s\n", v.GetModel())
	}
}

func main() {
	vehicles := []Vehicle{
		VehicleInfo{Model: "Speedster 2000", Type: MotorcycleType},
		VehicleInfo{Model: "Family Sedan", Type: CarType},
		VehicleInfo{Model: "Road Devourer", Type: TruckType},
	}

	for _, v := range vehicles {
		DirectToLift(v)
	}

}
