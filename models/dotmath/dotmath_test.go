package dotmath

import "testing"

type inBetweenTest struct {
	points []float64
	part   float64
	mid    float64
}

type pointOnPathTest struct {
	points []Point
	part   float64
	mid    Point
}

var inBetweenTests = []inBetweenTest{
	{[]float64{0, 0}, 1.0, 0},
	{[]float64{0, 0}, 0.5, 0},
	{[]float64{0, 0}, 0.25, 0},
	{[]float64{0, 0}, 0.125, 0},
	{[]float64{0, 0}, 0.0, 0},

	{[]float64{0, 4}, 1.0, 4},
	{[]float64{0, 4}, 0.5, 2},
	{[]float64{0, 4}, 0.25, 1},
	{[]float64{0, 4}, 0.125, 0.5},
	{[]float64{0, 4}, 0.0, 0},
}

var pointOnPathTests = []pointOnPathTest{
	{[]Point{Point{0, 0}, Point{4, 4}}, 1.0, Point{4, 4}},
	{[]Point{Point{0, 0}, Point{4, 4}}, 0.5, Point{2, 2}},
	{[]Point{Point{0, 0}, Point{4, 4}}, 0.25, Point{1, 1}},
	{[]Point{Point{0, 0}, Point{4, 4}}, 0.125, Point{0.5, 0.5}},
	{[]Point{Point{0, 0}, Point{4, 4}}, 0.0, Point{0, 0}},

	{[]Point{Point{4, 4}, Point{0, 0}}, 1.0, Point{0, 0}},
	{[]Point{Point{4, 4}, Point{0, 0}}, 0.5, Point{2, 2}},
	{[]Point{Point{4, 4}, Point{0, 0}}, 0.25, Point{3, 3}},
	{[]Point{Point{4, 4}, Point{0, 0}}, 0.125, Point{3.5, 3.5}},
	{[]Point{Point{4, 4}, Point{0, 0}}, 0.0, Point{4, 4}},

	{[]Point{Point{0, 0}, Point{-4, -4}}, 1.0, Point{-4, -4}},
	{[]Point{Point{0, 0}, Point{-4, -4}}, 0.5, Point{-2, -2}},
	{[]Point{Point{0, 0}, Point{-4, -4}}, 0.25, Point{-1, -1}},
	{[]Point{Point{0, 0}, Point{-4, -4}}, 0.125, Point{-0.5, -0.5}},
	{[]Point{Point{0, 0}, Point{-4, -4}}, 0.0, Point{0, 0}},

	{[]Point{Point{0, 0}, Point{0, 0}}, 1.0, Point{0, 0}},
	{[]Point{Point{0, 0}, Point{0, 0}}, 0.5, Point{0, 0}},
	{[]Point{Point{0, 0}, Point{0, 0}}, 0.25, Point{0, 0}},
	{[]Point{Point{0, 0}, Point{0, 0}}, 0.125, Point{0, 0}},
	{[]Point{Point{0, 0}, Point{0, 0}}, 0.0, Point{0, 0}},
}

func TestPointOnPath(t *testing.T) {
	for _, test := range pointOnPathTests {
		v := PointOnPath(test.points[0], test.points[1], test.part)
		if v != test.mid {
			t.Error(
				"For", test.points,
				"with part", test.part,
				"expected", test.mid,
				"got", v,
			)
		}
	}
}

func TestInBetween(t *testing.T) {
	for _, test := range inBetweenTests {
		v := inBetween(test.points[0], test.points[1], test.part)
		if v != test.mid {
			t.Error(
				"For", test.points,
				"with part", test.part,
				"expected", test.mid,
				"got", v,
			)
		}
	}
}
