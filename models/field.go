package models

import (
	"math"
	"sort"
	"strconv"
)

type FieldSide int

const (
	Left = iota
	Right
)

// fieldLayout Type Definition
// Defines the coordinate grid for a football field or indoor floor.
// Units of the grid are "steps" in the "x to 5" format.
// Usually, an 8 to 5 (22.5 in) step is used, but anything can be used here.
//
// Members:
// StepsPerFiveYards: Scale used for the size of one step. A value of
// 		8 = 22.5in. step.
// Width: Total field width in steps.
// SideToSideLines: Map of vertical line names to steps from middle.
// Height: Total field height in steps.
// FrontToBackLines: Map of horizonal line names to steps from front.
type fieldLayout struct {
	StepsPerFiveYards float64
	Width             float64
	SideToSideLines   map[string]float64
	Height            float64
	FrontToBackLines  map[string]float64
}

// makeFootball creates a fieldLayout for a standard football field.
// 300 feet side-to-side, 160 feet front-to-back.
// Yardlines are 5 yards apart.
// Hash calculations are delegated to other methods.
func makeFootball(stepsPerFiveYards float64, yardsBetweenLines float64,
	hashes map[string]float64) *fieldLayout {
	midLine := 50.
	width := midLine / 2.5 * stepsPerFiveYards

	sideToSideLines := make(map[string]float64)
	sideToSideLines["50"] = 0
	for line := int(midLine) - 5; line >= 0; line -= 5 {
		steps := (midLine - float64(line)) * stepsPerFiveYards / yardsBetweenLines
		sideToSideLines["A"+strconv.Itoa(line)] = -1 * steps
		sideToSideLines["B"+strconv.Itoa(line)] = steps
	}

	height := 10.5 * stepsPerFiveYards

	return &fieldLayout{
		StepsPerFiveYards: stepsPerFiveYards,
		Width:             width,
		SideToSideLines:   sideToSideLines,
		Height:            height,
		FrontToBackLines:  hashes,
	}
}

// MakeNCAAFootball creates a standard NCAA footabll field (DCI, TX high school
// standard)
// Front hash: 60 feet from front sideline
// Back hash: 100 feet from front sideline
//
// Since not every line is an even number of steps from the front sideline:
// "BH" in map is actually 97.5 feet from front sideline.
// "BSL" in map is actually 157.5 feet from front sideline.
func MakeNCAAFootball(stepsPerFiveYards float64) *fieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  4 * stepsPerFiveYards,
		"BH":  6.5 * stepsPerFiveYards,
		"BSL": 10.5 * stepsPerFiveYards,
	}

	return makeFootball(stepsPerFiveYards, 5, hashes)
}

// MakeHSFootball creates a standard high school footabll field.
// Front hash: 53'4" from front sideline.
// Back hash: 106'8" from front sideline.
//
// Since not every line is an even number of steps from the front sideline:
// "FH" in map is actually 52.5 feet from the front sideline.
// "BH" in map is actually 105 feet from front sideline.
// "BSL" in map is actually 157.5 feet from front sideline.
func MakeHSFootball(stepsPerFiveYards float64) *fieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  3.5 * stepsPerFiveYards,
		"BH":  7 * stepsPerFiveYards,
		"BSL": 10.5 * stepsPerFiveYards,
	}

	return makeFootball(stepsPerFiveYards, 5, hashes)
}

// LineSide determines whice side of the field a given line is on.
func LineSide(line string) FieldSide {
	switch line[0] {
	case 'A':
		return Left
	case 'B':
		return Right
	default:
		panic("Illegal line")
	}
}

// closestLine finds the name of the line closest to the given coordindate.
func closestLine(m map[string]float64, steps float64) string {
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
