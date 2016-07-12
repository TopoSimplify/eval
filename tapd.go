package eval

import (
  . "simplex/geom"
)

//Total Area of Polygonal Displacement
func TAPD(oline []*Point, genidx []int) {
  var n       = len(genidx)
  var dispmts = make([]float64, n)
  var segidx  = zip(genidx[0: n - 1], genidx[1: n])

  for _, idx := range segidx {
    if idx[1] - idx[0] > 1 {
      var plyidx  = _range(idx[0], idx[1] + 1)
      var poly    = make([]*Point , 0)

      for _, i := range plyidx {
        poly = append(poly, oline[i])
      }
      dispmts = append(dispmts, NewLinearRing(poly).Area())
    }
  }

  var slo = dist(segs(oline))
  var as  = sum(dispmts)
  return as / slo
}

