package eval

import "github.com/intdxdt/geom"

//difference indices as result of constraint
func AtDiff(const_at, normal_at []int) []int {
	if len(const_at) < len(normal_at) {
		panic("const at length is less than normal at length")
	}
	return difference(const_at, normal_at)
}

//extra vertices
func ExtraVertices(pln []*geom.Point, const_at, normal_at []int) []*geom.Point {
	var diff = AtDiff(const_at, normal_at)
	coords := make([]*geom.Point, 0)
	for _, i := range diff {
		coords = append(coords, pln[i])
	}
	return coords
}

//percentage extra vertices
func PercExtra(const_at, normal_at []int) float64 {
	var pratio = float64(len(const_at)) / float64(len(normal_at))
	if pratio < 1 {
		panic("invalid const normal ratio ")
	}
	return pratio
}
