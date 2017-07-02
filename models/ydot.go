package models

import (
	"log"
	"strconv"
	"strings"
)

// type FrontOrBehind int
//
// const (
// 	Front = iota
// 	Behind
// )
//
// // A YDot represents a front-to-back coordinate.
// // Steps: number of steps from the nearest line of reference.
// // FrontOrBehind: which side of the reference line the coordinate is on.
// // 		Front is toward the audience, Behind is away from the audience.
// // Line: The closest front-to-back line of reference, like a hash or sideline.
// type YDot struct {
// 	Steps         float64
// 	FrontOrBehind FrontOrBehind
// 	Line          string
// }

// // Translates a yDot into a Cartesian coordinate, with the unit being steps.
// // The front of the field is 0, behind is positive, in front is negative
// func (d *YDot) dotToNum(f *fieldLayout) float64 {
// 	var sign float64
// 	if d.FrontOrBehind == Front {
// 		sign = -1
// 	} else {
// 		sign = 1
// 	}
//
// 	return sign*d.Steps + f.FrontToBackLines[d.Line]
// }

// Converts a string to a yDot.
// steps: Any integer or float input is acceptable
// FrontOrBehind: Front is "F", Behind is "B"
// line: Must be a valid line for the given fieldLayout
func parseYDot(s string, f *fieldLayout) float64 {
	list := strings.Split(s, " ")

	steps, err := strconv.ParseFloat(list[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	FrontOrBehind := list[1]
	var sign float64 = 1
	if FrontOrBehind == "F" {
		sign = -1
	} else if FrontOrBehind != "B" {
		log.Fatalf("Illegal: %s", FrontOrBehind)
	}

	line := list[2]
	var y float64
	var ok bool
	if y, ok = f.FrontToBackLines[line]; !ok {
		log.Fatalf("Illegal line %s", line)
	}

	return sign*steps + y

}
