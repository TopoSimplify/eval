package eval

import "github.com/intdxdt/geom"

func distance(segs []*geom.Segment) float64 {
	return sum(segLengths(segs))
}
