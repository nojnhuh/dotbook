package models

import (
	"testing"

	"github.com/nojnhuh/dotbook/models/dotmath"
)

func TestEquals(t *testing.T) {
	p1 := &Dot{Name: "TestP1", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	p2 := &Dot{Name: "TestP2", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	var equalTests = []struct {
		d1, d2   *Dot
		expected bool
	}{
		{nil, nil, true},
		{
			nil,
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: nil},
			false,
		},
		{
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: nil},
			nil, false,
		},
		{
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: nil},
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: nil},
			true,
		},
		{
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: p1},
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: p1},
			true,
		},
		{
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: p1},
			&Dot{Name: "123A", HoldCounts: 8, MoveCounts: 16, Point: dotmath.Point{X: -12, Y: 15}, BodyCenter: false, PrevDot: p2},
			true,
		},
	}

	for _, test := range equalTests {
		v := test.d1.equals(test.d2)
		if v != test.expected {
			t.Errorf("For dots %v and %v, expected %v, got %v", test.d1, test.d2, test.expected, v)
		}
	}
}

func TestBodyToFootDot(t *testing.T) {
	d0 := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0,
		Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	var bodyToFootDotTests = []struct {
		dest, expected *Dot
	}{
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: true, PrevDot: nil},
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: true, PrevDot: nil},
		},
		{
			&Dot{Name: "VeryGoodDotName", MoveCounts: 9001, HoldCounts: 0.7, Point: dotmath.Point{X: 21, Y: 13}, BodyCenter: true, PrevDot: nil},
			&Dot{Name: "VeryGoodDotName", MoveCounts: 9001, HoldCounts: 0.7, Point: dotmath.Point{X: 21, Y: 13}, BodyCenter: true, PrevDot: nil},
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: true, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: true, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
		{
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: true, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 16, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
		{
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: d0},
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: d0},
		},
	}

	for _, test := range bodyToFootDotTests {
		v := test.dest.BodyToFootDot()
		if !v.equals(test.expected) {
			t.Errorf("From %v to %v, expected %v, got %v", test.dest.PrevDot,
				test.dest, test.expected, v)
		}
	}
}

func TestDotOnCount(t *testing.T) {
	p := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	var dotOnCountTests = []struct {
		dot       *Dot
		count     float64
		shouldErr bool
		expected  *Dot
	}{
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil},
			0, false,
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil},
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil},
			1, true, nil,
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil},
			-1, true, nil,
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
			0, false,
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
			1, true, nil,
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
			9, true, nil,
		},
		{
			&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
			-1, true, nil,
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p},
			-1, true, nil,
		},

		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			8, false,
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			4, false,
			&Dot{Name: "", MoveCounts: 4, HoldCounts: 0, Point: dotmath.Point{X: 4, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			2, false,
			&Dot{Name: "", MoveCounts: 2, HoldCounts: 0, Point: dotmath.Point{X: 2, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			1, false,
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 1, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			1, false,
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
			1, false,
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p},
		},
		{
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: true, PrevDot: p},
			1, false,
			&Dot{Name: "", MoveCounts: 1, HoldCounts: 0, Point: dotmath.Point{X: 16, Y: 0}, BodyCenter: false, PrevDot: p},
		},
	}

	for _, test := range dotOnCountTests {
		v, err := test.dot.DotOnCount(test.count)
		if (err != nil) == test.shouldErr {
			if !test.shouldErr {
				if !v.equals(test.expected) {
					t.Errorf("From %v to %v on count %f of %f, expected %v, got %v",
						test.dot.PrevDot, test.dot, test.count, test.dot.MoveCounts, test.expected, v)
				}
			}
		} else {
			t.Errorf("From %v to %v, expected error %v, got %v", test.dot.PrevDot,
				test.dot, test.shouldErr, err != nil)
		}
	}
}

func TestDistance(t *testing.T) {
	p := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	var distanceTests = []struct {
		dot      *Dot
		expected float64
	}{
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}, 0},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p}, 0},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 3, Y: 4}, BodyCenter: false, PrevDot: p}, 5},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: -3, Y: 4}, BodyCenter: false, PrevDot: p}, 5},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 3, Y: -4}, BodyCenter: false, PrevDot: p}, 5},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: -3, Y: -4}, BodyCenter: false, PrevDot: p}, 5},
	}

	for _, test := range distanceTests {
		v := test.dot.Distance()
		if v != test.expected {
			t.Errorf("From %v to %v, expected %v, got %v", test.dot.PrevDot,
				test.dot, test.expected, v)
		}
	}
}

func TestStepSize(t *testing.T) {
	p := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	field := MakeNCAAFootball(8, 5)
	var stepSizeTests = []struct {
		dot      *Dot
		expected float64
	}{
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p}, 0},
		{&Dot{Name: "", MoveCounts: 4, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p}, 4},
		{&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p}, 8},
		{&Dot{Name: "", MoveCounts: 16, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p}, 16},
		{&Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 4, Y: 0}, BodyCenter: false, PrevDot: p}, 0},
		{&Dot{Name: "", MoveCounts: 2, HoldCounts: 0, Point: dotmath.Point{X: 4, Y: 0}, BodyCenter: false, PrevDot: p}, 4},
		{&Dot{Name: "", MoveCounts: 4, HoldCounts: 0, Point: dotmath.Point{X: 4, Y: 0}, BodyCenter: false, PrevDot: p}, 8},
		{&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 4, Y: 0}, BodyCenter: false, PrevDot: p}, 16},
	}

	for _, test := range stepSizeTests {
		v := test.dot.StepSize(field)
		if v != test.expected {
			t.Errorf("From %v to %v on field %v, expected %f, got %f",
				test.dot.PrevDot, test.dot, field, test.expected, v)
		}
	}
}

func TestCrossingCount(t *testing.T) {
	p0 := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: nil}
	p1 := &Dot{Name: "", MoveCounts: 0, HoldCounts: 0, Point: dotmath.Point{X: -16, Y: 0}, BodyCenter: false, PrevDot: nil}
	field := MakeNCAAFootball(8, 5)
	var crossCountTests = []struct {
		dot      *Dot
		expected []dotmath.CrossCount
	}{
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 0}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 0, Y: 92387432}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 7.9999, Y: 0}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{
				{Line: "B45", Count: 8},
			},
		},
		{
			&Dot{Name: "", MoveCounts: 7, HoldCounts: 0, Point: dotmath.Point{X: 8, Y: 0}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{
				{Line: "B45", Count: 7},
			},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 16, Y: 0}, BodyCenter: false, PrevDot: p0},
			[]dotmath.CrossCount{
				{Line: "B45", Count: 4},
				{Line: "B40", Count: 8},
			},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 16, Y: 0}, BodyCenter: false, PrevDot: p1},
			[]dotmath.CrossCount{
				{Line: "A45", Count: 2},
				{Line: "A50", Count: 4},
				{Line: "B50", Count: 4},
				{Line: "B45", Count: 6},
				{Line: "B40", Count: 8},
			},
		},
		{
			&Dot{Name: "", MoveCounts: 8, HoldCounts: 0, Point: dotmath.Point{X: 16, Y: -34762834}, BodyCenter: false, PrevDot: p1},
			[]dotmath.CrossCount{
				{Line: "A45", Count: 2},
				{Line: "A50", Count: 4},
				{Line: "B50", Count: 4},
				{Line: "B45", Count: 6},
				{Line: "B40", Count: 8},
			},
		},
	}

	for _, test := range crossCountTests {
		v := test.dot.CrossingCounts(field)
		if !dotmath.CrossCountSliceEqual(v, test.expected) {
			t.Errorf("From %v to %v with field %v, expected %v, got %v", test.dot.PrevDot,
				test.dot, field, test.expected, v)
		}
	}
}
