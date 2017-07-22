package models

import (
	"fmt"
	"strconv"
	"strings"
)

// parseXDot converts a string to an Point coordinate.
// steps: Any integer or float input is acceptable
// insideOrOutside: Inside is "I", Outside is "O"
// yardline: Must be a valid yardline for the given FieldLayout
func parseXDot(s string, f *FieldLayout) (float64, error) {
	list := strings.Split(s, " ")

	steps, err := strconv.ParseFloat(list[0], 64)
	if err != nil {
		return 0, err
	}

	line := list[2]
	var x float64
	var ok bool
	if x, ok = f.SideToSideLines[line]; !ok {
		return 0, fmt.Errorf("Illegal line: %s", line)
	}

	var sign float64
	switch line[0] {
	case 'A':
		sign = -1
	case 'B':
		sign = 1
	}

	var inOutSign float64
	if x != 0 {
		switch list[1] {
		case "I":
			inOutSign = -1 * sign
		case "O":
			inOutSign = sign
		default:
			return 0, fmt.Errorf("Illegal side: %s", list[1])
		}
	} else {
		inOutSign = sign
	}

	return x + inOutSign*steps, nil
}

// parseYDot converts a string to a yDot.
// steps: Any integer or float input is acceptable
// FrontOrBehind: Front is "F", Behind is "B"
// line: Must be a valid line for the given FieldLayout
func parseYDot(s string, f *FieldLayout) (float64, error) {
	list := strings.Split(s, " ")

	steps, err := strconv.ParseFloat(list[0], 64)
	if err != nil {
		return 0, err
	}

	FrontOrBehind := list[1]
	var sign float64 = 1
	if FrontOrBehind == "F" {
		sign = -1
	} else if FrontOrBehind != "B" {
		return 0, fmt.Errorf("Illegal: %s", FrontOrBehind)
	}

	line := list[2]
	var y float64
	var ok bool
	if y, ok = f.FrontToBackLines[line]; !ok {
		return 0, fmt.Errorf("Illegal line %s", line)
	}

	return sign*steps + y, nil
}
