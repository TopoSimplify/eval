package eval

import (
	"github.com/intdxdt/math"
)

//removed vertices ratio
func rm_ratio(rm, at int) float64 {
	return float64(rm) / float64(at)
}

//percentage redundancy
func perc_redundancy(rm, at []int) float64 {
	return math.Round(rm_ratio(len(rm), len(at))*100.0, 0)
}

//percentage interesting
func perc_intrest(rm, at []int) float64 {
	return 100.0 - perc_redundancy(rm, at)
}
