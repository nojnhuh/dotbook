package models

import (
	"math"
	"sort"
	"strconv"
)

// fieldLayout Type Definition
// Defines the coordinate grid for a football field or indoor floor.
// Units of the grid are "steps" in the "x to 5" format.
// Usually, an 8 to 5 (22.5 in) step is used, but anything can be used here.
type fieldLayout struct {
	// Scale used for the size of one step. A value of 8 = 22.5 in. step.
	StepsPerFiveYards float64

	// Width in steps
	width float64

	// Map of vertical line names to steps from middle
	// Assumes the list represents a symmetrical field, so this map only stores
	// half of the actual lines
	SideToSideLines map[string]float64

	// Height in steps
	height float64

	// Map of horizonal line names to steps from front
	FrontToBackLines map[string]float64
}

// Creates a fieldLayout for a standard football field
// 300 feet side-to-side, 160 feet front-to-back
// Yardlines are 5 yards apart
// Hash calculations are delegated to other methods
func makeFootball(StepsPerFiveYards float64, yardsBetweenLines float64,
	hashes map[string]float64) *fieldLayout {
	midLine := 50.
	width := midLine / 2.5 * StepsPerFiveYards

	// sideToSideLines := make([]lineNumPair, int(midLine)*2/int(yardsBetweenLines)+1)
	SideToSideLines := make(map[string]float64)
	for line := int(midLine); line >= 0; line -= 5 {
		steps := (midLine - float64(line)) * StepsPerFiveYards / yardsBetweenLines
		SideToSideLines[strconv.Itoa(line)] = steps
	}

	height := 10.5 * StepsPerFiveYards

	return &fieldLayout{
		StepsPerFiveYards: StepsPerFiveYards,
		width:             width,
		SideToSideLines:   SideToSideLines,
		height:            height,
		FrontToBackLines:  hashes,
	}

}

// Represents a standard NCAA footabll field (DCI, TX high school standard)
// Front hash: 60 feet from front sideline
// Back hash: 100 feet from front sideline
//
// Since not every line is an even number of steps from the front sideline,
// "BH" in map is actually 97.5 feet from front sideline
// "BSL" in map is actually 157.5 feet from front sideline
func MakeNCAAFootball(StepsPerFiveYards float64) *fieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  4 * StepsPerFiveYards,
		"BH":  6.5 * StepsPerFiveYards,
		"BSL": 10.5 * StepsPerFiveYards,
	}

	return makeFootball(StepsPerFiveYards, 5, hashes)
}

// Represents a standard high school footabll field.
// Front hash: 53' 4" from front sideline
// Back hash: 106' 8" from front sideline
//
// Since not every line is an even number of steps from the front sideline,
// "FH" in map is actually 52.5 feet from the front sideline
// "BH" in map is actually 105 feet from front sideline
// "BSL" in map is actually 157.5 feet from front sideline
func MakeHSFootball(StepsPerFiveYards float64) *fieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  3.5 * StepsPerFiveYards,
		"BH":  7 * StepsPerFiveYards,
		"BSL": 10.5 * StepsPerFiveYards,
	}

	return makeFootball(StepsPerFiveYards, 5, hashes)
}

// Given a coordinate a number of steps from the main referece find the name of
// the line closest to that coordindate
func closestLine(m map[string]float64, steps float64, symmetrical bool) string {
	if symmetrical {
		steps = math.Abs(steps)
	}
	min_distance := math.MaxFloat64
	var min_line string
	var distance float64

	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		distance = math.Abs(m[k] - steps)
		if distance <= min_distance {
			min_distance = distance
			min_line = k
		}
	}

	return min_line
}
