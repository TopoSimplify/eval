package eval

import (
    . "simplex/vect"
)

//computes the deflections of a series of vectors
func deflections(vs []*Vect) []float64 {
    var n = len(vs) - 1
    if n < 0 {
        n = 0
    }
    var angles = make([]float64, n)
    for i := 0; i < n; i++ {
        angles[i] = defln(vs[i], vs[i + 1])
    }
    return angles
}

//deflection between two vectors
func defln(a, b *Vect) float64 {
    return DeflectionAngle(a.D(), b.D())
}