// Stores the string definitions for each object
package models

import (
	"fmt"
	"math"
)

func (db *Dotbook) String() string {
	s := fmt.Sprintf("<Dotbook name: %s>\n", db.Name)
	for _, d := range db.Dots {
		s += fmt.Sprintln(d.String(db.Field))
		if d.PrevDot != nil {
			s += fmt.Sprintf("Step size = %.1f to 5\n", db.StepSize(d))
		}
		for _, cross := range db.CrossingCounts(d) {
			s += fmt.Sprintf("Cross side %s %s on count %.2f\n", cross.Side, cross.Line, cross.Count)
		}
		if d.PrevDot != nil {
			s += fmt.Sprintf("Midset: %.0f counts\n", d.MoveCounts/2)
			mid := db.DotOnCount(d, d.MoveCounts/2)
			s += fmt.Sprintln(mid.coordString(db.Field))
			if d.BodyCenter {
				s += fmt.Sprintf("Foot dot:\n")
				foot := db.bodyToFootDot(d)
				s += fmt.Sprintln(foot.coordString(db.Field))
			}
		}
		s += "\n"
	}
	return s
}

func (d *Dot) String(f *fieldLayout) string {
	s := fmt.Sprintf("Set %s\n%.0f counts\nhold %.0f\n%s", d.Name, d.MoveCounts, d.HoldCounts, d.coordString(f))
	if d.BodyCenter {
		s += "\nBody-center"
	}
	return s
}

func (d *Dot) coordString(f *fieldLayout) string {
	x := d.Point.X
	y := d.Point.Y
	return fmt.Sprintf("%s\n%s", numToXDotString(x, f),
		numToYDotString(y, f))
}

func (f *fieldLayout) String() string {
	return fmt.Sprintf("<Field w: %.3g, h: %.2g>", f.width, f.height)
}

// Converts a Cartesian coordinate back to an xDot string using the same
// coordinate system as dotToNum.
func numToXDotString(n float64, f *fieldLayout) string {
	var side string
	if n > 0 {
		side = "2"
	} else {
		side = "1"
	}

	line := closestLine(f.SideToSideLines, n, true)

	var insideOrOutside string
	var steps float64
	if math.Abs(n) > f.width/2 {
		insideOrOutside = "O"
		steps = math.Abs(n) - f.width/2
	} else {
		if math.Mod(math.Abs(n), f.StepsPerFiveYards) >= f.StepsPerFiveYards/2 {
			insideOrOutside = "I"
			steps = f.StepsPerFiveYards - math.Mod(math.Abs(n),
				f.StepsPerFiveYards)
		} else {
			insideOrOutside = "O"
			steps = math.Mod(math.Abs(n), f.StepsPerFiveYards/2)
		}
	}
	// fmt.Printf("%.2f is %s\n", n, fmt.Sprintf("%.3g %s %s %s", steps, insideOrOutside, side, line))

	return fmt.Sprintf("%.3g %s %s %s", steps, insideOrOutside, side, line)
}

// Converts a Cartesian coordinate back to a yDot string using the same
// coordinate system as dotToNum.
func numToYDotString(n float64, f *fieldLayout) string {
	line := closestLine(f.FrontToBackLines, n, false)
	steps := n - f.FrontToBackLines[line]

	var frontOrBehind string
	if steps < 0 {
		frontOrBehind = "F"
		steps = math.Abs(steps)
	} else {
		frontOrBehind = "B"
	}

	return fmt.Sprintf("%.4g %s%s", steps, frontOrBehind, line)
}
