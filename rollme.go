package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func usage(errormessage string) {
	println("============================================================================================")
	println("Error: ", errormessage)
	println("============================================================================================")
	println("rollme")
	println("(c) 2019 ilPinguino")
	println("GPLv3")
	println()
	println("USAGE")
	println("============================================================================================")
	println("rollme <Dice> [dice]")
	println()
	println("Dice syntax: 1d6")
	println("1 = Count of dice")
	println("6 = Count of sides (6 is a normal dice, 20 is a 20-sided dice, 2 is a coin)")
	println("The number of dice expressions per invocation is unlimited. Roll as much as you like.")
	println("ISSUES")
	println("============================================================================================")
	println("This project is perfect and does not have issues, because it is written by god himself. \nIf you want to leave offerings, find me on github.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rolls := os.Args[1:]

	if len(rolls) == 0 {
		usage("No rolls found.")
	}

	// Each roll consists of the DnD writing style - for example, 1d6 would be 1 dice with 6 sides.
	validroll := false
	for _, roll := range rolls {

		fmt.Println("Roll ", roll, ": ")
		rolldata := strings.Split(roll, "d")
		if len(rolldata) != 2 {
			fmt.Println("Syntax Error in roll", roll)
			continue
		}
		count, err := strconv.Atoi(rolldata[0])
		if err != nil {
			fmt.Println("Syntax Error in roll", roll)
			continue
		}
		dice, err := strconv.Atoi(rolldata[1])
		if err != nil {
			fmt.Println("Syntax Error in roll", roll)
			continue
		}

		for i := 0; i < count; i++ {
			fmt.Println(i, ":", rand.Intn(dice)+1) // Intn poduces 0 < n < max, we want to eliminate 1
			validroll = true
		}
	}

	if !validroll {
		usage("No valid rolls found.")
	}

}
