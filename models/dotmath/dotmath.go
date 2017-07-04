package dotmath

import (
	"math"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	// fmt.Printf("New point %s\n", &Point{x, y})
	return Point{x, y}
}

func AddPoints(a, b Point) Point {
	return NewPoint(a.X+b.X, a.Y+b.Y)
}

func ScalarMult(a Point, s float64) Point {
	return NewPoint(a.X*s, a.Y*s)
}

// inBetween calculates a float64inate that is 100*part percent of the way from
// a to b
func inBetween(a, b, part float64) float64 {
	return a + (b-a)*part
}

// pointOnPath returns a point that is 100*part percent of the way from a to b
func PointOnPath(a, b Point, part float64) Point {
	return Point{inBetween(a.X, b.X, part), inBetween(a.Y, b.Y, part)}
}

// Distance calculates the straight-line distance between two points
func Distance(a, b Point) float64 {
	// fmt.Printf("Distance: %.2f\n", math.Sqrt(math.Pow(a.X-b.X, 2)+math.Pow(a.Y-b.Y, 2)))
	// fmt.Printf("X Distance: %.2f\n", math.Abs(a.X-b.X))
	// fmt.Printf("Y Distance: %.2f\n", math.Abs(a.Y-b.Y))
	// fmt.Printf("a.X: %.2f\n", a.X)
	// fmt.Printf("b.X: %.2f\n", b.X)
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

// SegmentSize returns the size of each segment to get from a to b in segs steps
func SegmentSize(a, b Point, segs float64) float64 {
	// fmt.Printf("Counts: %.0f\n", segs)
	return Distance(a, b) / segs
}

func BodyToFootDot(p, prev Point, counts float64) Point {
	xDist := p.X - prev.X
	yDist := p.Y - prev.Y
	xAdjust := 2 * counts * xDist / (2*counts - 1)
	yAdjust := 2 * counts * yDist / (2*counts - 1)
	x := prev.X + xAdjust
	y := prev.Y + yAdjust
	return Point{x, y}
}
