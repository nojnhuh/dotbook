package models

import (
	"fmt"
	"math"
)

var (
	roundStepsToNearest    = 0.25
	roundStepSizeToNearest = 0.1
	countSubdivisions      = []string{"", "e", "&", "a"}
)

func (db *Dotbook) String() string {
	s := fmt.Sprintf("<Dotbook name: %s>\n", db.Name)
	for _, d := range db.Dots {
		s += fmt.Sprintln(d.String(db.Field))
		if d.PrevDot != nil {
			s += fmt.Sprintf("Step size: %.1f to 5\n",
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
			mid, err := d.DotOnCount(d.MoveCounts / 2)
			if err != nil {
				s += "Error calculating midset\n"
			} else {
				s += fmt.Sprintln(mid.CoordString(db.Field))
			}
			if d.BodyCenter {
				s += fmt.Sprintf("Foot dot:\n")
				foot := d.BodyToFootDot()
				s += fmt.Sprintln(foot.CoordString(db.Field))
			}
		}
		s += "\n"
	}
	return s
}

func (d *Dot) String(f *FieldLayout) string {
	s := fmt.Sprintf("Set %s\n%.0f counts\nhold %.0f\n%s", d.Name, d.MoveCounts,
		d.HoldCounts, d.CoordString(f))
	if d.BodyCenter {
		s += "\nBody-center"
	}
	return s
}

// CoordString constructs a string from a Point in dot notation.
func (d *Dot) CoordString(f *FieldLayout) string {
	return fmt.Sprintf("%s\n%s", d.NumToXDotString(f),
		d.NumToYDotString(f))
}

func (f *FieldLayout) String() string {
	return fmt.Sprintf("<Field w: %.3g, h: %.2g>", f.Width, f.Height)
}

// NumToXDotString converts a Cartesian coordinate back to an xDot string using
// the same coordinate system as dotToNum.
func (d *Dot) NumToXDotString(f *FieldLayout) string {
	n := d.Point.X
	line := closestLine(f.SideToSideLines, n)

	var insideOrOutside string
	var steps float64
	if math.Abs(n) > f.Width/2 {
		insideOrOutside = "O"
		steps = math.Abs(n) - f.Width/2
	} else {
		if math.Mod(math.Abs(n), f.StepsBetweenLines) >= f.StepsBetweenLines/2 {
			insideOrOutside = "I"
			steps = f.StepsBetweenLines - math.Mod(math.Abs(n),
				f.StepsBetweenLines)
		} else {
			insideOrOutside = "O"
			steps = math.Mod(math.Abs(n), f.StepsBetweenLines/2)
		}
	}

	// log.Printf("%.3g is %.3g %s %s", n, steps, insideOrOutside, line)
	return fmt.Sprintf("%.3g %s %s", roundToNearest(steps, roundStepsToNearest),
		insideOrOutside, line)
}

// NumToYDotString converts a Cartesian coordinate back to a yDot string using
// the same coordinate system as dotToNum.
func (d *Dot) NumToYDotString(f *FieldLayout) string {
	n := d.Point.Y
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

// CrossCountStrings returns an array of strings representing the crossing counts
// for the given dot
func (d *Dot) CrossCountStrings(f *FieldLayout) []string {
	s := []string{}
	for _, cross := range d.CrossingCounts(f) {
		lastCount := 0.
		if d.HoldCounts != 0 {
			lastCount = d.HoldCounts
		} else if d.PrevDot.MoveCounts != 0 {
			lastCount = d.PrevDot.MoveCounts
		} else if d.PrevDot.HoldCounts != 0 {
			lastCount = d.PrevDot.HoldCounts
		}
		s = append(s, fmt.Sprintf("Cross %s on count %s\n", cross.Line,
			formatCrossCount(cross.Count, lastCount)))
	}
	return s
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
