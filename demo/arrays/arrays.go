package main

import (
	"fmt"
	"time"
)
type Room struct {
	name string
	cleaned bool
}

func checkCleanliness(rooms []Room) {
	fmt.Println("Checking cleanliness of rooms...")
	time.Sleep(2 * time.Second)
	for i := range rooms {
		if rooms[i].cleaned {
			time.Sleep(2 * time.Second)
			fmt.Println(rooms[i].name, "is clean.")
		} else {
			time.Sleep(2 * time.Second)
			fmt.Println(rooms[i].name, "is dirty")
		}
	}
	time.Sleep(2 * time.Second)
}

func cleanRoom(rooms []Room) {
	for i := range rooms {
		if !rooms[i].cleaned {
			time.Sleep(2 * time.Second)
			fmt.Println("Cleaning", rooms[i].name)
			time.Sleep(2 * time.Second)
			rooms[i].cleaned = true
		} else {
			fmt.Println(rooms[i].name, "is already clean.")
		}
	}

	fmt.Println("All rooms have been cleaned.")
	time.Sleep(2 * time.Second)
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms[:])

	fmt.Println("Performing cleaning...")

	cleanRoom(rooms[:])

	checkCleanliness(rooms[:])

}
