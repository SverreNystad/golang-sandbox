package main

import "testing"

func TestAdd(t *testing.T) {
	result := Calculate(2, 3, "+")
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := Calculate(3, 2, "-")
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestMul(t *testing.T) {
	result := Calculate(2, 3, "*")
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result := Calculate(6, 3, "/")
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}
func TestDivZero(t *testing.T) {
	result := Calculate(6, 0, "/")
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestMod(t *testing.T) {
	result := Calculate(5, 3, "%")
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestExponent(t *testing.T) {
	result := Calculate(2, 3, "^")
	if result != 8 {
		t.Errorf("Expected 8, but got %d", result)
	}
}

func TestExponentZero(t *testing.T) {
	result := Calculate(2, 0, "^")
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

// func TestSqrt(t *testing.T) {
// 	result := sqrt(4.0)
// 	if result != 2.0 {
// 		t.Errorf("Expected 2.0, but got %f", result)
// 	}
// }
