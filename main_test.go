package main

import(
"testing"
)

func TestSum(t *testing.T){
	res := sum(1,1)
	if res != 2 {
		t.Errorf("sum was incorrect, got: %d want: %d.", res, 2)
	}
}

func TestMultiply(t *testin.T){
	res := multiply(2,3)
	if res != 6 {
		t.Errorf("multiply was incorrect, got: %d want %d.", res, 6)
	}
