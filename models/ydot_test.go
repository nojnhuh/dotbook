package models

// import (
// 	"reflect"
// 	"testing"
// )
//
// type yDotToNumTest struct {
// 	dot      YDot
// 	field    *fieldLayout
// 	expected float64
// }
//
// type numToYDotTest struct {
// 	num      float64
// 	field    *fieldLayout
// 	expected YDot
// }
//
// type parseYDotTest struct {
// 	in       string
// 	field    *fieldLayout
// 	expected *YDot
// }
//
// type yDotStringTest struct {
// 	in       YDot
// 	expected string
// }
//
// var yDotToNumTests = []yDotToNumTest{
// 	{YDot{9001, Front, "FSL"}, MakeNCAAFootball(8), -9001},
// 	{YDot{1, Front, "FSL"}, MakeNCAAFootball(8), -1},
// 	{YDot{0, Front, "FSL"}, MakeNCAAFootball(8), 0},
// 	{YDot{0, Behind, "FSL"}, MakeNCAAFootball(8), 0},
// 	{YDot{1, Behind, "FSL"}, MakeNCAAFootball(8), 1},
// 	{YDot{16, Behind, "FSL"}, MakeNCAAFootball(8), 16},
//
// 	{YDot{16, Front, "FH"}, MakeNCAAFootball(8), 16},
// 	{YDot{0, Front, "FH"}, MakeNCAAFootball(8), 32},
// 	{YDot{0, Behind, "FH"}, MakeNCAAFootball(8), 32},
// 	{YDot{10, Behind, "FH"}, MakeNCAAFootball(8), 42},
//
// 	{YDot{10, Front, "BH"}, MakeNCAAFootball(8), 42},
// 	{YDot{0, Front, "BH"}, MakeNCAAFootball(8), 52},
// 	{YDot{0, Behind, "BH"}, MakeNCAAFootball(8), 52},
// 	{YDot{16, Behind, "BH"}, MakeNCAAFootball(8), 68},
//
// 	{YDot{16, Front, "BSL"}, MakeNCAAFootball(8), 68},
// 	{YDot{0, Front, "BSL"}, MakeNCAAFootball(8), 84},
// 	{YDot{0, Behind, "BSL"}, MakeNCAAFootball(8), 84},
// 	{YDot{9001, Behind, "BSL"}, MakeNCAAFootball(8), 9001 + 84},
// }
//
// var numToYDotTests = []numToYDotTest{
// 	{-9001, MakeNCAAFootball(8), YDot{9001, Front, "FSL"}},
// 	{-1, MakeNCAAFootball(8), YDot{1, Front, "FSL"}},
// 	{0, MakeNCAAFootball(8), YDot{0, Behind, "FSL"}},
// 	{1, MakeNCAAFootball(8), YDot{1, Behind, "FSL"}},
// 	{16, MakeNCAAFootball(8), YDot{16, Behind, "FSL"}},
// 	{16.5, MakeNCAAFootball(8), YDot{15.5, Front, "FH"}},
// 	{31.5, MakeNCAAFootball(8), YDot{0.5, Front, "FH"}},
// 	{32, MakeNCAAFootball(8), YDot{0, Behind, "FH"}},
// 	{42, MakeNCAAFootball(8), YDot{10, Behind, "FH"}},
// 	{42.5, MakeNCAAFootball(8), YDot{9.5, Front, "BH"}},
// 	{52, MakeNCAAFootball(8), YDot{0, Behind, "BH"}},
// 	{68, MakeNCAAFootball(8), YDot{16, Front, "BSL"}},
// 	{68.5, MakeNCAAFootball(8), YDot{15.5, Front, "BSL"}},
// 	{84, MakeNCAAFootball(8), YDot{0, Behind, "BSL"}},
// 	{9001, MakeNCAAFootball(8), YDot{9001 - 84, Behind, "BSL"}},
// }
//
// var parseYDotTests = []parseYDotTest{
// 	{"3 F FSL", MakeNCAAFootball(8), &YDot{3, Front, "FSL"}},
// 	{"1.25 B BH", MakeNCAAFootball(8), &YDot{1.25, Behind, "BH"}},
// 	{"4.0 F FH", MakeNCAAFootball(8), &YDot{4, Front, "FH"}},
// 	{"12.5 B FSL", MakeNCAAFootball(8), &YDot{12.5, Behind, "FSL"}},
// }
//
// var yDotStringTests = []yDotStringTest{
// 	{YDot{3, Front, "FSL"}, "3 FFSL"},
// 	{YDot{1.25, Behind, "BH"}, "1.25 BBH"},
// 	{YDot{1.2500, Behind, "BH"}, "1.25 BBH"},
// 	{YDot{4, Front, "FH"}, "4 FFH"},
// 	{YDot{12.5, Behind, "FSL"}, "12.5 BFSL"},
// }
//
// func TestYDotToNum(t *testing.T) {
// 	for _, test := range yDotToNumTests {
// 		v := test.dot.dotToNum(test.field)
// 		if v != test.expected {
// 			t.Error(
// 				"For dot", test.dot,
// 				"on field", test.field,
// 				"expected", test.expected,
// 				"got", v,
// 			)
// 		}
// 	}
// }
//
// func TestNumToYDot(t *testing.T) {
// 	for _, test := range numToYDotTests {
// 		v := numToYDot(test.num, test.field)
// 		if !reflect.DeepEqual(*v, test.expected) {
// 			t.Error(
// 				"For num", test.num,
// 				"on field", test.field,
// 				"expected", test.expected,
// 				"got", v,
// 			)
// 		}
// 	}
// }
//
// func TestParseYDot(t *testing.T) {
// 	for _, test := range parseYDotTests {
// 		v := parseYDot(test.in, test.field)
// 		if !reflect.DeepEqual(*v, *test.expected) {
// 			t.Error(
// 				"For string", test.in,
// 				"expected", test.expected,
// 				"got", v,
// 			)
// 		}
// 	}
// }
//
// func TestYDotString(t *testing.T) {
// 	for _, test := range yDotStringTests {
// 		v := test.in.String()
// 		if !reflect.DeepEqual(v, test.expected) {
// 			t.Error(
// 				"For dot", test.in,
// 				"expected", test.expected,
// 				"got", v,
// 			)
// 		}
// 	}
// }
