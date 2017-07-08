package models

import (
	"github.com/nojnhuh/dotbook/models/dotmath"
)

// Represents a performer's dotbook for one show.
// Name: Any name by which to refer to this dotbook. Perhaps a show title.
// Dots: The ordered collection of dots.
// Field: The field layout for this dotbook. NCAA/HS football field, etc.
type Dotbook struct {
	Name  string
	Dots  []*Dot
	Field *fieldLayout
}

// NewDotbook creates a new Dotbook with no dots and returns a reference to it.
func NewDotbook(name string, field *fieldLayout) *Dotbook {
	return &Dotbook{name, []*Dot{}, field}
}

// AddDot creates a new dot from the arguments and adds it to the end of this
// Dotbook's list of dots.
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
		bodyCenter,
		prevDot,
	})
}

// AddDotByPointer adds the provided to the end of this Dotbook's list of dots.
func (db *Dotbook) AddDotByPointer(d *Dot) {
	d.PrevDot = db.Dots[len(db.Dots)-1]
	db.Dots = append(db.Dots, d)
}

// InsertDot inserts the given dot at the specified index in this Dotbook's
// list of dots.
func (db *Dotbook) InsertDot(d *Dot, i int) {
	d.PrevDot = db.Dots[i-1]
	db.Dots[i].PrevDot = d
	db.Dots = append(db.Dots, &Dot{})
	copy(db.Dots[i+1:], db.Dots[i:])
	db.Dots[i] = d
}

// getDotIndexByPointer finds the index of the specified dot in this
// Dotbook's list of dots and returns it, or -1 if not found.
func (db *Dotbook) getDotIndexByPointer(d *Dot) int {
	for i, dot := range db.Dots {
		if dot == d {
			return i
		}
	}
	return -1
}

// getDotIndexByName finds the index of the dot in this Dotbook's list of dots
// with the specified name and returns it, or -1 if not found.
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

// DeleteDotByPointer deletes the given dot from this Dotbook's list of dots.
func (db *Dotbook) DeleteDotByPointer(d *Dot) {
	db.DeleteDot(db.getDotIndexByPointer(d))
}

// DeleteDotByName deletes the first dot with the given name from this
// Dotbook's list of dots.
func (db *Dotbook) DeleteDotByName(name string) {
	db.DeleteDot(db.getDotIndexByName(name))
}
