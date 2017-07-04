package models

import (
	"log"
	"math"
	"sort"

	"github.com/nojnhuh/dotbook/models/dotmath"
)

// Represents a performer's dotbook for one show
// Name: Any name by which to refer to this dotbook. Perhaps a show title.
// Dots: The ordered collection of dots
// Field: The field layout for this dotbook. NCAA/HS football field, etc.
type Dotbook struct {
	Name  string
	Dots  []*Dot
	Field *fieldLayout
}

// Data returned from CrossingCounts method
// Line: Name of the line crossed
// Side: Side of the field the line is on
// Count: The count on which the line is crossed
type CrossCount struct {
	Line  string
	Side  string
	Count float64
}

// Creates a new Dotbook with no dots
func NewDotbook(name string, field *fieldLayout) *Dotbook {
	return &Dotbook{name, []*Dot{}, field}
}

// Creates a new dot from the arguments and adds it to the end of this Dotbook's
// list of dots
func (db *Dotbook) AddDot(name string, moveCounts, holdCounts float64,
	xdot, ydot string, bodyCenter bool) {
	f := db.Field
	var prevDot *Dot
	if len(db.Dots) > 0 {
		prevDot = db.Dots[len(db.Dots)-1]
	} else {
		prevDot = nil
	}
	db.Dots = append(db.Dots, &Dot{
		name,
		moveCounts,
		holdCounts,
		dotmath.NewPoint(parseXDot(xdot, f), parseYDot(ydot, f)),
		// parseYDot(ydot, f),
		bodyCenter,
		prevDot,
	})
}

// Adds the provided to the end of this Dotbook's list of dots
func (db *Dotbook) AddDotByPointer(d *Dot) {
	d.PrevDot = db.Dots[len(db.Dots)-1]
	db.Dots = append(db.Dots, d)
}

// Inserts the given dot at the specified index in this Dotbook's list of dots
func (db *Dotbook) InsertDot(d *Dot, i int) {
	d.PrevDot = db.Dots[i-1]
	db.Dots[i].PrevDot = d
	db.Dots = append(db.Dots, &Dot{})
	copy(db.Dots[i+1:], db.Dots[i:])
	db.Dots[i] = d
}

// Given a reference to a dot, finds its index in this Dotbook's list of dots
// and returns it, or -1 if not found
func (db *Dotbook) getDotIndexByPointer(d *Dot) int {
	for i, dot := range db.Dots {
		if dot == d {
			return i
		}
	}
	return -1
}

// Given the name of a dot, finds its index in this Dotbook's list of dots and
// returns it, or -1 if not found
func (db *Dotbook) getDotIndexByName(name string) int {
	for i, dot := range db.Dots {
		if dot.Name == name {
			return i
		}
	}
	return -1
}

// Deletes the dot at the given index in this Dotbook's list of dots
func (db *Dotbook) DeleteDot(i int) {
	db.Dots[i+1].PrevDot = db.Dots[i-1]
	copy(db.Dots[i:], db.Dots[i+1:])
	db.Dots[len(db.Dots)-1] = nil
	db.Dots = db.Dots[:len(db.Dots)-1]
}

// Deletes the given dot from this Dotbook's list of dots
func (db *Dotbook) DeleteDotByPointer(d *Dot) {
	db.DeleteDot(db.getDotIndexByPointer(d))
}

// Deletes the first dot with the given name from this Dotbook's list of dots
func (db *Dotbook) DeleteDotByName(name string) {
	db.DeleteDot(db.getDotIndexByName(name))
}

// Converts a foot dot to a body-center dot.
func (db *Dotbook) bodyToFootDot(d *Dot) *Dot {
	if d.BodyCenter {
		dot := *d
		if dot.PrevDot == nil {
			return &dot
		}
		dot.Point = dotmath.BodyToFootDot(d.Point, d.PrevDot.Point, d.MoveCounts)
		dot.BodyCenter = false
		return &dot
	} else {
		return d
	}
}

// Calculates an intermediate dot getting to the given dot based on the count
// parameter. Passing d.moveCounts/2 as count finds the midset.
func (db *Dotbook) DotOnCount(d *Dot, count float64) *Dot {
	if count > d.MoveCounts || count < 0 {
		log.Fatal("DotOnCount: Count %.0f invalid", count)
	}
	if d.PrevDot == nil {
		log.Fatalf("No dot before %s", d.Name)
	}
	t := count / float64(d.MoveCounts)
	dFoot := db.bodyToFootDot(d)
	prevFoot := db.bodyToFootDot(d.PrevDot)
	midPoint := dotmath.ScalarMult(dotmath.AddPoints(dFoot.Point, prevFoot.Point), t)
	mid := &Dot{d.Name, count, 0, midPoint, false, d}
	return mid
}

// Calculates the number of steps between this dot and the previous one.
func (db *Dotbook) Distance(d *Dot) float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := db.bodyToFootDot(d)
	prevFoot := db.bodyToFootDot(d.PrevDot)
	return dotmath.Distance(dFoot.Point, prevFoot.Point)
}

// Calculates the step size needed to get to the given dot
func (db *Dotbook) StepSize(d *Dot) float64 {
	if d.PrevDot == nil {
		return 0
	}
	dFoot := db.bodyToFootDot(d)
	prevFoot := db.bodyToFootDot(d.PrevDot)
	return db.Field.StepsPerFiveYards /
		dotmath.SegmentSize(dFoot.Point, prevFoot.Point, d.MoveCounts)
}

// Calculates the crossing counts getting to the given dot
func (db *Dotbook) CrossingCounts(d *Dot) []CrossCount {
	f := db.Field
	counts := []CrossCount{}
	if d.PrevDot == nil {
		return counts
	}
	prevX := d.PrevDot.Point.X
	thisX := d.Point.X
	for line, steps := range f.SideToSideLines {
		if (prevX < steps) != (thisX <= steps) {
			count := math.Abs((steps - prevX) / (prevX - thisX))
			counts = append(counts,
				CrossCount{line, "2", count * d.MoveCounts})
		}
		steps *= -1
		if (prevX < steps) != (thisX <= steps) && steps != 0 {
			count := math.Abs((steps - prevX) / (prevX - thisX))
			counts = append(counts,
				CrossCount{line, "1", count * d.MoveCounts})
		}
	}

	sort.Slice(counts[:], func(i, j int) bool {
		return counts[i].Count < counts[j].Count
	})
	return counts
}
