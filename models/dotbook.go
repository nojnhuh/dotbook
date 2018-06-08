package models

import (
	"fmt"
	"log"
)

// Dotbook represents a performer's dotbook for one show.
// Name: Any name by which to refer to this dotbook. Perhaps a show title.
// Dots: The ordered collection of dots.
// Field: The field layout for this dotbook. NCAA/HS football field, etc.
type Dotbook struct {
	Name  string       `json:"name"`
	Dots  []*Dot       `json:"dots"`
	Field *FieldLayout `json:"field"`
}

// NewDotbook creates a new Dotbook with no dots and returns a reference to it.
func NewDotbook(name string, field *FieldLayout) (*Dotbook, error) {
	if field == nil {
		return nil, fmt.Errorf("Dotbook field layout cannot be nil")
	}
	if len(name) == 0 {
		return nil, fmt.Errorf("Dotbook name cannot be empty")
	}
	return &Dotbook{Name: name, Dots: []*Dot{}, Field: field}, nil
}

// AddDot creates a new dot from the arguments and adds it to the end of this
// Dotbook's list of dots.
func (db *Dotbook) AddDot(name string, moveCounts, holdCounts float64,
	xdot, ydot string, bodyCenter bool) error {
	if len(name) == 0 {
		return fmt.Errorf("Dot name cannot be empty")
	}
	var prevDot *Dot
	if len(db.Dots) > 0 {
		prevDot = db.Dots[len(db.Dots)-1]
	} else {
		prevDot = nil
	}
	f := db.Field
	xCoord, err := parseXDot(xdot, f)
	if err != nil {
		log.Printf("Error creating coordinate from '%s'.\n", xdot)
		return err
	}
	yCoord, err := parseYDot(ydot, f)
	if err != nil {
		log.Printf("Error creating coordinate from '%s'.\n", ydot)
		return err
	}
	db.Dots = append(db.Dots, NewDot(name, moveCounts, holdCounts, xCoord, yCoord, bodyCenter, prevDot))
	return nil
}

// AddDotByPointer adds the provided to the end of this Dotbook's list of dots.
func (db *Dotbook) AddDotByPointer(d *Dot) {
	if len(db.Dots) == 0 {
		d.PrevDot = nil
	} else {
		d.PrevDot = db.Dots[len(db.Dots)-1]
	}
	db.Dots = append(db.Dots, d)
}

// InsertDot inserts the given dot at the specified index in this Dotbook's
// list of dots.
func (db *Dotbook) InsertDot(d *Dot, i int) error {
	dotsLength := len(db.Dots)
	if i < 0 || i > dotsLength && i != 0 {
		return fmt.Errorf("Index %v is not between 0 and %v", i, dotsLength)
	}

	if i == 0 {
		d.PrevDot = nil
	} else {
		d.PrevDot = db.Dots[i-1]
	}
	if i < dotsLength {
		db.Dots[i].PrevDot = d
	}
	db.Dots = append(db.Dots, new(Dot))
	copy(db.Dots[i+1:], db.Dots[i:])
	db.Dots[i] = d
	return nil
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

// DeleteDot Deletes the dot at the given index in this Dotbook's list of dots
func (db *Dotbook) DeleteDot(i int) error {
	if i < 0 || i >= len(db.Dots) {
		return fmt.Errorf("Index out of bounds: %v", i)
	}
	if i < len(db.Dots)-1 {
		if i == 0 {
			db.Dots[1].PrevDot = nil
		} else {
			db.Dots[i+1].PrevDot = db.Dots[i-1]
		}
		copy(db.Dots[i:], db.Dots[i+1:])
	}
	db.Dots[len(db.Dots)-1] = nil
	db.Dots = db.Dots[:len(db.Dots)-1]
	return nil
}

// DeleteDotByPointer deletes the given dot from this Dotbook's list of dots.
func (db *Dotbook) DeleteDotByPointer(d *Dot) error {
	return db.DeleteDot(db.getDotIndexByPointer(d))
}

// DeleteDotByName deletes the first dot with the given name from this
// Dotbook's list of dots.
func (db *Dotbook) DeleteDotByName(name string) error {
	return db.DeleteDot(db.getDotIndexByName(name))
}
