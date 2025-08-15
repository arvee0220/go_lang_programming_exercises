//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
	"reflect"
)

type Player struct {
	Health int
	MaxHealth int
	Energy int
	MaxEnergy int
	Name string
}

func printChange(statName string, prevVal, newVal int) {
	switch {
	case newVal == prevVal:
		fmt.Printf("No changes made to %s\nplayer %s: %v\n", statName, statName, newVal)
	case newVal < prevVal:
		fmt.Printf("Player %s reduced by %v\n", statName, prevVal - newVal)
	case newVal > prevVal:
		fmt.Printf("Player %s increased by %v\n", statName, newVal - prevVal)
	}

}

func (player *Player) modifyHealth(h, mh int) {
	prevHealth := player.Health
	prevMaxHealth := player.MaxHealth

	player.Health += h
	player.MaxHealth += mh

	printChange("health", prevHealth, player.Health)
	printChange("max health", prevMaxHealth, player.MaxHealth)
}

func (player *Player) modifyEnergy(e, me int) {
	prevEnergy := player.Energy
	prevMaxEnergy := player.MaxEnergy

	player.Energy += e
	player.MaxEnergy += me

	printChange("energy", prevEnergy, player.Energy)
	printChange("max energy", prevMaxEnergy, player.MaxEnergy)
}

func structToMap(p Player) map[string]int {
	result := make(map[string]int)
	val := reflect.ValueOf(p)
	prop := reflect.TypeOf(p)

	for i := 0; i < val.NumField(); i++ {
		fieldName := prop.Field(i).Name
		fieldValue := val.Field(i)

		if fieldValue.Kind() == reflect.Int {
			result[fieldName] = int(fieldValue.Int())
		}
	}

	return result
}

func main() {
	player := Player{90, 100, 50, 100, "Max"}

	player.modifyHealth(5, 10)
	player.modifyEnergy(-5, 10)

	playerStat := structToMap(player)

	fmt.Printf("%v's current stat:\n", player.Name)
	for i, v := range playerStat {	
		fmt.Println(i, v)
	}
}
