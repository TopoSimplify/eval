package eval

import (
    . "simplex/geom"
    . "simplex/vect"
)

//Total Length of Vector Difference
func TLVD(oline []*Point, genidx []int) float64 {
    n := len(genidx)
    dispmts := make([]float64, 0)
    segidx := zip(genidx[0: n - 1], genidx[1: n])

    for _, idx := range segidx {
        segvect := NewVect(&Options{
            A : oline[idx[0]],
            B : oline[idx[1]],
        })

        rmpnts := make([]*Point, 0)
        for _, i := range _range(idx[0] + 1, idx[1]) {
            rmpnts = append(rmpnts, oline[i])
        }

        for _, pnt := range rmpnts {
            dispmts = append(dispmts, segvect.DistanceToPoint(pnt))
        }
    }

    var slo = dist(Segments(oline))
    var vls = sum(dispmts)
    return vls / slo
}
