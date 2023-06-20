package dice

type DiceGame struct {
	NumDice int
	Die     *Die
}

func NewDiceGame(numDice int, die *Die) *DiceGame {
	return &DiceGame{NumDice: numDice, Die: die}
}

// FixedValuesAcrossRolls will compute the probability that each roll in `numRolls` will have the specified `numValues` of dice locked to a particular value
// e.g. probability of rolling two 3s five times in a row would be FixedValuesAcrossRolls(2, 5)
// e.g. probability of rolling three Yahtzees in a row would be FixedValuesAcrossRolls(5, 3)
func FixedValuesAcrossRolls(numValues, numRolls int) float64 {
	return float64(0)
}
