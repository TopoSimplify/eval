package eval

import (
	. "github.com/intdxdt/geom"
	. "github.com/intdxdt/vect"
)


//Create segments
func segments(coords []*Point) []*Segment {
	segs := make([]*Segment, len(coords)-1)
	for i := 0; i < len(coords)-1; i++ {
		j := i + 1
		segs[i] = NewSegment(coords[i], coords[j])
	}
	return segs
}


//Segment vectors
func segment_vectors(segs []*Segment) []*Vect {
	var n = len(segs)
	var segv = make([]*Vect, n)
	for i := 0; i < n; i++ {
		segv[i] = NewVect(&Options{
			A: segs[i].A,
			B: segs[i].B,
		})
	}
	return segv
}


//segments length
func seg_lengths(segs []*Segment) []float64 {
	dists := make([]float64, len(segs))
	for i, v := range segs {
		dists[i] = v.A.Distance(v.B)
	}
	return dists
}
