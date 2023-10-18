package main

import "testing"

type test struct {
	data   []int
	answer int
}

// BenchmarkSum Run benchmark tests for functions
func BenchmarkSum(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{3, 9}, answer: 12},
		{data: []int{10, 10, 10}, answer: 30},
		{data: []int{7, 1, 100}, answer: 108},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			sum(v.data...)
		}
	}
}

func BenchmarkMultiple(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{3, 9}, answer: 12},
		{data: []int{10, 10, 10}, answer: 30},
		{data: []int{7, 1, 100}, answer: 108},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			multiple(v.data...)
		}
	}
}
