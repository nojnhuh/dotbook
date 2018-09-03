package models

import (
	"reflect"
	"testing"
)

func TestMakeFootball(t *testing.T) {
	stepsBetweenLines := 16.
	field := makeFootball(stepsBetweenLines, 10, nil)
	if field.Height != 84 || field.Width != 160 {
		t.Errorf("Incorrect dimensions: %.0fx%.0f, expected 160x84", field.Width, field.Height)
	}
	if field.StepsBetweenLines != stepsBetweenLines {
		t.Errorf("Incorrect StepsBetweenLines: %.0f, expected %.0f", field.StepsBetweenLines, stepsBetweenLines)
	}

	var sideToSideLines = map[string]float64{
		"A0":  -5 * stepsBetweenLines,
		"A10": -4 * stepsBetweenLines,
		"A20": -3 * stepsBetweenLines,
		"A30": -2 * stepsBetweenLines,
		"A40": -1 * stepsBetweenLines,
		"A50": -0 * stepsBetweenLines,
		"B50": 0 * stepsBetweenLines,
		"B40": 1 * stepsBetweenLines,
		"B30": 2 * stepsBetweenLines,
		"B20": 3 * stepsBetweenLines,
		"B10": 4 * stepsBetweenLines,
		"B0":  5 * stepsBetweenLines,
	}
	if !reflect.DeepEqual(field.SideToSideLines, sideToSideLines) {
		t.Errorf("Incorrect side-to-side lines. Expected %v, got %v", field.SideToSideLines, sideToSideLines)
	}
}

func TestMakeNCAAFootball(t *testing.T) {
	stepsBetweenLines := 8.
	yardsBetweenLines := 5.
	ncaa := MakeNCAAFootball(stepsBetweenLines, yardsBetweenLines)
	if ncaa.Height != 84 || ncaa.Width != 160 {
		t.Errorf("Incorrect dimensions: %.0fx%.0f, expected 160x84", ncaa.Width, ncaa.Height)
	}
	if ncaa.StepsBetweenLines != stepsBetweenLines {
		t.Errorf("Incorrect StepsBetweenLines: %.0f, expected %.0f", ncaa.StepsBetweenLines, stepsBetweenLines)
	}
	var sideToSideLines = map[string]float64{
		"A0":  -10 * stepsBetweenLines,
		"A5":  -9 * stepsBetweenLines,
		"A10": -8 * stepsBetweenLines,
		"A15": -7 * stepsBetweenLines,
		"A20": -6 * stepsBetweenLines,
		"A25": -5 * stepsBetweenLines,
		"A30": -4 * stepsBetweenLines,
		"A35": -3 * stepsBetweenLines,
		"A40": -2 * stepsBetweenLines,
		"A45": -1 * stepsBetweenLines,
		"A50": -0 * stepsBetweenLines,
		"B50": 0 * stepsBetweenLines,
		"B45": 1 * stepsBetweenLines,
		"B40": 2 * stepsBetweenLines,
		"B35": 3 * stepsBetweenLines,
		"B30": 4 * stepsBetweenLines,
		"B25": 5 * stepsBetweenLines,
		"B20": 6 * stepsBetweenLines,
		"B15": 7 * stepsBetweenLines,
		"B10": 8 * stepsBetweenLines,
		"B5":  9 * stepsBetweenLines,
		"B0":  10 * stepsBetweenLines,
	}
	if !reflect.DeepEqual(ncaa.SideToSideLines, sideToSideLines) {
		t.Errorf("Incorrect side-to-side lines. Expected %v, got %v", ncaa.SideToSideLines, sideToSideLines)
	}

	var frontToBackLines = map[string]float64{
		"FSL": 0,
		"FH":  4 * stepsBetweenLines,
		"BH":  6.5 * stepsBetweenLines,
		"BSL": 10.5 * stepsBetweenLines,
	}
	if !reflect.DeepEqual(ncaa.FrontToBackLines, frontToBackLines) {
		t.Errorf("Incorrect front-to-back lines. Expected %v, got %v", ncaa.FrontToBackLines, frontToBackLines)
	}

}

func TestMakeHSFootball(t *testing.T) {
	stepsBetweenLines := 8.
	yardsBetweenLines := 5.
	hs := MakeHSFootball(stepsBetweenLines, yardsBetweenLines)
	if hs.Height != 84 || hs.Width != 160 {
		t.Errorf("Incorrect dimensions: %.0fx%.0f, expected 160x84", hs.Width, hs.Height)
	}
	if hs.StepsBetweenLines != stepsBetweenLines {
		t.Errorf("Incorrect StepsBetweenLines: %.0f, expected %.0f", hs.StepsBetweenLines, stepsBetweenLines)
	}
	var sideToSideLines = map[string]float64{
		"A0":  -10 * stepsBetweenLines,
		"A5":  -9 * stepsBetweenLines,
		"A10": -8 * stepsBetweenLines,
		"A15": -7 * stepsBetweenLines,
		"A20": -6 * stepsBetweenLines,
		"A25": -5 * stepsBetweenLines,
		"A30": -4 * stepsBetweenLines,
		"A35": -3 * stepsBetweenLines,
		"A40": -2 * stepsBetweenLines,
		"A45": -1 * stepsBetweenLines,
		"A50": -0 * stepsBetweenLines,
		"B50": 0 * stepsBetweenLines,
		"B45": 1 * stepsBetweenLines,
		"B40": 2 * stepsBetweenLines,
		"B35": 3 * stepsBetweenLines,
		"B30": 4 * stepsBetweenLines,
		"B25": 5 * stepsBetweenLines,
		"B20": 6 * stepsBetweenLines,
		"B15": 7 * stepsBetweenLines,
		"B10": 8 * stepsBetweenLines,
		"B5":  9 * stepsBetweenLines,
		"B0":  10 * stepsBetweenLines,
	}
	if !reflect.DeepEqual(hs.SideToSideLines, sideToSideLines) {
		t.Errorf("Incorrect side-to-side lines. Expected %v, got %v", hs.SideToSideLines, sideToSideLines)
	}

	var frontToBackLines = map[string]float64{
		"FSL": 0,
		"FH":  3.5 * stepsBetweenLines,
		"BH":  7 * stepsBetweenLines,
		"BSL": 10.5 * stepsBetweenLines,
	}
	if !reflect.DeepEqual(hs.FrontToBackLines, frontToBackLines) {
		t.Errorf("Incorrect front-to-back lines. Expected %v, got %v", hs.FrontToBackLines, frontToBackLines)
	}

}

func TestClosestLine(t *testing.T) {
	var closestSideToSideLineTests = []struct {
		lines    map[string]float64
		steps    float64
		expected string
	}{
		// Side to Side line tests
		{MakeNCAAFootball(8, 5).SideToSideLines, 0, "B50"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 4, "B50"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 4.1, "B45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 8, "B45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 12, "B45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 12.1, "B40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 16, "B40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 20, "B40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 20.1, "B35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 24, "B35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 28, "B35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 28.1, "B30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 32, "B30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 36, "B30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 36.1, "B25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 40, "B25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 44, "B25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 44.1, "B20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 48, "B20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 52, "B20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 52.1, "B15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 56, "B15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 60, "B15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 60.1, "B10"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 64, "B10"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 68, "B5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 68.1, "B5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 72, "B5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 76, "B5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 76.1, "B0"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 80, "B0"},
		{MakeNCAAFootball(8, 5).SideToSideLines, 9001, "B0"},

		// Side to Side line tests
		{MakeNCAAFootball(8, 5).SideToSideLines, -4.1, "A45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -8, "A45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -12, "A45"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -12.1, "A40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -16, "A40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -20, "A40"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -20.1, "A35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -24, "A35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -28, "A35"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -28.1, "A30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -32, "A30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -36, "A30"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -36.1, "A25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -40, "A25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -44, "A25"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -44.1, "A20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -48, "A20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -52, "A20"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -52.1, "A15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -56, "A15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -60, "A15"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -60.1, "A10"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -64, "A10"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -68, "A5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -68.1, "A5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -72, "A5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -76, "A5"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -76.1, "A0"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -80, "A0"},
		{MakeNCAAFootball(8, 5).SideToSideLines, -9001, "A0"},

		// Front to Back line tests
		{MakeNCAAFootball(8, 5).FrontToBackLines, -9001, "FSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 0, "FSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 16, "FSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 16.1, "FH"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 32, "FH"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 42, "FH"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 42.1, "BH"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 52, "BH"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 68, "BSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 68.1, "BSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 84, "BSL"},
		{MakeNCAAFootball(8, 5).FrontToBackLines, 9001, "BSL"},

		// Front to Back line tests
		{MakeHSFootball(8, 5).FrontToBackLines, -9001, "FSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 0, "FSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 14, "FSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 14.1, "FH"},
		{MakeHSFootball(8, 5).FrontToBackLines, 28, "FH"},
		{MakeHSFootball(8, 5).FrontToBackLines, 42, "FH"},
		{MakeHSFootball(8, 5).FrontToBackLines, 42.1, "BH"},
		{MakeHSFootball(8, 5).FrontToBackLines, 56, "BH"},
		{MakeHSFootball(8, 5).FrontToBackLines, 70, "BSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 70.1, "BSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 84, "BSL"},
		{MakeHSFootball(8, 5).FrontToBackLines, 9001, "BSL"},
	}

	for _, test := range closestSideToSideLineTests {
		v := closestLine(test.lines, test.steps)
		if v != test.expected {
			t.Errorf("For lines %v with steps %f, expected %s, got %s", test.lines,
				test.steps, test.expected, v)
		}
	}
}

func TestLineSide(t *testing.T) {
	f := MakeNCAAFootball(8, 5)
	var lineSideTests = []struct {
		line      string
		expected  FieldSide
		shouldErr bool
	}{
		{"A30", Left, false},
		{"A50", Left, false},
		{"B50", Left, false},
		{"B30", Right, false},

		{"A51", Left, true},
		{"C50", Left, true},
	}

	for _, test := range lineSideTests {
		v, err := f.LineSide(test.line)
		if (err != nil) == test.shouldErr {
			if v != test.expected {
				t.Errorf("For lines %s, expected %v, got %v", test.line, test.expected,
					v)
			}
		} else {
			t.Errorf("Line %s error: %v, expected %v", test.line, err != nil, test.shouldErr)
		}
	}
}
