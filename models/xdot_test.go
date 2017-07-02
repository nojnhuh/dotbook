package models

//
// import (
// 	"reflect"
// 	"testing"
// )
//
// type xDotToNumTest struct {
// 	dot      XDot
// 	field    *fieldLayout
// 	expected float64
// }
//
// type numToXDotTest struct {
// 	num      float64
// 	field    *fieldLayout
// 	expected XDot
// }
//
// type parseXDotTest struct {
// 	in       string
// 	field    *fieldLayout
// 	expected *XDot
// }
//
// type xDotStringTest struct {
// 	in       XDot
// 	expected string
// }
//
// var xDotToNumTests = []xDotToNumTest{
// 	{XDot{0, Inside, Left, "50"}, MakeNCAAFootball(8), 0},
// 	{XDot{0, Outside, Left, "50"}, MakeNCAAFootball(8), 0},
// 	{XDot{0, Inside, Right, "50"}, MakeNCAAFootball(8), 0},
// 	{XDot{0, Outside, Right, "50"}, MakeNCAAFootball(8), 0},
// 	{XDot{1, Inside, Left, "50"}, MakeNCAAFootball(8), -1},
// 	{XDot{1, Outside, Left, "50"}, MakeNCAAFootball(8), -1},
// 	{XDot{1, Inside, Right, "50"}, MakeNCAAFootball(8), 1},
// 	{XDot{1, Outside, Right, "50"}, MakeNCAAFootball(8), 1},
//
// 	{XDot{1, Inside, Left, "45"}, MakeNCAAFootball(8), -7},
// 	{XDot{0, Inside, Left, "45"}, MakeNCAAFootball(8), -8},
// 	{XDot{0, Outside, Left, "45"}, MakeNCAAFootball(8), -8},
// 	{XDot{1, Outside, Left, "45"}, MakeNCAAFootball(8), -9},
//
// 	{XDot{1, Inside, Right, "45"}, MakeNCAAFootball(8), 7},
// 	{XDot{0, Inside, Right, "45"}, MakeNCAAFootball(8), 8},
// 	{XDot{0, Outside, Right, "45"}, MakeNCAAFootball(8), 8},
// 	{XDot{1, Outside, Right, "45"}, MakeNCAAFootball(8), 9},
//
// 	{XDot{1, Inside, Left, "0"}, MakeNCAAFootball(8), -79},
// 	{XDot{0, Inside, Left, "0"}, MakeNCAAFootball(8), -80},
// 	{XDot{0, Outside, Left, "0"}, MakeNCAAFootball(8), -80},
// 	{XDot{1, Outside, Left, "0"}, MakeNCAAFootball(8), -81},
// 	{XDot{9001, Outside, Left, "0"}, MakeNCAAFootball(8), -9001 - 80},
//
// 	{XDot{1, Inside, Right, "0"}, MakeNCAAFootball(8), 79},
// 	{XDot{0, Inside, Right, "0"}, MakeNCAAFootball(8), 80},
// 	{XDot{0, Outside, Right, "0"}, MakeNCAAFootball(8), 80},
// 	{XDot{1, Outside, Right, "0"}, MakeNCAAFootball(8), 81},
// 	{XDot{9001, Outside, Right, "0"}, MakeNCAAFootball(8), 9001 + 80},
// }
//
// var numToXDotTests = []numToXDotTest{
// 	{0, MakeNCAAFootball(8), XDot{0, Outside, Left, "50"}},
// 	{-1, MakeNCAAFootball(8), XDot{1, Outside, Left, "50"}},
// 	{1, MakeNCAAFootball(8), XDot{1, Outside, Right, "50"}},
// 	{4, MakeNCAAFootball(8), XDot{4, Inside, Right, "50"}},
// 	{-4, MakeNCAAFootball(8), XDot{4, Inside, Left, "50"}},
//
// 	{80, MakeNCAAFootball(8), XDot{0, Outside, Right, "0"}},
// 	{79, MakeNCAAFootball(8), XDot{1, Inside, Right, "0"}},
// 	{-80, MakeNCAAFootball(8), XDot{0, Outside, Left, "0"}},
// 	{-79, MakeNCAAFootball(8), XDot{1, Inside, Left, "0"}},
//
// 	{9001, MakeNCAAFootball(8), XDot{9001 - 80, Outside, Right, "0"}},
// 	{-9001, MakeNCAAFootball(8), XDot{9001 - 80, Outside, Left, "0"}},
//
// 	{10, MakeNCAAFootball(8), XDot{2, Outside, Right, "45"}},
// 	{-10, MakeNCAAFootball(8), XDot{2, Outside, Left, "45"}},
// 	{8, MakeNCAAFootball(8), XDot{0, Outside, Right, "45"}},
// 	{-8, MakeNCAAFootball(8), XDot{0, Outside, Left, "45"}},
// 	{6, MakeNCAAFootball(8), XDot{2, Inside, Right, "45"}},
// 	{-6, MakeNCAAFootball(8), XDot{2, Inside, Left, "45"}},
//
// 	{0, MakeHSFootball(8), XDot{0, Outside, Left, "50"}},
// 	{-1, MakeHSFootball(8), XDot{1, Outside, Left, "50"}},
// 	{1, MakeHSFootball(8), XDot{1, Outside, Right, "50"}},
// 	{4, MakeHSFootball(8), XDot{4, Inside, Right, "50"}},
// 	{-4, MakeHSFootball(8), XDot{4, Inside, Left, "50"}},
//
// 	{80, MakeHSFootball(8), XDot{0, Outside, Right, "0"}},
// 	{79, MakeHSFootball(8), XDot{1, Inside, Right, "0"}},
// 	{-80, MakeHSFootball(8), XDot{0, Outside, Left, "0"}},
// 	{-79, MakeHSFootball(8), XDot{1, Inside, Left, "0"}},
//
// 	{9001, MakeHSFootball(8), XDot{9001 - 80, Outside, Right, "0"}},
// 	{-9001, MakeHSFootball(8), XDot{9001 - 80, Outside, Left, "0"}},
//
// 	{10, MakeHSFootball(8), XDot{2, Outside, Right, "45"}},
// 	{-10, MakeHSFootball(8), XDot{2, Outside, Left, "45"}},
// 	{8, MakeHSFootball(8), XDot{0, Outside, Right, "45"}},
// 	{-8, MakeHSFootball(8), XDot{0, Outside, Left, "45"}},
// 	{6, MakeHSFootball(8), XDot{2, Inside, Right, "45"}},
// 	{-6, MakeHSFootball(8), XDot{2, Inside, Left, "45"}},
// }
//
// var parseXDotTests = []parseXDotTest{
// 	{"1.5 I 1 40", MakeNCAAFootball(8), &XDot{1.5, Inside, Left, "40"}},
// 	{"2 O 2 5", MakeNCAAFootball(8), &XDot{2, Outside, Right, "5"}},
// }
//
// var xDotStringTests = []xDotStringTest{
// 	{XDot{1.5, Inside, Left, "40"}, "1.5 I 1 40"},
// 	{XDot{2, Outside, Right, "5"}, "2 O 2 5"},
// }
//
// func TestXDotToNum(t *testing.T) {
// 	for _, test := range xDotToNumTests {
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
// func TestNumToXDot(t *testing.T) {
// 	for _, test := range numToXDotTests {
// 		v := numToXDot(test.num, test.field)
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
// func TestParseXDot(t *testing.T) {
// 	for _, test := range parseXDotTests {
// 		v := parseXDot(test.in, test.field)
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
// func TestXDotString(t *testing.T) {
// 	for _, test := range xDotStringTests {
// 		v := test.in.String()
// 		if v != test.expected {
// 			t.Error(
// 				"For dot", test.in,
// 				"expected", test.expected,
// 				"got", v,
// 			)
// 		}
// 	}
// }
