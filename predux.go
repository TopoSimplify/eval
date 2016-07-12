package eval

import (
  . "simplex/util/math"
)

//removed vertices ratio
func rmratio(rm , at int) float64{
    return float64(rm) / float64(at)
}

//percentage redundancy
func predundancy(rm , at []int) float64{
  return Round(rmratio(len(rm), len(at)) * 100.0, 0)
}

//percentage interesting
func pintrest(rm , at []int) float64{
  return 100.0 - predundancy(rm , at )
}
