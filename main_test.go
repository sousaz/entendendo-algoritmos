package main

import (
	"testing"

	"github.com/sousaz/entendendo-algoritmos/binary_search"
)

type test struct {
	data   []int
	target int
	answer int
}

func TestBinarySearch(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 5, answer: 5},
		test{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 7, answer: 7},
		test{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 12, answer: -1},
	}

	for _, v := range tests {
		r := binary_search.BinarySearch(v.data, v.target)
		if r != v.answer {
			t.Errorf("Esperado %d mas obteve %d", v.answer, r)
		}
	}
}
