package models

import (
	"github.com/nojnhuh/dotbook/models/dotmath"
)

// Represents one full coordinate
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
