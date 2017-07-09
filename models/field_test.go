package models

import (
	"testing"
)

type closestLine_test struct {
	lines    map[string]float64
	steps    float64
	expected string
}

var closestSideToSideLine_tests = []closestLine_test{
	// Side to Side line tests
	{MakeNCAAFootball(8).SideToSideLines, 0, "50"},
	{MakeNCAAFootball(8).SideToSideLines, 4, "B45"},
	{MakeNCAAFootball(8).SideToSideLines, 4.1, "B45"},
	{MakeNCAAFootball(8).SideToSideLines, 8, "B45"},
	{MakeNCAAFootball(8).SideToSideLines, 12, "B45"},
	{MakeNCAAFootball(8).SideToSideLines, 12.1, "B40"},
	{MakeNCAAFootball(8).SideToSideLines, 16, "B40"},
	{MakeNCAAFootball(8).SideToSideLines, 20, "B40"},
	{MakeNCAAFootball(8).SideToSideLines, 20.1, "B35"},
	{MakeNCAAFootball(8).SideToSideLines, 24, "B35"},
	{MakeNCAAFootball(8).SideToSideLines, 28, "B35"},
	{MakeNCAAFootball(8).SideToSideLines, 28.1, "B30"},
	{MakeNCAAFootball(8).SideToSideLines, 32, "B30"},
	{MakeNCAAFootball(8).SideToSideLines, 36, "B30"},
	{MakeNCAAFootball(8).SideToSideLines, 36.1, "B25"},
	{MakeNCAAFootball(8).SideToSideLines, 40, "B25"},
	{MakeNCAAFootball(8).SideToSideLines, 44, "B25"},
	{MakeNCAAFootball(8).SideToSideLines, 44.1, "B20"},
	{MakeNCAAFootball(8).SideToSideLines, 48, "B20"},
	{MakeNCAAFootball(8).SideToSideLines, 52, "B20"},
	{MakeNCAAFootball(8).SideToSideLines, 52.1, "B15"},
	{MakeNCAAFootball(8).SideToSideLines, 56, "B15"},
	{MakeNCAAFootball(8).SideToSideLines, 60, "B15"},
	{MakeNCAAFootball(8).SideToSideLines, 60.1, "B10"},
	{MakeNCAAFootball(8).SideToSideLines, 64, "B10"},
	{MakeNCAAFootball(8).SideToSideLines, 68, "B5"},
	{MakeNCAAFootball(8).SideToSideLines, 68.1, "B5"},
	{MakeNCAAFootball(8).SideToSideLines, 72, "B5"},
	{MakeNCAAFootball(8).SideToSideLines, 76, "B5"},
	{MakeNCAAFootball(8).SideToSideLines, 76.1, "B0"},
	{MakeNCAAFootball(8).SideToSideLines, 80, "B0"},
	{MakeNCAAFootball(8).SideToSideLines, 9001, "B0"},

	// Side to Side line tests
	{MakeNCAAFootball(8).SideToSideLines, -4.1, "A45"},
	{MakeNCAAFootball(8).SideToSideLines, -8, "A45"},
	{MakeNCAAFootball(8).SideToSideLines, -12, "A45"},
	{MakeNCAAFootball(8).SideToSideLines, -12.1, "A40"},
	{MakeNCAAFootball(8).SideToSideLines, -16, "A40"},
	{MakeNCAAFootball(8).SideToSideLines, -20, "A40"},
	{MakeNCAAFootball(8).SideToSideLines, -20.1, "A35"},
	{MakeNCAAFootball(8).SideToSideLines, -24, "A35"},
	{MakeNCAAFootball(8).SideToSideLines, -28, "A35"},
	{MakeNCAAFootball(8).SideToSideLines, -28.1, "A30"},
	{MakeNCAAFootball(8).SideToSideLines, -32, "A30"},
	{MakeNCAAFootball(8).SideToSideLines, -36, "A30"},
	{MakeNCAAFootball(8).SideToSideLines, -36.1, "A25"},
	{MakeNCAAFootball(8).SideToSideLines, -40, "A25"},
	{MakeNCAAFootball(8).SideToSideLines, -44, "A25"},
	{MakeNCAAFootball(8).SideToSideLines, -44.1, "A20"},
	{MakeNCAAFootball(8).SideToSideLines, -48, "A20"},
	{MakeNCAAFootball(8).SideToSideLines, -52, "A20"},
	{MakeNCAAFootball(8).SideToSideLines, -52.1, "A15"},
	{MakeNCAAFootball(8).SideToSideLines, -56, "A15"},
	{MakeNCAAFootball(8).SideToSideLines, -60, "A15"},
	{MakeNCAAFootball(8).SideToSideLines, -60.1, "A10"},
	{MakeNCAAFootball(8).SideToSideLines, -64, "A10"},
	{MakeNCAAFootball(8).SideToSideLines, -68, "A5"},
	{MakeNCAAFootball(8).SideToSideLines, -68.1, "A5"},
	{MakeNCAAFootball(8).SideToSideLines, -72, "A5"},
	{MakeNCAAFootball(8).SideToSideLines, -76, "A5"},
	{MakeNCAAFootball(8).SideToSideLines, -76.1, "A0"},
	{MakeNCAAFootball(8).SideToSideLines, -80, "A0"},
	{MakeNCAAFootball(8).SideToSideLines, -9001, "A0"},

	// Front to Back line tests
	{MakeNCAAFootball(8).FrontToBackLines, -9001, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 0, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 16, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 16.1, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 32, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 42, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 42.1, "BH"},
	{MakeNCAAFootball(8).FrontToBackLines, 52, "BH"},
	{MakeNCAAFootball(8).FrontToBackLines, 68, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 68.1, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 84, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 9001, "BSL"},

	// Front to Back line tests
	{MakeHSFootball(8).FrontToBackLines, -9001, "FSL"},
	{MakeHSFootball(8).FrontToBackLines, 0, "FSL"},
	{MakeHSFootball(8).FrontToBackLines, 14, "FSL"},
	{MakeHSFootball(8).FrontToBackLines, 14.1, "FH"},
	{MakeHSFootball(8).FrontToBackLines, 28, "FH"},
	{MakeHSFootball(8).FrontToBackLines, 42, "FH"},
	{MakeHSFootball(8).FrontToBackLines, 42.1, "BH"},
	{MakeHSFootball(8).FrontToBackLines, 56, "BH"},
	{MakeHSFootball(8).FrontToBackLines, 70, "BSL"},
	{MakeHSFootball(8).FrontToBackLines, 70.1, "BSL"},
	{MakeHSFootball(8).FrontToBackLines, 84, "BSL"},
	{MakeHSFootball(8).FrontToBackLines, 9001, "BSL"},
}

func TestClosestLine(t *testing.T) {
	for _, test := range closestSideToSideLine_tests {
		v := closestLine(test.lines, test.steps)
		if v != test.expected {
			t.Error(
				"For field", test.lines,
				"with steps", test.steps,
				"expected", test.expected,
				"got", v,
			)
		}
	}
}
