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

func TestMultiply(t *testing.T){
	res := multiply(2,3)
	if res != 6 {
		t.Errorf("multiply was incorrect, got: %d want %d.", res, 6)
	}
}

func TestDivide(t *testing.T){
	res := divide(6,2)
	if res !=3{
	t.Errorf("divide was incorrect, got %d want %d.", res, 3)
	}
}
