package num

import (
	"testing"
)

type testval_rec struct {
	x float64
	floor_x int64
	ceil_x int64
}

// Test data from https://www.aa.quae.nl/en/reken/juliaansedag.html
var testvals = []testval_rec {
	//  x   fl  cl
	{ -5.0,	-5,	-5},
	{ -4.9,	-5,	-4},
	{ -4.2,	-5,	-4},
	{ -4.0,	-4,	-4},
	{ -0.2,	-1,	0},
	{ 0.0, 	0, 	0},
	{ 0.2, 	0, 	1},
	{ 4.0, 	4, 	4},
	{ 4.2, 	4, 	5},
	{ 4.9, 	4, 	5},
	{ 5.0, 	5, 	5},
}

func TestFloor(t *testing.T) {
	for _, v := range testvals {
		result := Floor(v.x)
		if result != v.floor_x {
			t.Fatalf("floor(%v): expected %v, got %v", v.x, v.floor_x, result)
		}
	}
}

func TestCeil(t *testing.T) {
	for _, v := range testvals {
		result := Ceil(v.x)
		if result != v.ceil_x {
			t.Fatalf("ceil(%v): expected %v, got %v", v.x, v.ceil_x, result)
		}
	}
}
