package main

import (
	"fmt"
	"github.com/deasa/tools/pkg/dice"
	"github.com/deasa/tools/pkg/format"
)

func main() {
	die := dice.NewDie(6)

	fmt.Printf("probability of a single number in a single roll: %s\n", format.FloatToPercentage(die.SingleRoll()))
	fmt.Printf("probability of a single number across 2 rolls: %s\n", format.FloatToPercentage(die.SingleValueAcrossRolls(2)))
	fmt.Printf("probability of a single number across 6 rolls: %s\n", format.FloatToPercentage(die.SingleValueAcrossRolls(6)))
}
