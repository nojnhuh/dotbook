package models

import (
	"testing"
)

func MakeTestDotbook() (*Dotbook, error) {
	db, err := NewDotbook("test", MakeNCAAFootball(8, 5))
	if err != nil {
		return nil, err
	}
	d0 := &Dot{Name: "d0"}
	d1 := &Dot{Name: "d1"}
	d2 := &Dot{Name: "d2"}
	db.AddDotByPointer(d0)
	db.AddDotByPointer(d1)
	db.AddDotByPointer(d2)
	return db, nil
}

func TestNewDotbook(t *testing.T) {
	var newDotbookTests = []struct {
		name      string
		field     *FieldLayout
		shouldErr bool
	}{
		{"Test DB", MakeNCAAFootball(8, 5), false},
		{"Test DB", MakeHSFootball(8, 5), false},
		{"", MakeNCAAFootball(8, 5), true},
		{"Test DB", nil, true},
	}
	for _, test := range newDotbookTests {
		db, err := NewDotbook(test.name, test.field)
		if (err != nil) != test.shouldErr {
			t.Errorf("Creation of %v expected error %v, got error %v.", db, test.shouldErr, err != nil)
		} else if db != nil {
			if db == nil {
				t.Errorf("New dotbook should not be nil")
			}
			if db.Name != test.name {
				t.Errorf("New dotbook name incorrect. Expected %v, got %v.", test.name, db.Name)
			}
			if len(db.Dots) != 0 {
				t.Errorf("New dotbook should have no dots, but has %d.", len(db.Dots))
			}
			if db.Field == nil {
				t.Errorf("Field should not be nil.")
			}
		}
	}
}

func TestAddDot(t *testing.T) {
	db, err := NewDotbook("Test", MakeNCAAFootball(8, 5))
	if err != nil {
		panic(err)
	}
	var addDotTests = []struct {
		db                     Dotbook
		name                   string
		moveCounts, holdCounts float64
		xdot, ydot             string
		bodyCenter             bool
		shouldErr              bool
	}{
		{*db, "1", 0, 8, "0 O A50", "0 B FSL", false, false},
		{*db, "1", 0, 8, "Not an XDot", "0 B FSL", false, true},
		{*db, "1", 0, 8, "0 O A50", "Not a YDot", false, true},
		{*db, "1", 0, 0, "0 O A50", "0 B FSL", false, false},
	}

	for _, test := range addDotTests {
		oldLength := len(test.db.Dots)
		err := test.db.AddDot(test.name, test.moveCounts, test.holdCounts, test.xdot, test.ydot, test.bodyCenter)
		if (err != nil) != test.shouldErr {
			t.Errorf("Adding dot %s %s to %v expected error %v, got error %v.", test.xdot, test.ydot, test.db, test.shouldErr, err != nil)
		} else if err == nil {
			if len(test.db.Dots) != oldLength+1 {
				t.Errorf("Expected %d dots, got %d dots.", oldLength+1, len(test.db.Dots))
			}
		}
	}
}

func TestAddDotByPointer(t *testing.T) {
	db, err := NewDotbook("Test", MakeNCAAFootball(8, 5))
	if err != nil {
		t.Error("Making new dotbook should not fail.")
	}
	dot := new(Dot)
	db.AddDotByPointer(dot)
	if len(db.Dots) != 1 {
		t.Errorf("Dotbook should have 1 dot, has %v.", len(db.Dots))
	}
	if !dot.equals(db.Dots[0]) {
		t.Error("New dot is not equal to dot in slice.")
	}
	if dot.PrevDot != nil {
		t.Error("First dot PrevDot should be nil.")
	}
	dot2 := new(Dot)
	dot2.MoveCounts = 8
	if dot.equals(dot2) {
		t.Error("Dots should not be equal")
	}
	db.AddDotByPointer(dot2)
	if !dot2.equals(db.Dots[len(db.Dots)-1]) {
		t.Error("Second new dot is not the last dot")
	}
	if !dot.equals(dot2.PrevDot) {
		t.Error("PrevDot not set correctly")
	}
}

func TestInsertDot(t *testing.T) {
	db, err := NewDotbook("Test", MakeNCAAFootball(8, 5))
	if err != nil {
		t.Error("Making new dotbook should not fail.")
	}
	d1 := new(Dot)
	err = db.InsertDot(d1, 1)
	if err == nil {
		t.Error("Should not be able to insert a dot at index 1 in an empty dotbook")
	}
	err = db.InsertDot(d1, 0)
	if err != nil {
		t.Error("Inserting dot at index 0 in an empty dotbook should not fail")
	}
	if len(db.Dots) != 1 {
		t.Error("Adding new dot does not change slice length")
	}
	d3 := new(Dot)
	err = db.InsertDot(d3, 1)
	if err != nil {
		t.Errorf("Inserting dot at end of list should not fail. Error: %v", err)
	}
	if d3.PrevDot != d1 {
		t.Error("PrevDot link not set correctly")
	}
	d0 := new(Dot)
	err = db.InsertDot(d0, 0)
	if err != nil {
		t.Error("Inserting at the beginning of the list should not fail")
	}
	if d1.PrevDot != d0 {
		t.Error("PrevDot on existing dot not set correctly")
	}
	if d0.PrevDot != nil {
		t.Error("PrevDot on new first dot not nil")
	}
	d2 := new(Dot)
	err = db.InsertDot(d2, 2)
	if err != nil {
		t.Error("Inserting dot in middle should not fail")
	}
	if d2.PrevDot != d1 {
		t.Error("PrevDot link on new dot not set correctly")
	}
	if d3.PrevDot != d2 {
		t.Error("PrevDot link on existing dot not set correctly")
	}
}

func TestGetDotIndexByPointer(t *testing.T) {
	db, err := NewDotbook("test", MakeNCAAFootball(8, 5))
	if err != nil {
		t.Error("Making new dotbook should not fail.")
	}
	d0 := &Dot{Name: "d0"}
	d1 := &Dot{Name: "d1"}
	d2 := &Dot{Name: "d2"}
	db.Dots = []*Dot{d0, d1, d2}
	var getGotIndexByPointerTests = []struct {
		dot      *Dot
		expected int
	}{
		{d0, 0},
		{d1, 1},
		{d2, 2},
		{new(Dot), -1},
	}

	for _, test := range getGotIndexByPointerTests {
		v := db.getDotIndexByPointer(test.dot)
		if v != test.expected {
			t.Errorf("Expected to find %v at %v, got %v", test.dot.Name, test.expected, v)
		}
	}
}

func TestGetDotIndexByName(t *testing.T) {
	db, err := NewDotbook("test", MakeNCAAFootball(8, 5))
	if err != nil {
		t.Error("Making new dotbook should not fail.")
	}
	d0 := &Dot{Name: "d0"}
	d1 := &Dot{Name: "d1"}
	d2 := &Dot{Name: "d2"}
	db.AddDotByPointer(d0)
	db.AddDotByPointer(d1)
	db.AddDotByPointer(d2)
	var getGotIndexByNameTests = []struct {
		dot      string
		expected int
	}{
		{"d0", 0},
		{"d1", 1},
		{"d2", 2},
		{"", -1},
	}

	for _, test := range getGotIndexByNameTests {
		v := db.getDotIndexByName(test.dot)
		if v != test.expected {
			t.Errorf("Expected to find %v at %v, got %v", test.dot, test.expected, v)
		}
	}
}

func TestDeleteDot(t *testing.T) {
	db, err := MakeTestDotbook()
	if err != nil {
		t.Error("Making new dotbook should not fail.")
	}
	db.DeleteDot(0)
	if len(db.Dots) != 2 {
		t.Errorf("Expected length 2, got length %v", len(db.Dots))
	}
	if db.Dots[0].PrevDot != nil {
		t.Error("PrevDot not set correctly")
	}
	db, _ = MakeTestDotbook()
	db.DeleteDot(1)
	if len(db.Dots) != 2 {
		t.Errorf("Expected length 2, got length %v", len(db.Dots))
	}
	if db.Dots[1].PrevDot != db.Dots[0] {
		t.Error("PrevDot not set correctly")
	}
	db, _ = MakeTestDotbook()
	db.DeleteDot(2)
	if len(db.Dots) != 2 {
		t.Errorf("Expected length 2, got length %v", len(db.Dots))
	}
	if db.Dots[1].PrevDot != db.Dots[0] {
		t.Error("PrevDot not set correctly")
	}
}
