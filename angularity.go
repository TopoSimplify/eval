package eval

import (
    .  "simplex/geom"
)

func angularity(coords []*Point) []float64{
    return deflections(
        segment_vectors(
            segments(coords)))
}