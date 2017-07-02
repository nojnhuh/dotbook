package dotmath

import (
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	// fmt.Printf("New point %s\n", &Point{x, y})
	return Point{x, y}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func AddPoints(a, b Point) Point {
	return NewPoint(a.x+b.x, a.y+b.y)
}

func ScalarMult(a Point, s float64) Point {
	return NewPoint(a.x*s, a.y*s)
}

// inBetween calculates a float64inate that is 100*part percent of the way from
// a to b
func inBetween(a, b, part float64) float64 {
	return a + (b-a)*part
}

// pointOnPath returns a point that is 100*part percent of the way from a to b
func PointOnPath(a, b Point, part float64) Point {
	return Point{inBetween(a.x, b.x, part), inBetween(a.y, b.y, part)}
}

// Distance calculates the straight-line distance between two points
func Distance(a, b Point) float64 {
	// fmt.Printf("Distance: %.2f\n", math.Sqrt(math.Pow(a.x-b.x, 2)+math.Pow(a.y-b.y, 2)))
	// fmt.Printf("X Distance: %.2f\n", math.Abs(a.x-b.x))
	// fmt.Printf("Y Distance: %.2f\n", math.Abs(a.y-b.y))
	// fmt.Printf("a.x: %.2f\n", a.x)
	// fmt.Printf("b.x: %.2f\n", b.x)
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// SegmentSize returns the size of each segment to get from a to b in segs steps
func SegmentSize(a, b Point, segs float64) float64 {
	// fmt.Printf("Counts: %.0f\n", segs)
	return Distance(a, b) / segs
}

func BodyToFootDot(p, prev Point, counts float64) Point {
	xDist := p.x - prev.x
	yDist := p.y - prev.y
	xAdjust := 2 * counts * xDist / (2*counts - 1)
	yAdjust := 2 * counts * yDist / (2*counts - 1)
	x := prev.x + xAdjust
	y := prev.y + yAdjust
	return Point{x, y}
}
