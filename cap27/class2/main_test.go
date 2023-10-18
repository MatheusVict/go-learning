package main

import "testing"

type test struct {
	data   []int
	answer int
}

// TestSumTable Table tests
func TestSumTable(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{3, 9}, answer: 12},
		{data: []int{10, 10, 10}, answer: 30},
		{data: []int{7, 1, 100}, answer: 108},
	}
	for _, v := range tests {
		su := sum(v.data...)
		if su != v.answer {
			t.Error("Expected: ", v.answer, " got: ", su)
		}
	}
}

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
