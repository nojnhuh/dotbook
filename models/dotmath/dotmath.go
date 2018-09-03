// Package dotmath handles the math behind calculating characteristics of
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
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// CrossCount represents data returned from CrossingCounts method
// Line: Name of the line crossed
// Side: Side of the field the line is on
// Count: The count on which the line is crossed
type CrossCount struct {
	Line  string  `json:"line"`
	Count float64 `json:"count"`
}

// CrossCounts is an array of CrossCount objects with a special sorting
// implementation
type CrossCounts []CrossCount

// NewPoint creates a new Cartesian coordinate from the provided x and y values.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Equal tests two Points for equality
func Equal(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
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

// PointOnPath returns a point that is 100*part percent of the way from a to b.
func PointOnPath(a, b Point, part float64) Point {
	return NewPoint(inBetween(a.X, b.X, part), inBetween(a.Y, b.Y, part))
}

// Distance calculates the straight-line distance between two points.
func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

// SegmentSize returns the size of each segment to get from a to b in segs steps.
func SegmentSize(a, b Point, segs float64) float64 {
	if segs == 0 {
		return 0
	}
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
	if thisX == prevX {
		return counts
	}
	for line, steps := range lines {
		if (prevX < steps) != (thisX <= steps) {
			when := math.Abs((steps - prevX) / (prevX - thisX))
			counts = append(counts, CrossCount{line, when})
		} else if thisX == steps {
			counts = append(counts, CrossCount{line, 1})
		}
	}

	sort.Slice(counts[:], func(i, j int) bool {
		return counts[i].Count < counts[j].Count
	})
	return counts
}

func crossCountEqual(c1, c2 CrossCount) bool {
	return c1.Line == c2.Line && c1.Count == c2.Count
}

func (slice CrossCounts) Len() int {
	return len(slice)
}

func (slice CrossCounts) Less(i, j int) bool {
	return slice[i].Line < slice[j].Line
}

func (slice CrossCounts) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// CrossCountSliceEqual determines if two slices of CrossCount objects are equal
func CrossCountSliceEqual(a, b CrossCounts) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	sort.Sort(a)
	sort.Sort(b)

	for i := range a {
		if !crossCountEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}
