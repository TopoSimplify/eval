package eval

import (
    .  "simplex/geom"
)

func Angularity(coords []*Point) []float64{
    return dflns(segvects(segs(coords)))
}