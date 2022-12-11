package main

import(
"testing"
)

func TestMain(t *testing.T){
	res := sum(1,1)
	if res != 2 {
		t.Errorf("sum was incorrect, got: %d want: %d.", res, 2)
	}
}
