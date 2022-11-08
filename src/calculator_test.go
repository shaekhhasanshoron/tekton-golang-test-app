package src

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMinus(t *testing.T) {
	total := Minus(10, 5)
	if total != 5 {
		t.Errorf("Minus was incorrect, got: %d, want: %d.", total, 5)
	}
}

//
//func TestMinus2(t *testing.T) {
//	total := Minus(5, 5)
//	if total != 5 {
//		t.Errorf("Minus was incorrect, got: %d, want: %d.", total, 5)
//	}
//}
