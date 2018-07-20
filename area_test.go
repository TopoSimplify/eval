package eval

import (
	"testing"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
)

func TestWKT(t *testing.T) {
	var g = goblin.Goblin(t)
	var areaA, areaB = 31263.0267879, 30875.6418663
	var wkta = "LINESTRING ( 280 320, 300 360, 280 420, 340 460, 440 440, 440 500, 480 500, 560 400, 520 360, 620 380, 619.9847625162381 453.03751576722357, 680 400, 740 500 )"
	var wktb = "LINESTRING ( 280 320, 300 360, 280 420, 340 460, 440 440, 440 500, 480 500, 560 400, 520 360, 620 380, 620 480, 680 400, 740 500 )"
	g.Describe("Diplacement Area", func() {
		g.It("test displacement area", func() {
			g.Assert(areaA).Eql(math.Round(
				displacementArea(geom.NewLineStringFromWKT(wkta).Coordinates()),
				7,
			))
			g.Assert(areaB).Eql(math.Round(
				displacementArea(geom.NewLineStringFromWKT(wktb).Coordinates()),
				7,
			))
		})
	})
}
