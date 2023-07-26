package main

import (
	"fmt"
	"time"
	"math/rand"
)
//here is the stuff for the player object
type player struct {

	name string
	level int
	health int
	xp int
}
//here is the constructor for making a new player, we only need one player.
func newPlayer(name string) *player {

	p := player{name: name}
	p.level = 1
	p.health = 100
	p.xp = 0

	return &p
}
//to declare a new player object you just declare a new one like so:
//playerObject := player{name: name, health: 100, level: 100, xp: 10}

//to set the parameters of the player object just call it like so:
//playerObject.health = num
//playerObject.level = num

//we only need one player so we should only need to make one of these objects.

func main () {

	fmt.Println(newPlayer("Rory"))
	playerObject := player{name: "cory", health: 69, level: 1, xp: 100}

	fmt.Println(playerObject)
	playerObject.health = 42
	fmt.Println(playerObject)
	//fmt.Println(diceRoll(20))
	//fmt.Println(abilityCheck(10))
}

//abiltycheck will roll the dice and apply the modifier score accordingly.
//to use the function you put your ability score as the modifier int.
func abilityCheck(modifier int) int {
	roll := diceRoll(20)
	
	switch modifier {
	
	case 1:
		return roll - 5
        case 2, 3:
                return roll - 4
        case 4, 5:
                return roll - 3
        case 6, 7:
                return roll - 2	
	case 8, 9:
                return roll - 1
	case 10, 11:
                return roll
	case 12, 13:
                return roll + 1
	case 14, 15:
                return roll + 2
        case 16, 17:
                return roll + 3
        case 18, 19:
                return roll + 4
        case 20, 21:
                return roll + 5
        case 22, 23:
                return roll + 6
        case 24, 25:
                return roll + 7
        case 26, 27:
                return roll + 8
	case 28, 29:
		return roll + 9
	case 30:
		return roll + 10
	}
	return roll
}

//dice roll will roll a dice of any sides. or in other words, generate a number between 1 and the number of sides.
func diceRoll(sides int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
	a := r1.Intn(sides)
	if (a == 0) {
	a++
	}
	return a
}
