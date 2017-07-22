package models

import (
	"fmt"

	"github.com/nojnhuh/dotbook/models/dotmath"
)

// Dot represents one full coordinate
// Name: Set number, like "7" or "123A"
// MoveCounts: Number of counts needed to get to this dot
// HoldCounts: Number of counts held at this dot
// Point: A Cartesian coordinate representing the dot's location on a field
// BodyCenter: indicates whether at this dot should be under the performer's
// 		foot or it should be between the performer's feet on the last count
// PrevDot: A reference to the dot before this one
type Dot struct {
	Name       string
	MoveCounts float64
	HoldCounts float64
	Point      dotmath.Point
	BodyCenter bool
	PrevDot    *Dot
}

// equals tests two Dots for equality
func (d *Dot) equals(d2 *Dot) bool {
	if d == nil && d2 == nil {
		return true
	}
	if d == nil || d2 == nil {
		return false
	}
	return d.Name == d2.Name &&
		d.MoveCounts == d2.MoveCounts &&
		d.HoldCounts == d2.HoldCounts &&
		dotmath.Equal(d.Point, d2.Point) &&
		d.BodyCenter == d2.BodyCenter //&&
	// d.PrevDot.equals(d2.PrevDot)
}

// bodyToFootDot converts a foot dot to a body-center dot.
func (d *Dot) bodyToFootDot() *Dot {
	if d.BodyCenter {
		dot := *d
		if dot.PrevDot == nil {
			return &dot
		}
		dot.Point = dotmath.BodyToFootDot(d.Point, d.PrevDot.Point, d.MoveCounts)
		dot.BodyCenter = false
		return &dot
	}
	return d
}

// DotOnCount calculates an intermediate dot getting to the given dot based on
// the count parameter. Passing d.moveCounts/2 as count finds the midset.
func (d *Dot) DotOnCount(count float64) (*Dot, error) {
	if count > d.MoveCounts || count < 0 {
		return nil, fmt.Errorf("DotOnCount: Count %.0f invalid", count)
	}
	if d.PrevDot == nil || d.MoveCounts == 0 {
		return d, nil
	}
	t := count / d.MoveCounts
	dFoot := d.bodyToFootDot()
	prevFoot := d.PrevDot.bodyToFootDot()
	midPoint := dotmath.ScalarMult(dotmath.AddPoints(dFoot.Point, prevFoot.Point), t)
	mid := &Dot{d.Name, count, 0, midPoint, false, d}
	return mid, nil
}

// Distance calculates the number of steps between this dot and the previous one.
func (d *Dot) Distance() float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := d.bodyToFootDot()
	prevFoot := d.PrevDot.bodyToFootDot()
	return dotmath.Distance(dFoot.Point, prevFoot.Point)
}

// StepSize calculates the step size needed to get to the given dot in the form
// of "x to 5" where x steps are needed to go 5 yards.
func (d *Dot) StepSize(f *FieldLayout) float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := d.bodyToFootDot()
	prevFoot := d.PrevDot.bodyToFootDot()
	segSize := dotmath.SegmentSize(dFoot.Point, prevFoot.Point, d.MoveCounts)
	if segSize != 0 {
		return f.StepsBetweenLines / segSize
	}
	return 0
}

// CrossingCounts calculates the crossing counts getting to the given dot
func (d *Dot) CrossingCounts(f *FieldLayout) []dotmath.CrossCount {
	if d.PrevDot == nil {
		return []dotmath.CrossCount{}
	}
	crosses := dotmath.CrossingCounts(d.Point, d.PrevDot.Point, f.SideToSideLines)
	for i, cross := range crosses {
		crosses[i].Count = cross.Count * d.MoveCounts
	}
	return crosses
}
