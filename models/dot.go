package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/nojnhuh/dotbook/models/dotmath"
)

// Dot represents one full coordinate
type Dot struct {
	// Database ID
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	// Set number, like "7" or "123A"
	Name string `json:"name"`

	// Number of counts needed to get to this dot
	MoveCounts float64 `json:"moveCounts"`

	// Number of counts held at this dot
	HoldCounts float64 `json:"holdCounts"`

	// A Cartesian coordinate representing the dot's location on a field
	Point dotmath.Point `json:"point"`

	// indicates whether at this dot should be under the performer's foot or it
	// should be between the performer's feet on the last count
	BodyCenter bool `json:"bodyCenter"`

	// A reference to the dot before this one
	PrevDot *Dot `json:"-"`
}

// DotDetails contains the data for a Dot as well as derived fields
type DotDetails struct {
	// Set number, like "7" or "123A"
	Name string `json:"name"`

	// Number of counts needed to get to this dot
	MoveCounts float64 `json:"moveCounts"`

	// Number of counts held at this dot
	HoldCounts float64 `json:"holdCounts"`

	// A Cartesian coordinate representing the dot's location on a field
	Point dotmath.Point `json:"point"`

	// Indicates whether at this dot should be under the performer's foot or it
	// should be between the performer's feet on the last count
	BodyCenter bool `json:"bodyCenter"`

	// The dot halfway on the path to this dot
	Midset *Dot `json:"midset"`

	// The foot-on-dot equivalent for this dot
	FootDot *Dot `json:"footDot"`

	// The "_ to 5" step size required to get to this dot
	StepSize float64 `json:"stepSize"`

	// The lines crossed and counts on the path to this dot
	CrossingCounts dotmath.CrossCounts `json:"crossingCounts"`
}

// NewDot returns a reference to a new Dot object
func NewDot(name string, moveCounts, holdCounts float64,
	xdot, ydot float64, bodyCenter bool, prevDot *Dot) *Dot {
	return &Dot{
		ID:         bson.NewObjectId(),
		Name:       name,
		MoveCounts: moveCounts,
		HoldCounts: holdCounts,
		Point:      dotmath.NewPoint(xdot, ydot),
		BodyCenter: bodyCenter,
		PrevDot:    prevDot,
	}
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

// BodyToFootDot converts a foot dot to a body-center dot.
func (d *Dot) BodyToFootDot() *Dot {
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
	dFoot := d.BodyToFootDot()
	prevFoot := d.PrevDot.BodyToFootDot()
	midPoint := dotmath.ScalarMult(dotmath.AddPoints(dFoot.Point, prevFoot.Point), t)
	return NewDot(d.Name, count, 0, midPoint.X, midPoint.Y, false, d), nil
}

// Distance calculates the number of steps between this dot and the previous one.
func (d *Dot) Distance() float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := d.BodyToFootDot()
	prevFoot := d.PrevDot.BodyToFootDot()
	return dotmath.Distance(dFoot.Point, prevFoot.Point)
}

// StepSize calculates the step size needed to get to the given dot in the form
// of "x to 5" where x steps are needed to go 5 yards.
func (d *Dot) StepSize(f *FieldLayout) float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := d.BodyToFootDot()
	prevFoot := d.PrevDot.BodyToFootDot()
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

// GetDetails returns additional information about the dot, including step size,
// midset, crossing counts, and foot dot (if body-center)
func (d *Dot) GetDetails(f *FieldLayout) *DotDetails {
	dd := &DotDetails{
		Name:           d.Name,
		MoveCounts:     d.MoveCounts,
		HoldCounts:     d.HoldCounts,
		Point:          d.Point,
		BodyCenter:     d.BodyCenter,
		Midset:         nil,
		FootDot:        d.BodyToFootDot(),
		StepSize:       d.StepSize(f),
		CrossingCounts: d.CrossingCounts(f),
	}

	mid, err := d.DotOnCount(d.MoveCounts / 2)
	if err != nil {
		panic(err)
	} else {
		dd.Midset = mid
	}

	return dd
}
