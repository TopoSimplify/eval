package eval

import (
	"github.com/intdxdt/geom"
)

func displacement_area(coords []*geom.Point) float64 {
	n := len(coords)
	var sub_seg *geom.Segment
	blade := geom.NewSegment(coords[0], coords[n-1])
	subpolys := make([][]*geom.Segment, 0)
	segs := segments(coords)
	cur := make([]*geom.Segment, 0)

	for i := 0; i < len(segs); i++ {
		if i == 0 || i == len(segs)-1 {
			cur = append(cur, segs[i])
		} else {
			s := segs[i]
			pnts, bln := blade.SegSegIntersection(s, false)
			if bln {
				sub_seg = geom.NewSegment(s.A, pnts[0])
				cur = append(cur, sub_seg)
				//---------------------------------
				subpolys = append(subpolys, cur)
				cur = make([]*geom.Segment, 0)
				//---------------------------------
				sub_seg = geom.NewSegment(pnts[len(pnts)-1], s.B)
				cur = append(cur, sub_seg)
			} else {
				sub_seg = segs[i]
				cur = append(cur, sub_seg)
			}
		}
	}

	if len(cur) > 0 {
		subpolys = append(subpolys, cur)
	}
	var areas float64
	for _, segs := range subpolys {
		p := segs_to_polygon(segs)
		areas += p.Area()
	}
	return areas
}

func segs_to_polygon(segs []*geom.Segment) *geom.Polygon {
	var coords = make([]*geom.Point, 0)
	for _, seg := range segs {
		coords = append(coords, seg.A)
		coords = append(coords, seg.B)
	}
	return geom.NewPolygon(coords)
}
