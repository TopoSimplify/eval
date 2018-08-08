package eval

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/vect"
)

//Create segments
func segments(coords []geom.Point) []*geom.Segment {
	segs := make([]*geom.Segment, len(coords)-1)
	for i := 0; i < len(coords)-1; i++ {
		j := i + 1
		segs[i] = geom.NewSegmentAB(&coords[i], &coords[j])
	}
	return segs
}

//Segment vectors
func segmentVectors(segs []*geom.Segment) []*vect.Vect {
	var n = len(segs)
	var segv = make([]*vect.Vect, n)
	for i := 0; i < n; i++ {
		segv[i] = vect.NewVect( *segs[i].A, *segs[i].B)
	}
	return segv
}

//segments length
func segLengths(segs []*geom.Segment) []float64 {
	dists := make([]float64, len(segs))
	for i, v := range segs {
		dists[i] = v.A.Distance(v.B)
	}
	return dists
}
