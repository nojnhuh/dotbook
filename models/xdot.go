package models

import (
	"log"
	"strconv"
	"strings"
)

// // Type definition for InsideOrOutside field in XDot
// type InsideOrOutside int
// type Side int
//
// // for InsideOrOutside
// const (
// 	Inside = iota
// 	Outside
// )
//
// // For Side
// const (
// 	Left  = 1
// 	Right = 2
// )
//
// // An XDot represents a side-to-side coordinate.
// // Steps: number of steps from the closest line of reference.
// // InsideOrOutside: which side of the line the coordinate is off of.
// // 		Inside is toward the center, outside is away from the center.
// // Side: Which side of the field the coordinate is on, Left or Right.
// // 		(From audience perspective)
// // Line: The closest side-to-side line of reference, like a yardline.
// type XDot struct {
// 	Steps           float64
// 	InsideOrOutside InsideOrOutside
// 	Side            Side
// 	Line            string
// }

// Translates an xDot into a Cartesian coordinate, with the unit being steps.
// Mid-field is 0, left is negative, right is positive.
// func (d *XDot) dotToNum(f *fieldLayout) float64 {
// 	var sign float64
// 	if d.Side == Left {
// 		sign = -1
// 	} else {
// 		sign = 1
// 	}
//
// 	x := f.SideToSideLines[d.Line]
// 	var in_out_sign float64
// 	if x != 0 {
// 		if d.InsideOrOutside == Inside {
// 			in_out_sign = -1 * sign
// 		} else {
// 			in_out_sign = sign
// 		}
// 	} else {
// 		in_out_sign = sign
// 	}
//
// 	return sign*x + in_out_sign*d.Steps
// }

// Converts a string to an xDot.
// steps: Any integer or float input is acceptable
// insideOrOutside: Inside is "I", Outside is "O"
// side: Left is "1", Right is "2"
// yardline: Must be a valid yardline for the given fieldLayout
func parseXDot(s string, f *fieldLayout) float64 {
	list := strings.Split(s, " ")

	steps, err := strconv.ParseFloat(list[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	var sign float64
	switch list[2] {
	case "1":
		sign = -1
	case "2":
		sign = 1
	default:
		log.Fatal("Illegal field side ", list[2])
	}

	line := list[3]
	var x float64
	var ok bool
	if x, ok = f.SideToSideLines[line]; !ok {
		log.Fatalf("Illegal line %s", line)
	}

	var in_out_sign float64
	if x != 0 {
		switch list[1] {
		case "I":
			in_out_sign = -1 * sign
		case "O":
			in_out_sign = sign
		default:
			log.Fatal("Illegal side ", list[1])
		}
	} else {
		in_out_sign = sign
	}

	return sign*x + in_out_sign*steps
}
