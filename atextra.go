package eval

import "github.com/intdxdt/geom"

//difference indices as result of constraint
func AtDiff(constAt, normalAt []int) []int {
	if len(constAt) < len(normalAt) {
		panic("const at length is less than normal at length")
	}
	return difference(constAt, normalAt)
}

//extra vertices
func ExtraVertices(pln []*geom.Point, constAt, normalAt []int) []*geom.Point {
	var diff = AtDiff(constAt, normalAt)
	coords := make([]*geom.Point, 0)
	for _, i := range diff {
		coords = append(coords, pln[i])
	}
	return coords
}

//percentage extra vertices
func PercExtra(constAt, normalAt []int) float64 {
	var pratio = float64(len(constAt)) / float64(len(normalAt))
	if pratio < 1 {
		panic("invalid const normal ratio ")
	}
	return pratio
}
