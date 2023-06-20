package format

import "fmt"

func FloatToPercentage(f float64) string {
	return fmt.Sprintf("%.8f%%", f*100)
}
