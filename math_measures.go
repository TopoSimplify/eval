package eval

import (
	"math"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/vect"
)

//Percentage Change in Angularity
func PC_Angularity(original, simple []*geom.Point) float64 {
	var angs = angularity(simple)
	var ango = angularity(original)
	return (math.Abs(sum(angs)) / math.Abs(sum(ango))) * 100.0
}

//Percentage Change in Curvilinear Segments
func PC_CurvlinearSegs(original, simple []*geom.Point) float64 {
	var angs = angularity(simple)
	var ango = angularity(original)
	var sc = float64(numDeflections(angs))
	var oc = float64(numDeflections(ango))
	return (sc / oc ) * 100.0
}

//Total Length of Vector Difference
func TL_VectorDiff(oline []*geom.Point, genidx []int) float64 {
	n := len(genidx)
	vls := make([]float64, 0)
	segidx := zip(genidx[0: n-1], genidx[1: n])

	for _, idx := range segidx {
		segvect := vect.NewVect(&vect.Options{
			A: oline[idx[0]],
			B: oline[idx[1]],
		})

		rmpnts := make([]*geom.Point, 0)
		for _, i := range _range(idx[0]+1, idx[1]) {
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
func TA_PolyDisplacement(original []*geom.Point, genidx []int) float64 {
	var as = make([]float64, 0)
	var n = len(genidx)
	var segidx = zip(genidx[0: n-1], genidx[1: n])

	for _, idx := range segidx {
		if idx[1]-idx[0] > 1 {
			var plyidx = _range(idx[0], idx[1]+1)
			var poly = make([]*geom.Point, 0)

			for _, i := range plyidx {
				poly = append(poly, original[i])
			}
			area := displacementArea(poly)
			as = append(as, area)
		}
	}
	var slo = distance(segments(original))
	return math.Log(sum(as)) / slo
}
