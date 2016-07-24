package eval

import (
    . "simplex/geom"
    . "simplex/vect"
    "math"
)

//Percentage Change in Angularity
func PCAngle(original, simple []*Point) float64 {
    var angs = angularity(simple)
    var ango = angularity(original)
    return (math.Abs(sum(angs)) / math.Abs(sum(ango))) * 100.0
}

//Percentage Change in Curvilinear Segments
func PCCS(original, simple []*Point) float64 {
    var angs = angularity(simple)
    var ango = angularity(original)
    var sc = float64(num_dflns(angs))
    var oc = float64(num_dflns(ango))
    return (sc / oc ) * 100.0
}

//Total Length of Vector Difference
func TLVD(oline []*Point, genidx []int) float64 {
    n := len(genidx)
    vls := make([]float64, 0)
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
            vls = append(vls, segvect.DistanceToPoint(pnt))
        }
    }
    var slo = distance(segments(oline))
    return math.Log(sum(vls)) / slo
}



//Total Area of Polygonal Displacement
func TAPD(original []*Point, genidx []int) float64 {
    var as = make([]float64, 0)
    var n = len(genidx)
    var segidx = zip(genidx[0: n - 1], genidx[1: n])

    for _, idx := range segidx {
        if idx[1] - idx[0] > 1 {
            var plyidx = _range(idx[0], idx[1] + 1)
            var poly = make([]*Point, 0)

            for _, i := range plyidx {
                poly = append(poly, original[i])
            }
            area := displacement_area_complex(poly)
            as = append(as, area)
        }
    }
    var slo = distance(segments(original))
    return math.Log(sum(as)) / slo
}




