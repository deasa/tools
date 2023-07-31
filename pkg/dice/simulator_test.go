package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRollValue(t *testing.T) {
	tests := []struct {
		name        string
		die         *Die
		wantValueIn []int
	}{
		{
			name:        "d6",
			die:         NewDie(6),
			wantValueIn: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:        "d1",
			die:         NewDie(1),
			wantValueIn: []int{1},
		},
		{
			name:        "d20",
			die:         NewDie(20),
			wantValueIn: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RollValue(tt.die)

			assert.Contains(t, tt.wantValueIn, got, "got unexpected value from dice roll")
		})
	}

	t.Run("d0", func(t *testing.T) {
		assert.Panics(t, func() {
			NewDie(0)
		})
	})
}

func TestSimulator_SingleDie(t *testing.T) {
	tests := []struct {
		name      string
		die       *Die
		numRolls int
		desiredValues []int
		wantValue float64
	}{
		{
			name:      "d6",
			die:       NewDie(6),
			numRolls: 1000,
			desiredValues:
			wantValue: .333,
		},
		{
			name:      "d1",
			die:       NewDie(1),
			wantValue: 1,
		},
		{
			name:      "d20",
			die:       NewDie(20),
			wantValue: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSimulator(tt.numRolls, tt.die)
			got := s.SingleDie()

			assert.Equal(t, tt.wantValue, got, "got unexpected value from dice roll")
		})
	}
}
