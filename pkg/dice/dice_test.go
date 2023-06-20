package dice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSingleRoll() {
	d := NewDie(6)
	fmt.Println(d.SingleRoll())
	// Output: 0.16666666666666666
}

func TestSingleRoll(t *testing.T) {
	testCases := []struct {
		sides           int
		wantProbability float64
	}{
		{
			sides:           6,
			wantProbability: 0.1667,
		},
		{
			sides:           12,
			wantProbability: 0.0833,
		},
		{
			sides:           1200,
			wantProbability: 0.000833,
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprint(tt.sides), func(t *testing.T) {
			d := NewDie(tt.sides)
			got := d.SingleRoll()
			assert.InDelta(t, got, tt.wantProbability, 0.0001, "incorrect probability")
		})
	}
}

func TestSingleValueAcrossRolls(t *testing.T) {
	testCases := []struct {
		numSides        int
		numRolls        int
		wantProbability float64
	}{
		{
			numSides:        6,
			numRolls:        2,
			wantProbability: 0.02778,
		},
		{
			numSides:        20,
			numRolls:        2,
			wantProbability: 0.0025,
		},
		{
			numSides:        6,
			numRolls:        30,
			wantProbability: 4.52e-24,
		},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprint(tt.numRolls), func(t *testing.T) {
			d := NewDie(tt.numSides)
			got := d.SingleValueAcrossRolls(tt.numRolls)

			assert.InDelta(t, got, tt.wantProbability, 0.0001, "incorrect probability")
		})
	}
}

func TestDie_XOfAKindSingleRoll(t *testing.T) {
	type fields struct {
		sides float64
	}
	type args struct {
		numDesired int
		numDice    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "yahtzee",
			fields: fields{
				sides: 6,
			},
			args: args{
				numDesired: 5,
				numDice:    5,
			},
			want: 0.0007778,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Die{
				sides: tt.fields.sides,
			}
			assert.Equal(t, tt.want, d.XOfAKindSingleRoll(tt.args.numDesired, tt.args.numDice), "XOfAKindSingleRoll(%v, %v)", tt.args.numDesired, tt.args.numDice)
		})
	}
}
