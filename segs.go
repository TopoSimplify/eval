package eval

import (
    . "simplex/geom"
)

//Create segments
func Segments(coords []*Point) []*Segment {
    segs := make([]*Segment, len(coords) - 1)
    for i := 0; i < len(coords) - 1; i++ {
        segs[i] = NewSegment(coords[i], coords[i + 1])
    }
    return segs
}
