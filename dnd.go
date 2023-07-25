package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main () {

	//fmt.Println(diceRoll(20))
	fmt.Println(abilityCheck(10))
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
