package main

import "testing"

func TestSum(t *testing.T) {
	x := sum(3, 2, 1)
	result := 6
	if x != result {
		t.Error("Expected: ", result, " got: ", x)
	}
}

func TestMultiple(t *testing.T) {
	x := multiple(1, 3, 3)
	result := 9
	if x != result {
		t.Error("Expected: ", result, " got: ", x)
	}
}

//Test to fail
/*func TestSumToFail(t *testing.T) {
	x := sum(3, 2, 1)
	result := 7
	if x != result {
		t.Error("Expected: ", result, " received: ", x)
	}
}*/
