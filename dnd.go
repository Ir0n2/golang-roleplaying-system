package main

import (
	"fmt"
	"time"
	"math/rand"
)
//enemy object, conatins health, and name values.
type enemy struct {

	name string
	health int
	level int
	weapon string
}

//here is the stuff for the player object
type player struct {

	name string
	level int
	health int
	weapon string
}
//here is the constructor for making a new player, we only need one player.
func newPlayer(name string) *player {

	p := player{name: name}
	p.level = 1
	p.health = 100
	p.weapon = "sword"

	return &p

}

//to declare a new player object you just declare a new one like so:
//playerObject := player{name: name, health: 100, level: 100, xp: 10}

//to set the parameters of the player object just call it like so:
//playerObject.health = num
//playerObject.level = num

//we only need one player so we should only need to make one of these objects.

func main () {

var dicknballs string

	playerObject := player{name: "cory", health: 30, level: 1, weapon: "Great Sword"}
        
	enemyObject := enemy{name: "Goblin", health: 30, level: 1,weapon: "mace"}



	loop: for {
		fmt.Println("dungeons and dragons demo\n0: break loop\n1: combat demo\n")
		fmt.Scanln(&dicknballs)
		switch dicknballs {

			case "0", "break":
				break loop
			case "1", "combat":
				combat(playerObject, enemyObject)

		}
	}

}

func printPlayerStats(pObject player) {
	fmt.Println("name:", pObject.name)
	fmt.Println("health:", pObject.health)
	fmt.Println("level:", pObject.level)
	fmt.Println("weapon:", pObject.weapon)
}

func printEnemyStats(pObject enemy) {
        fmt.Println("name:", pObject.name)
        fmt.Println("health:", pObject.health)
        fmt.Println("level:", pObject.level)
        fmt.Println("weapon:", pObject.weapon)
}


func combat(pObject player, eObject enemy) {
var a string
	looper: for {
		
		fmt.Println("0:break\n1:attack\n2:dodge")

		fmt.Scanln(&a)

		switch a {
			
		case "break", "0", "b":
			break looper
		case "attack", "a", "1":
		
		fmt.Println("your turn")
		//roll the dice
		dmg := diceRoll(20)
		//does damage
		eObject.health -= dmg
		//print results
		fmt.Println("rolled", dmg)	
		
		fmt.Println(eObject.name, ":", eObject.health)
			if (eObject.health <= 0) {
				fmt.Println("you win, combat over")
				break looper
			}

		//enemies turn right after you attack
		fmt.Println("enemy turn")
		//roll the dice
		enemyDmg := diceRoll(20)
		//do damage
		pObject.health -= enemyDmg
		//print results
		fmt.Println("enemie rolled", enemyDmg)
		
		fmt.Println(pObject.name, ":", pObject.health)
			if (pObject.health <= 0) {
				fmt.Println("You've fallen, darkness consumes you")
				break looper
			}
		//dodging can either result in the enemy missing their attack or you taking 10 dmg
		//I belive this 50/50 gamble will make combat more exciting
		case "dodge", "2":

			roll := diceRoll(20)
			if (roll > 10) {
				fmt.Println(eObject.name ,"misses")
			} else {
				pObject.health -= 10
				fmt.Println(eObject.name ,"hits you for 10 pts of health")
			}
		}
	}
	
	//fmt.Println("test successful")

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
