package tests

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})

	if v != 1.5 {
		t.Error("Expected 1.5, got: ", v)
	}
}

func Average(float64s []float64) float64 {
	total := 0.0
	size := len(float64s)
	for _, v := range float64s {
		total += v
	}

	return total / float64(size)
}
