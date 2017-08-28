package eval

import (
	"simplex/geom"
)

func angularity(coords []*geom.Point) []float64 {
	return deflections(segment_vectors(segments(coords)))
}
