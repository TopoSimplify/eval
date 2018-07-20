package eval

import (
	"github.com/intdxdt/vect"
	"github.com/intdxdt/geom"
)

//computes the deflections of a series of vectors
func deflections(vs []*vect.Vect) []float64 {
	var n = len(vs) - 1
	if n < 0 {
		n = 0
	}
	var angles = make([]float64, n)
	for i := 0; i < n; i++ {
		angles[i] = deflection(vs[i], vs[i+1])
	}
	return angles
}

//deflection between two vectors
func deflection(a, b *vect.Vect) float64 {
	return geom.DeflectionAngle(a.Direction(), b.Direction())
}
