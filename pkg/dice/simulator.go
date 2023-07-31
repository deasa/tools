package dice

import (
	"golang.org/x/exp/slices"
	"math/rand"
	"time"
)

type Simulator struct {
	numRolls int
	die      *Die
	rand     *rand.Rand
}

func NewSimulator(numRolls int, die *Die) *Simulator {
	return &Simulator{
		numRolls: numRolls,
		die:      die,
		rand:     rand.New(rand.NewSource(int64(time.Now().Nanosecond()))),
	}
}

func (s *Simulator) SingleDie(desiredValues []int) float64 {
	wins := 0
	for i := 0; i < s.numRolls; i++ {
		v := RollValue(s.die)
		if slices.Contains(desiredValues, v) {
			wins++
		}
	}

	return float64(wins) / float64(s.numRolls)
}

func RollValue(die *Die) int {
	sides := int(die.sides)
	if sides%1 != 0 {
		panic("die must have an integer number of sides")
	}

	return rand.Intn(sides) + 1
}
