//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import (
	"fmt"
	"time"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServerInfo(serverStatuses map[string]int) {
	totalServers := len(serverStatuses)
	online := 0
	offline := 0
	maintenance := 0
	retired := 0

	for _, status := range serverStatuses {
		switch status {
		case Online:
			online += 1
		case Offline:
			offline += 1
		case Maintenance:
			maintenance += 1
		case Retired:
			retired += 1
		}
	}
	fmt.Println("Total servers:", totalServers)
	fmt.Println("Online servers:", online)
	fmt.Println("Offline servers:", offline)
	fmt.Println("Maintenance servers:", maintenance)
	fmt.Println("Retired servers:", retired)
	fmt.Println()
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatuses := make(map[string]int)	

	for _, server := range servers {
		serverStatuses[server] = Online
		fmt.Println("Server", server, "is set to Online")
		time.Sleep(100 * time.Millisecond)
	}

	displayServerInfo(serverStatuses)
	serverStatuses["darkstar"] = Retired
	serverStatuses["aiur"] = Offline
	displayServerInfo(serverStatuses)

	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
		fmt.Println("Server", server, "is set to Maintenance")
		time.Sleep(100 * time.Millisecond)
	}

	displayServerInfo(serverStatuses)
}
