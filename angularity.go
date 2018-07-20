package eval

import (
	"github.com/intdxdt/geom"
)

func angularity(coords []geom.Point) []float64 {
	return deflections(segmentVectors(segments(coords)))
}
