package eval

import (
  . "simplex/vect"
)

//deflections
func dflns(vsegs []*Vect) []float64{
  var n = len(vsegs) - 1
  if n < 0 {
    n = 0
  }
  var angles = make([]float64, n)
  for i := 0;  i < n;  i++ {
    angles[i] = DeflectionAngle(vsegs[i].D(), vsegs[i + 1].D())
  }
  return angles
}
