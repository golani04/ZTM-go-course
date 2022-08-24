//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerInfo(servers map[string]int) {
	fmt.Println("--- Server Statuses --")
	fmt.Println("Number of servers", len(servers))

	stats := make(map[int]int)
	for _, status := range servers {
		if status > Retired {
			panic("unhandled server status")
		}
		stats[status] += 1
	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")
	fmt.Println()
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatuses := make(map[string]int)
	for _, server := range servers {
		serverStatuses[server] = Online
	}

	printServerInfo(serverStatuses)

	serverStatuses["darkstar"] = Retired
	serverStatuses["aiur"] = Offline

	printServerInfo(serverStatuses)

	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}

	printServerInfo(serverStatuses)

}
