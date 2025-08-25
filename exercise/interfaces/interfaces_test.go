package main

import "testing"

func TestDirectToLift(t *testing.T) {
	vehicles := []Vehicle{
		VehicleInfo{Model: "Speedster 2000", Type: MotorcycleType},
		VehicleInfo{Model: "Family Sedan", Type: CarType},
		VehicleInfo{Model: "Road Devourer", Type: TruckType},
	}

	for _, v := range vehicles {
		DirectToLift(v)
	}
}

func TestGetType(t *testing.T) {
	v1 := VehicleInfo{Model: "Speedster 2000", Type: MotorcycleType}
	if v1.GetType() != "Motorcycle" {
		t.Errorf("Expected Motorcycle, got %s", v1.GetType())
	}

	v2 := VehicleInfo{Model: "Family Sedan", Type: CarType}
	if v2.GetType() != "Car" {
		t.Errorf("Expected Car, got %s", v2.GetType())
	}

	v3 := VehicleInfo{Model: "Road Devourer", Type: TruckType}
	if v3.GetType() != "Truck" {
		t.Errorf("Expected Truck, got %s", v3.GetType())
	}
}

func TestGetModel(t *testing.T) {
	v := VehicleInfo{Model: "Speedster 2000", Type: MotorcycleType}
	if v.GetModel() != "Speedster 2000" {
		t.Errorf("Expected Speedster 2000, got %s", v.GetModel())
	}
}