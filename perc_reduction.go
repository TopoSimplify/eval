package eval

import (
	"github.com/intdxdt/math"
)

//removed vertices ratio
func rmRatio(rm, at int) float64 {
	return float64(rm) / float64(at)
}

//percentage redundancy
func percRedundancy(rm, at []int) float64 {
	return math.Round(rmRatio(len(rm), len(at))*100.0, 0)
}

//percentage interesting
func perc_intrest(rm, at []int) float64 {
	return 100.0 - percRedundancy(rm, at)
}
