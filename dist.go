package eval

import . "simplex/geom"

func dist(segs []*Segment) float64 {
    return sum(seglengths(segs))
}
