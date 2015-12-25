package sample

import (
	"testing"
)

func TestSum(t *testing.T){
	expected := 10
	result := Sum(3, 7)
	if result != expected {
		t.Errorf("Expected %v, but result %v", expected, result)
	}
}
