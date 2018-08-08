package eval

import (
	"github.com/intdxdt/geom"
)

func displacementArea(coords []geom.Point) float64 {
	var n = len(coords)
	var subSeg *geom.Segment
	var blade = geom.NewSegmentAB(&coords[0], &coords[n-1])
	var subpolys = make([][]*geom.Segment, 0)
	var segs = segments(coords)
	var cur = make([]*geom.Segment, 0)

	for i := 0; i < len(segs); i++ {
		if i == 0 || i == len(segs)-1 {
			cur = append(cur, segs[i])
		} else {
			s := segs[i]
			pnts := blade.SegSegIntersection(s)
			if len(pnts) > 0 {
				subSeg = geom.NewSegmentAB(s.A, &pnts[0].Point)
				cur = append(cur, subSeg)
				//---------------------------------
				subpolys = append(subpolys, cur)
				cur = make([]*geom.Segment, 0)
				//---------------------------------
				subSeg = geom.NewSegmentAB(&pnts[len(pnts)-1].Point, s.B)
				cur = append(cur, subSeg)
			} else {
				subSeg = segs[i]
				cur = append(cur, subSeg)
			}
		}
	}

	if len(cur) > 0 {
		subpolys = append(subpolys, cur)
	}
	var areas float64
	for _, segs := range subpolys {
		p := segsToPolygon(segs)
		areas += p.Area()
	}
	return areas
}

func segsToPolygon(segs []*geom.Segment) *geom.Polygon {
	var coords = make([]geom.Point, 0)
	for _, seg := range segs {
		coords = append(coords, *seg.A)
		coords = append(coords, *seg.B)
	}
	return geom.NewPolygon(coords)
}
