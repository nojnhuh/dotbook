package models

import (
	"testing"
)

type closestLine_test struct {
	lines    map[string]float64
	steps    float64
	sym      bool
	expected string
}

var closestSideToSideLine_tests = []closestLine_test{
	// Side to Side line tests
	{MakeNCAAFootball(8).SideToSideLines, 0, true, "50"},
	{MakeNCAAFootball(8).SideToSideLines, 4, true, "50"},
	{MakeNCAAFootball(8).SideToSideLines, 4.1, true, "45"},
	{MakeNCAAFootball(8).SideToSideLines, 8, true, "45"},
	{MakeNCAAFootball(8).SideToSideLines, 12, true, "45"},
	{MakeNCAAFootball(8).SideToSideLines, 12.1, true, "40"},
	{MakeNCAAFootball(8).SideToSideLines, 16, true, "40"},
	{MakeNCAAFootball(8).SideToSideLines, 20, true, "40"},
	{MakeNCAAFootball(8).SideToSideLines, 20.1, true, "35"},
	{MakeNCAAFootball(8).SideToSideLines, 24, true, "35"},
	{MakeNCAAFootball(8).SideToSideLines, 28, true, "35"},
	{MakeNCAAFootball(8).SideToSideLines, 28.1, true, "30"},
	{MakeNCAAFootball(8).SideToSideLines, 32, true, "30"},
	{MakeNCAAFootball(8).SideToSideLines, 36, true, "30"},
	{MakeNCAAFootball(8).SideToSideLines, 36.1, true, "25"},
	{MakeNCAAFootball(8).SideToSideLines, 40, true, "25"},
	{MakeNCAAFootball(8).SideToSideLines, 44, true, "25"},
	{MakeNCAAFootball(8).SideToSideLines, 44.1, true, "20"},
	{MakeNCAAFootball(8).SideToSideLines, 48, true, "20"},
	{MakeNCAAFootball(8).SideToSideLines, 52, true, "20"},
	{MakeNCAAFootball(8).SideToSideLines, 52.1, true, "15"},
	{MakeNCAAFootball(8).SideToSideLines, 56, true, "15"},
	{MakeNCAAFootball(8).SideToSideLines, 60, true, "15"},
	{MakeNCAAFootball(8).SideToSideLines, 60.1, true, "10"},
	{MakeNCAAFootball(8).SideToSideLines, 64, true, "10"},
	{MakeNCAAFootball(8).SideToSideLines, 68, true, "5"},
	{MakeNCAAFootball(8).SideToSideLines, 68.1, true, "5"},
	{MakeNCAAFootball(8).SideToSideLines, 72, true, "5"},
	{MakeNCAAFootball(8).SideToSideLines, 76, true, "5"},
	{MakeNCAAFootball(8).SideToSideLines, 76.1, true, "0"},
	{MakeNCAAFootball(8).SideToSideLines, 80, true, "0"},
	{MakeNCAAFootball(8).SideToSideLines, 9001, true, "0"},

	// Front to Back line tests
	{MakeNCAAFootball(8).FrontToBackLines, -9001, false, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 0, false, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 16, false, "FSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 16.1, false, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 32, false, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 42, false, "FH"},
	{MakeNCAAFootball(8).FrontToBackLines, 42.1, false, "BH"},
	{MakeNCAAFootball(8).FrontToBackLines, 52, false, "BH"},
	{MakeNCAAFootball(8).FrontToBackLines, 68, false, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 68.1, false, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 84, false, "BSL"},
	{MakeNCAAFootball(8).FrontToBackLines, 9001, false, "BSL"},
}

func TestClosestLine(t *testing.T) {
	for _, test := range closestSideToSideLine_tests {
		v := closestLine(test.lines, test.steps, test.sym)
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
