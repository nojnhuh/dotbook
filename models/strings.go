package models

import (
	"fmt"
	"math"
)

var (
	roundStepsToNearest    = 0.25
	roundStepSizeToNearest = 0.1
	countSubdivisions      = []string{"", "e", "+", "a"}
)

func (db *Dotbook) String() string {
	s := fmt.Sprintf("<Dotbook name: %s>\n", db.Name)
	for _, d := range db.Dots {
		s += fmt.Sprintln(d.String(db.Field))
		if d.PrevDot != nil {
			s += fmt.Sprintf("Step size = %.1f to 5\n",
				roundToNearest(d.StepSize(db.Field), roundStepSizeToNearest))
		}
		for _, cross := range d.CrossingCounts(db.Field) {
			lastCount := 0.
			if d.HoldCounts != 0 {
				lastCount = d.HoldCounts
			} else if d.PrevDot.MoveCounts != 0 {
				lastCount = d.PrevDot.MoveCounts
			} else if d.PrevDot.HoldCounts != 0 {
				lastCount = d.PrevDot.HoldCounts
			}
			s += fmt.Sprintf("Cross %s on count %s\n", cross.Line,
				formatCrossCount(cross.Count, lastCount))
		}
		if d.PrevDot != nil {
			s += fmt.Sprintf("Midset: %.0f counts\n", d.MoveCounts/2)
			mid := d.DotOnCount(d.MoveCounts / 2)
			s += fmt.Sprintln(mid.coordString(db.Field))
			if d.BodyCenter {
				s += fmt.Sprintf("Foot dot:\n")
				foot := d.bodyToFootDot()
				s += fmt.Sprintln(foot.coordString(db.Field))
			}
		}
		s += "\n"
	}
	return s
}

func (d *Dot) String(f *fieldLayout) string {
	s := fmt.Sprintf("Set %s\n%.0f counts\nhold %.0f\n%s", d.Name, d.MoveCounts,
		d.HoldCounts, d.coordString(f))
	if d.BodyCenter {
		s += "\nBody-center"
	}
	return s
}

// coordString constructs a string from a Point in dot notation.
func (d *Dot) coordString(f *fieldLayout) string {
	x := d.Point.X
	y := d.Point.Y
	return fmt.Sprintf("%s\n%s", numToXDotString(x, f),
		numToYDotString(y, f))
}

func (f *fieldLayout) String() string {
	return fmt.Sprintf("<Field w: %.3g, h: %.2g>", f.Width, f.Height)
}

// Converts a Cartesian coordinate back to an xDot string using the same
// coordinate system as dotToNum.
func numToXDotString(n float64, f *fieldLayout) string {

	line := closestLine(f.SideToSideLines, n)

	var insideOrOutside string
	var steps float64
	if math.Abs(n) > f.Width/2 {
		insideOrOutside = "O"
		steps = math.Abs(n) - f.Width/2
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

	// log.Printf("%.3g is %.3g %s %s", n, steps, insideOrOutside, line)
	return fmt.Sprintf("%.3g %s %s", roundToNearest(steps, roundStepsToNearest),
		insideOrOutside, line)
}

// Converts a Cartesian coordinate back to a yDot string using the same
// coordinate system as dotToNum.
func numToYDotString(n float64, f *fieldLayout) string {
	line := closestLine(f.FrontToBackLines, n)
	steps := n - f.FrontToBackLines[line]

	var frontOrBehind string
	if steps < 0 {
		frontOrBehind = "F"
		steps = math.Abs(steps)
	} else {
		frontOrBehind = "B"
	}

	return fmt.Sprintf("%.4g %s%s", roundToNearest(steps, roundStepsToNearest),
		frontOrBehind, line)
}

// roundToNearest rounds a float to the nearest fraction defined by frac
func roundToNearest(n, frac float64) float64 {
	return float64(int64(n/frac+0.5)) * frac
}

func formatCrossCount(count, lastCount float64) string {
	if int(count) == 0 {
		count += lastCount
	}
	return fmt.Sprintf("%d%s", int(count),
		countSubdivisions[int(float64(len(countSubdivisions))*math.Mod(count, 1))])
}
