package dice

import "math"

type Die struct {
	sides float64
}

func NewDie(sides int) *Die {
	if sides < 1 {
		panic("die must have at least one side")
	}
	return &Die{sides: float64(sides)}
}

// SingleRoll will compute the probability of any value being rolled on a die with a given number of sides
func (d *Die) SingleRoll() float64 {
	return float64(1) / d.sides
}

// SingleValueAcrossRolls will compute the probability of `numRolls` number of rolls resolving to a desired value each time
// e.g. probability of rolling six 6s in a row would be SingleValueAcrossRolls(6) called from a die with die.sides = 6
// e.g. probability of rolling 1, 2, 3, 4, 5, 6 in that order would also be SingleValueAcrossRolls(6)
func (d *Die) SingleValueAcrossRolls(numRolls int) float64 {
	var prob float64
	for i := 0; i < numRolls; i++ {
		if i == 0 {
			prob += d.SingleRoll()
			continue
		}
		prob *= d.SingleRoll()
	}
	return prob
}

// XOfAKindSingleRoll will compute the probability of rolling x number of a given value in a single roll
// e.g. probability of rolling two 6s in a single roll out of 5 dice would be XOfAKindSingleRoll(2, 5)
// e.g. probability of rolling a Yahtzee in a single roll would be XOfAKindSingleRoll(5, 5)
func (d *Die) XOfAKindSingleRoll(numDesired, numDice int) float64 {
	if numDice == 0 {
		return 0
	}

	possibilities := math.Pow(d.sides, float64(numDice))

	return float64(numDesired) / possibilities
}
