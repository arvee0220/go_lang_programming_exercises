//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestModifyHealth(t *testing.T) {
	player := Player{Health: 50, MaxHealth: 100, Energy: 30, MaxEnergy: 100, Name: "Hero"}

	tests := []struct {
		h, mh       int
		wantHealth  int
		wantMaxHealth int
	}{
		{20, 0, 70, 100},   // Increase health
		{-30, 0, 40, 100},  // Decrease health
		{0, 20, 40, 120},   // Increase max health
		{0, -10, 40, 110},  // Decrease max health
		{100, 0, 110, 110}, // Health should not exceed max health
		{-200, 0, 0, 110},  // Health should not go below 0
	}

	for _, tt := range tests {
		player.ModifyHealth(tt.h, tt.mh)
		if player.Health != tt.wantHealth || player.MaxHealth != tt.wantMaxHealth {
			t.Errorf("ModifyHealth(%d, %d) = (Health: %d, MaxHealth: %d); want (Health: %d, MaxHealth: %d)",
				tt.h, tt.mh, player.Health, player.MaxHealth, tt.wantHealth, tt.wantMaxHealth)
		}
	}
}

func TestModifyEnergy(t *testing.T) {
	player := Player{Health: 50, MaxHealth: 100, Energy: 30, MaxEnergy: 100, Name: "Hero"}

	tests := []struct {
		e, me       int
		wantEnergy  int
		wantMaxEnergy int
	}{
		{20, 0, 50, 100},   // Increase energy
		{-30, 0, 20, 100},  // Decrease energy
		{0, 20, 20, 120},   // Increase max energy
		{0, -10, 20, 110},  // Decrease max energy
		{100, 0, 110, 110}, // Energy should not exceed max energy
		{-200, 0, 0, 110},  // Energy should not go below 0
	}

	for _, tt := range tests {
		player.ModifyEnergy(tt.e, tt.me)
		if player.Energy != tt.wantEnergy || player.MaxEnergy != tt.wantMaxEnergy {
			t.Errorf("ModifyEnergy(%d, %d) = (Energy: %d, MaxEnergy: %d); want (Energy: %d, MaxEnergy: %d)",
				tt.e, tt.me, player.Energy, player.MaxEnergy, tt.wantEnergy, tt.wantMaxEnergy)
		}
	}
}