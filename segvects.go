package eval

import (
    . "simplex/geom"
    . "simplex/vect"
)

//Segment vectors
func segment_vectors(segs []*Segment) []*Vect {
    var n = len(segs)
    var segv = make([]*Vect, n)
    for i := 0; i < n; i++ {
        segv[i] = NewVect(&Options{
            A : segs[i].A,
            B : segs[i].B,
        })
    }
    return segv
}
