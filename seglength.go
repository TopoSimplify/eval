package eval

import (
    . "simplex/geom"
)

//segments length
func seglengths(segs []*Segment) []float64 {
    dists := make([]float64, len(segs))
    for i, v := range segs {
        dists[i] = v.A.Distance(v.B)
    }
    return dists
}
