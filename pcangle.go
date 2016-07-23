package eval

import (
    . "simplex/geom"
    "math"
)

//percentage change in angularity
func PCAngle(original, simple []*Point) float64{
    var ang_s = Angularity(simple)
    var ang_o = Angularity(original)
    var abs_angs = math.Abs(sum(ang_s))
    var abs_ango = math.Abs(sum(ang_o))
    return (abs_angs / abs_ango) * 100.0
}
