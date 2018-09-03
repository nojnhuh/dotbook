package models

import (
	"testing"
)

func TestParseXDot(t *testing.T) {
	field := MakeNCAAFootball(8, 5)
	var parseXDotTests = []struct {
		in        string
		expected  float64
		shouldErr bool
	}{
		{"0 I A50", 0, false},
		{"1 I A50", -1, false},
		{"1 O A50", -1, false},
		{"1 I B50", 1, false},
		{"1 O B50", 1, false},
		{"2 O A45", -10, false},
		{"2 I A45", -6, false},
		{"4 I A45", -4, false},
		{"0 I A45", -8, false},
		{"0 O A45", -8, false},
		{"2 O B45", 10, false},
		{"2 I B45", 6, false},
		{"4 I B45", 4, false},
		{"0 I B45", 8, false},
		{"0 O B45", 8, false},

		{"Not even close to being a real dot", 0, true},
		{"2 I ANotAYardline", 0, true},
		{"NotANumber I 50", 0, true},
		{"2 NotAnAcceptableValue 50", 0, true},
		{"2 N A40", 0, true},
	}

	for _, test := range parseXDotTests {
		v, err := parseXDot(test.in, field)
		if (err != nil) == test.shouldErr {
			if v != test.expected {
				t.Errorf("For input '%v' on NCAA field, expected %f, got %f", test.in, test.expected, v)
			}
		} else {
			t.Errorf("Input '%v' error: %v, expected %v", test.in, err != nil, test.shouldErr)
		}
	}
}

func TestParseYDot(t *testing.T) {
	ncaa := MakeNCAAFootball(8, 5)
	hs := MakeHSFootball(8, 5)
	var parseYDotTests = []struct {
		in                       string
		ncaaExpected, hsExpected float64
		shouldErr                bool
	}{
		{"0 B FSL", 0, 0, false},
		{"0 F FSL", 0, 0, false},
		{"1 F FSL", -1, -1, false},
		{"9001 F FSL", -9001, -9001, false},
		{"1 B FSL", 1, 1, false},
		{"14 B FSL", 14, 14, false},
		{"9001 B FSL", 9001, 9001, false},
		{"28 F FH", 4, 0, false},
		{"32 F FH", 0, -4, false},
		{"0 F FH", 32, 28, false},
		{"0 B FH", 32, 28, false},
		{"10 B FH", 42, 38, false},
		{"10 F BH", 42, 46, false},

		{"Not even close to being a real dot", 0, 0, true},
		{"2 F ANotAYardline", 0, 0, true},
		{"NotANumber F FH", 0, 0, true},
		{"2 NotAnAcceptableValue BH", 0, 0, true},
		{"2 N BSL", 0, 0, true},
	}

	for _, test := range parseYDotTests {
		v, err := parseYDot(test.in, ncaa)
		if (err != nil) == test.shouldErr {
			if v != test.ncaaExpected {
				t.Errorf("For input '%v' on NCAA field, expected %f, got %f", test.in, test.ncaaExpected, v)
			}
		}
		v, err = parseYDot(test.in, hs)
		if (err != nil) == test.shouldErr {
			if v != test.hsExpected {
				t.Errorf("For input '%v' on HS field, expected %f, got %f", test.in, test.hsExpected, v)
			}
		}
	}
}
