package models

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// FieldSide represents one half of a field, either left or right, from audience
// perspective.
type FieldSide int

// Valid values of type FieldSide
const (
	Left = iota
	Right
)

// FieldLayout Type Definition
// Defines the coordinate grid for a football field or indoor floor.
// Units of the grid are "steps" in the "x to 5" format.
// Usually, an 8 to 5 (22.5 in) step is used, but anything can be used here.
//
// Members:
// StepsBetweenLines: Scale used for the size of one step. A value of
// 		8 = 22.5in. step.
// Width: Total field width in steps.
// SideToSideLines: Map of vertical line names to steps from middle.
// Height: Total field height in steps.
// FrontToBackLines: Map of horizonal line names to steps from front.
type FieldLayout struct {
	StepsBetweenLines float64
	Width             float64
	SideToSideLines   map[string]float64
	Height            float64
	FrontToBackLines  map[string]float64
}

// makeFootball creates a FieldLayout for a standard football field.
// 300 feet side-to-side, 160 feet front-to-back.
// Yardlines are 5 yards apart.
// Hash calculations are delegated to other methods.
func makeFootball(stepsBetweenLines float64, yardsBetweenLines float64,
	hashes map[string]float64) *FieldLayout {
	midLine := 50.
	width := midLine * 2 / yardsBetweenLines * stepsBetweenLines

	sideToSideLines := make(map[string]float64)
	for line := int(midLine); line >= 0; line -= int(yardsBetweenLines) {
		steps := (midLine - float64(line)) * stepsBetweenLines / yardsBetweenLines
		sideToSideLines["A"+strconv.Itoa(line)] = -1 * steps
		sideToSideLines["B"+strconv.Itoa(line)] = steps
	}

	height := 10.5 * stepsBetweenLines / yardsBetweenLines * 5

	return &FieldLayout{
		StepsBetweenLines: stepsBetweenLines,
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
func MakeNCAAFootball(stepsBetweenLines, yardsBetweenLines float64) *FieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  4 * stepsBetweenLines / yardsBetweenLines * 5,
		"BH":  6.5 * stepsBetweenLines / yardsBetweenLines * 5,
		"BSL": 10.5 * stepsBetweenLines / yardsBetweenLines * 5,
	}

	return makeFootball(stepsBetweenLines, yardsBetweenLines, hashes)
}

// MakeHSFootball creates a standard high school footabll field.
// Front hash: 53'4" from front sideline.
// Back hash: 106'8" from front sideline.
//
// Since not every line is an even number of steps from the front sideline:
// "FH" in map is actually 52.5 feet from the front sideline.
// "BH" in map is actually 105 feet from front sideline.
// "BSL" in map is actually 157.5 feet from front sideline.
func MakeHSFootball(stepsBetweenLines, yardsBetweenLines float64) *FieldLayout {
	hashes := map[string]float64{
		"FSL": 0,
		"FH":  3.5 * stepsBetweenLines / yardsBetweenLines * 5,
		"BH":  7 * stepsBetweenLines / yardsBetweenLines * 5,
		"BSL": 10.5 * stepsBetweenLines / yardsBetweenLines * 5,
	}

	return makeFootball(stepsBetweenLines, 5, hashes)
}

// LineSide determines whice side of the field a given line is on.
func (f *FieldLayout) LineSide(line string) (FieldSide, error) {
	var x float64
	var ok bool
	if x, ok = f.SideToSideLines[line]; !ok {
		return Left, fmt.Errorf("Illegal line: %s", line)
	}
	if x > 0 {
		return Right, nil
	}
	return Left, nil
}

// closestLine finds the name of the line closest to the given coordindate.
func closestLine(m map[string]float64, steps float64) string {
	minDistance := math.MaxFloat64
	var minLine string
	var distance float64

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		distance = math.Abs(m[k] - steps)
		if distance <= minDistance {
			minDistance = distance
			minLine = k
		}
	}

	return minLine
}
