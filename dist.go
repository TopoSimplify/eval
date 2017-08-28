package eval

import . "simplex/geom"

func distance(segs []*Segment) float64 {
	return sum(seglengths(segs))
}
