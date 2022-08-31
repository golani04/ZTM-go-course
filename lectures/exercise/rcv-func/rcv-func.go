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

import "fmt"

type Name string
type Stat uint

type Player struct {
	name      Name
	health    Stat
	maxHealth Stat
	energy    Stat
	maxEnergy Stat
}

func (player *Player) addHealth(amount Stat) {
	result := player.health + amount
	if player.maxHealth < result {
		player.health = player.maxHealth
	} else {
		player.health = result
	}
	fmt.Println(player.name, "received", amount, "points to health. Health:", player.health)
}

func (player *Player) applyDamage(amount Stat) {
	result := int(player.health - amount)
	if result < 0 {
		player.health = 0
	} else {
		player.health = Stat(result)
	}
	fmt.Println(player.name, "received damaged of", amount, "points. Health:", player.health)
}

func (player *Player) addEnergy(amount Stat) {
	result := player.energy + amount
	if player.maxEnergy < result {
		player.energy = player.maxEnergy
	} else {
		player.energy = result
	}
	fmt.Println(player.name, "received", amount, "points to energy. Energy:", player.energy)
}

func (player *Player) consumeEnergy(amount Stat) {
	result := int(player.energy - amount)
	fmt.Println(result, player.energy)
	if result < 0 {
		player.energy = 0
	} else {
		player.energy = Stat(result)
	}
	fmt.Println(player.name, "consumed", amount, "points of the energy. Energy:", player.energy)
}

func main() {
	player := Player{"Steve", 100, 100, 200, 200}
	player.applyDamage(90)
	player.addHealth(11)
	player.consumeEnergy(50)
	player.addEnergy(23)

	// extremes
	player.consumeEnergy(10000)
	player.addEnergy(100000)

	player.applyDamage(101)
	player.addHealth(10001)
}
