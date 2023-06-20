package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test for floatToPercentage
func TestFloatToPercentage(t *testing.T) {
	testCases := []struct {
		name string
		in   float64
		want string
	}{
		{
			name: "zero",
			in:   0,
			want: "0.00%",
		},
		{
			name: "one",
			in:   1,
			want: "100.00%",
		},
		{
			name: "half",
			in:   0.5,
			want: "50.00%",
		},
		{
			name: "quarter",
			in:   0.25,
			want: "25.00%",
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := FloatToPercentage(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
