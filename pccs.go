package eval

import (
  .  "simplex/geom"
)


//percentage change in curvilinear segments
func PCCS(origline, simpline []*Point ) float64{
  var angs, ango = Angularity(simpline), Angularity(origline)
  var a, b  = _count(angs), _count(ango)
  return  (float64(a) / float64(b)) * 100.0
}

//count change in deflection angles
func _count(angles []float64) int {
  var k , sign  int
  for i := 0; i < len(angles); i++ {
    a := angles[i]
    if i == 0 {
      k += 1
      sign = _sign(a)
    } else {
      var sa = _sign(a)//prev
      if sa != sign { //new curv
        k += 1
        sign = sa
      }
    }
  }
  return k
}

//sign 1 or -1
func _sign(i float64 ) int {
  if i >= 0 {
    return  1
  }
    return -1
}
