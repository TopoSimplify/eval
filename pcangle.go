package eval

import (
    . "simplex/geom"
    "math"
)

//percentage change in angularity
func PCAngle(origline, simpline []*Point) {
    var angs = Angularity(simpline)
    var ango = Angularity(origline)
    var abs_angs = math.Abs(sum(angs))
    var abs_ango = math.Abs(sum(ango))
    return (abs_angs / abs_ango) * 100.0
}
