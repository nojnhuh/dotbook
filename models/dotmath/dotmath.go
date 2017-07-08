// The dotmath package handles the math behind calculating characteristics of
// dots
package dotmath

import (
	"math"
	"sort"
)

// The Point type represents a two-dimensional Cartesian corrdinate.
// X: x-coordinate of the Point.
// Y: y-coordinate of the Point.
type Point struct {
	X, Y float64
}

// Data returned from CrossingCounts method
// Line: Name of the line crossed
// Side: Side of the field the line is on
// Count: The count on which the line is crossed
type CrossCount struct {
	Line  string
	Count float64
}

// NewPoint creates a new Cartesian coordinate from the provided x and y values.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// AddPoints adds together 2 points like vectors.
func AddPoints(a, b Point) Point {
	return NewPoint(a.X+b.X, a.Y+b.Y)
}

// ScalarMult multiplies each coordinate by the given point by the given scalar.
func ScalarMult(a Point, s float64) Point {
	return NewPoint(a.X*s, a.Y*s)
}

// inBetween calculates a float64inate that is 100*part percent of the way from
// the scalar a to b.
func inBetween(a, b, part float64) float64 {
	return a + (b-a)*part
}

// pointOnPath returns a point that is 100*part percent of the way from a to b.
func PointOnPath(a, b Point, part float64) Point {
	return NewPoint(inBetween(a.X, b.X, part), inBetween(a.Y, b.Y, part))
}

// Distance calculates the straight-line distance between two points.
func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

// SegmentSize returns the size of each segment to get from a to b in segs steps.
func SegmentSize(a, b Point, segs float64) float64 {
	return Distance(a, b) / segs
}

// BodyToFootDot calculates the final foot placement of a body-center dot.
func BodyToFootDot(p, prev Point, counts float64) Point {
	xDist := p.X - prev.X
	yDist := p.Y - prev.Y
	xAdjust := 2 * counts * xDist / (2*counts - 1)
	yAdjust := 2 * counts * yDist / (2*counts - 1)
	x := prev.X + xAdjust
	y := prev.Y + yAdjust
	return NewPoint(x, y)
}

// CrossingCounts calculates when a path between two points will cross the
// given lines. "When" is denoted by the portion from 0 to 1 from prev to p.
func CrossingCounts(p, prev Point, lines map[string]float64) []CrossCount {
	counts := []CrossCount{}
	thisX := p.X
	prevX := prev.X
	for line, steps := range lines {
		if (prevX < steps) != (thisX <= steps) {
			when := math.Abs((steps - prevX) / (prevX - thisX))
			counts = append(counts,
				CrossCount{line, when})
		}
		// steps *= -1
		// if (prevX < steps) != (thisX <= steps) && steps != 0 {
		// 	count := math.Abs((steps - prevX) / (prevX - thisX))
		// 	counts = append(counts,
		// 		CrossCount{line, count * d.MoveCounts})
		// }
	}

	sort.Slice(counts[:], func(i, j int) bool {
		return counts[i].Count < counts[j].Count
	})
	return counts
}
