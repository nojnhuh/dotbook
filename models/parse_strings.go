package models

import (
	"log"
	"strconv"
	"strings"
)

// Converts a string to an Point coordinate.
// steps: Any integer or float input is acceptable
// insideOrOutside: Inside is "I", Outside is "O"
// yardline: Must be a valid yardline for the given fieldLayout
func parseXDot(s string, f *fieldLayout) float64 {
	list := strings.Split(s, " ")

	steps, err := strconv.ParseFloat(list[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	line := list[2]
	var x float64
	var ok bool
	if x, ok = f.SideToSideLines[line]; !ok {
		log.Fatal("Illegal line:", line)
	}

	var sign float64
	switch LineSide(line) {
	case Left:
		sign = -1
	case Right:
		sign = 1
	}

	var in_out_sign float64
	if x != 0 {
		switch list[1] {
		case "I":
			in_out_sign = -1 * sign
		case "O":
			in_out_sign = sign
		default:
			log.Fatal("Illegal side:", list[1])
		}
	} else {
		in_out_sign = sign
	}

	return x + in_out_sign*steps
}

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
