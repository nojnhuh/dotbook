package dotmath

import (
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	var equalTests = []struct {
		a, b     Point
		expected bool
	}{
		{Point{0, 0}, Point{0, 0}, true},
		{Point{0, 0}, Point{-0, -0}, true},
		{Point{0, 0}, Point{-1 * math.SmallestNonzeroFloat64, 0}, false},
		{Point{0, 0}, Point{0, -1 * math.SmallestNonzeroFloat64}, false},
		{Point{0, 0}, Point{math.SmallestNonzeroFloat64, 0}, false},
		{Point{0, 0}, Point{0, math.SmallestNonzeroFloat64}, false},
		{Point{math.MaxFloat64, math.MaxFloat64}, Point{math.MaxFloat64, math.MaxFloat64}, true},
	}

	for _, test := range equalTests {
		v := equal(test.a, test.b)
		if v != test.expected {
			t.Errorf("For points %v and %v, expected %v, got %v", test.a, test.b,
				test.expected, v)
		}
	}
}

func equal(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func TestNewPoint(t *testing.T) {
	var newPointTests = []struct {
		x, y     float64
		expected Point
	}{
		{0, 0, Point{0, 0}},
		{0, 1, Point{0, 1}},
		{1, 0, Point{1, 0}},
		{0, -1, Point{0, -1}},
		{-1, 0, Point{-1, 0}},
	}

	for _, test := range newPointTests {
		v := NewPoint(test.x, test.y)
		if !equal(v, test.expected) {
			t.Errorf("For coordinates X = %f and Y = %f, expected %v, got %v",
				test.x, test.y, test.expected, v)
		}
	}
}

func TestAddPoints(t *testing.T) {
	var addPointTests = []struct {
		a, b     Point
		expected Point
	}{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{Point{0, 0}, Point{0, 1}, Point{0, 1}},
		{Point{0, 0}, Point{1, 0}, Point{1, 0}},
		{Point{0, 0}, Point{0, -1}, Point{0, -1}},
		{Point{0, 0}, Point{-1, 0}, Point{-1, 0}},
		{Point{3, 21}, Point{16, 5}, Point{19, 26}},
		{Point{-3, -21}, Point{-16, -5}, Point{-19, -26}},
		{Point{-3, -21}, Point{16, 5}, Point{13, -16}},
		{Point{3, -21}, Point{-16, 5}, Point{-13, -16}},
	}

	for _, test := range addPointTests {
		v := AddPoints(test.a, test.b)
		if !equal(v, test.expected) {
			t.Errorf("For point %v plus %v, expected %v, got %v", test.a, test.b,
				test.expected, v)
		}
	}
}

func TestScalarMult(t *testing.T) {
	var scalarMultTests = []struct {
		point    Point
		s        float64
		expected Point
	}{
		{Point{0, 0}, 0, Point{0, 0}},
		{Point{0, 0}, 1, Point{0, 0}},
		{Point{0, 0}, 7, Point{0, 0}},
		{Point{1, 1}, 0, Point{0, 0}},
		{Point{43, 22}, 0, Point{0, 0}},
		{Point{1, 1}, 1, Point{1, 1}},
		{Point{1, 1}, -1, Point{-1, -1}},
		{Point{98, 0.00005}, 1, Point{98, 0.00005}},
		{Point{1, 1}, 9, Point{9, 9}},
		{Point{-1, -1}, 9, Point{-9, -9}},
		{Point{1, 1}, 0.09, Point{0.09, 0.09}},
		{Point{2, 1}, 0.09, Point{0.18, 0.09}},
		{Point{1, 2}, 0.09, Point{0.09, 0.18}},
	}

	for _, test := range scalarMultTests {
		v := ScalarMult(test.point, test.s)
		if !equal(v, test.expected) {
			t.Errorf("For point %v multiplied by %f, expected %v, got %v",
				test.point, test.s, test.expected, v)
		}
	}
}

func TestInBetween(t *testing.T) {
	var inBetweenTests = []struct {
		a, b, part, expected float64
	}{
		{0, 0, 1.0, 0},
		{0, 0, 0.5, 0},
		{0, 0, 0.25, 0},
		{0, 0, 0.125, 0},
		{0, 0, 0.0, 0},

		{0, 4, 1.0, 4},
		{0, 4, 0.5, 2},
		{0, 4, 0.25, 1},
		{0, 4, 0.125, 0.5},
		{0, 4, 0.0, 0},
	}

	for _, test := range inBetweenTests {
		v := inBetween(test.a, test.b, test.part)
		if v != test.expected {
			t.Errorf("For %v with part %v, expected %v, got %v", test.a, test.b,
				test.part, test.expected, v)
		}
	}
}

func TestPointOnPath(t *testing.T) {
	var pointOnPathTests = []struct {
		a, b     Point
		part     float64
		expected Point
	}{
		{Point{0, 0}, Point{4, 4}, 1.0, Point{4, 4}},
		{Point{0, 0}, Point{4, 4}, 0.5, Point{2, 2}},
		{Point{0, 0}, Point{4, 4}, 0.25, Point{1, 1}},
		{Point{0, 0}, Point{4, 4}, 0.125, Point{0.5, 0.5}},
		{Point{0, 0}, Point{4, 4}, 0.0, Point{0, 0}},

		{Point{4, 4}, Point{0, 0}, 1.0, Point{0, 0}},
		{Point{4, 4}, Point{0, 0}, 0.5, Point{2, 2}},
		{Point{4, 4}, Point{0, 0}, 0.25, Point{3, 3}},
		{Point{4, 4}, Point{0, 0}, 0.125, Point{3.5, 3.5}},
		{Point{4, 4}, Point{0, 0}, 0.0, Point{4, 4}},

		{Point{0, 0}, Point{-4, -4}, 1.0, Point{-4, -4}},
		{Point{0, 0}, Point{-4, -4}, 0.5, Point{-2, -2}},
		{Point{0, 0}, Point{-4, -4}, 0.25, Point{-1, -1}},
		{Point{0, 0}, Point{-4, -4}, 0.125, Point{-0.5, -0.5}},
		{Point{0, 0}, Point{-4, -4}, 0.0, Point{0, 0}},

		{Point{0, 0}, Point{0, 0}, 1.0, Point{0, 0}},
		{Point{0, 0}, Point{0, 0}, 0.5, Point{0, 0}},
		{Point{0, 0}, Point{0, 0}, 0.25, Point{0, 0}},
		{Point{0, 0}, Point{0, 0}, 0.125, Point{0, 0}},
		{Point{0, 0}, Point{0, 0}, 0.0, Point{0, 0}},
	}

	for _, test := range pointOnPathTests {
		v := PointOnPath(test.a, test.b, test.part)
		if !equal(v, test.expected) {
			t.Errorf("For points %v and %v with part %f, expected %v, got %v",
				test.a, test.b, test.part, test.expected, v)
		}
	}
}

func TestDistance(t *testing.T) {
	var distanceTests = []struct {
		a, b     Point
		expected float64
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{10, 12}, Point{10, 12}, 0},

		{Point{0, 0}, Point{10, 0}, 10},
		{Point{0, 0}, Point{0, 10}, 10},
		{Point{0, 0}, Point{-10, 0}, 10},
		{Point{0, 0}, Point{0, -10}, 10},
		{Point{10, 0}, Point{0, 0}, 10},
		{Point{-10, 0}, Point{0, 0}, 10},
		{Point{0, 10}, Point{0, 0}, 10},
		{Point{0, -10}, Point{0, 0}, 10},

		{Point{0, 0}, Point{3, 4}, 5},
		{Point{0, 0}, Point{-3, -4}, 5},
		{Point{0, 0}, Point{-3, 4}, 5},
		{Point{0, 0}, Point{3, -4}, 5},
		{Point{3, 4}, Point{0, 0}, 5},
		{Point{-3, 4}, Point{0, 0}, 5},
		{Point{3, -4}, Point{0, 0}, 5},
		{Point{-3, -4}, Point{0, 0}, 5},
	}

	for _, test := range distanceTests {
		v := Distance(test.a, test.b)
		if v != test.expected {
			t.Errorf("For points %v and %v, expected %f, got %f", test.a, test.b, test.expected, v)
		}
	}
}

func TestSegmentSize(t *testing.T) {
	var segmentSizeTests = []struct {
		a, b     Point
		segs     float64
		expected float64
	}{
		{Point{0, 0}, Point{0, 0}, 0, 0},
		{Point{0, 0}, Point{0, 0}, 3, 0},
		{Point{0, 0}, Point{0, 3}, 3, 1},
		{Point{0, 0}, Point{4, 3}, 5, 1},
		{Point{0, 0}, Point{4, 3}, 10, 0.5},
		{Point{0, 0}, Point{4, 3}, 2.5, 2},
	}

	for _, test := range segmentSizeTests {
		v := SegmentSize(test.a, test.b, test.segs)
		if v != test.expected {
			t.Errorf("For points %v and %v in %.0f segments, expected %f, got %f", test.a, test.b, test.segs, test.expected, v)
		}
	}
}

func TestBodyToFoot(t *testing.T) {
	var bodyToFootTests = []struct {
		from, to Point
		counts   float64
		expected Point
	}{
		{Point{0, 0}, Point{0, 0}, 0, Point{0, 0}},
		{Point{8, 7}, Point{0, 0}, 0, Point{8, 7}},
		{Point{0, 0}, Point{0, 8}, 0, Point{0, 0}},
		{Point{0, 0}, Point{8, 0}, 1, Point{16, 0}},
		{Point{0, 0}, Point{8, 0}, 2, Point{32. / 3, 0}},
		{Point{0, 0}, Point{8, 0}, 8, Point{128. / 15, 0}},
		{Point{0, 0}, Point{0, 8}, 8, Point{0, 128. / 15}},
		{Point{0, 0}, Point{3, 4}, 1, Point{6, 8}},
		{Point{0, 0}, Point{3, 4}, 2, Point{4, 16. / 3}},
	}

	for _, test := range bodyToFootTests {
		v := BodyToFootDot(test.to, test.from, test.counts)
		if !equal(v, test.expected) {
			t.Errorf("From %v to %v in %.0f counts, expected %v, got %v", test.from, test.to, test.counts, test.expected, v)
		}
	}
}

func crossCountEqual(c1, c2 CrossCount) bool {
	return c1.Line == c2.Line && c1.Count == c2.Count
}

func crossCountSliceEqual(a, b []CrossCount) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !crossCountEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestCrossingCounts(t *testing.T) {
	lines := map[string]float64{
		"A40": -16,
		"A45": -8,
		"50":  0,
		"B45": 8,
		"B40": 16,
	}
	var crossingCountTests = []struct {
		from, to Point
		expected []CrossCount
	}{
		{Point{0, 0}, Point{0, 0}, []CrossCount{}},
		{Point{0, -16}, Point{0, 16}, []CrossCount{}},
		{Point{-16, 0}, Point{16, 0}, []CrossCount{
			{"A45", 0.25},
			{"50", 0.5},
			{"B45", 0.75},
			{"B40", 1},
		}},
		{Point{0, 0}, Point{8, 0}, []CrossCount{
			{"B45", 1},
		}},
		{Point{2, 0}, Point{6, 0}, []CrossCount{}},
		{Point{-1, 0}, Point{1, 0}, []CrossCount{
			{"50", 0.5},
		}},
		{Point{-1, 1}, Point{1, 1}, []CrossCount{
			{"50", 0.5},
		}},
		{Point{-1, -9.6456}, Point{1, 4534}, []CrossCount{
			{"50", 0.5},
		}},
	}

	for _, test := range crossingCountTests {
		v := CrossingCounts(test.to, test.from, lines)
		if !crossCountSliceEqual(v, test.expected) {
			t.Errorf("From %v to %v with lines %v, expected %v, got %v", test.from,
				test.to, lines, test.expected, v)
		}
	}
}
