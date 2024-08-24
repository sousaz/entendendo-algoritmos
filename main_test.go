package main

import (
	"reflect"
	"testing"

	"github.com/sousaz/entendendo-algoritmos/binary_search"
	"github.com/sousaz/entendendo-algoritmos/selection_sort"
)

type test struct {
	data   []int
	target int
	answer int
}

type test2 struct {
	data   []int
	answer []int
}

func TestBinarySearch(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 5, answer: 5},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 7, answer: 7},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 12, answer: -1},
	}

	for _, v := range tests {
		r := binary_search.BinarySearch(v.data, v.target)
		if r != v.answer {
			t.Errorf("Esperado %d mas obteve %d", v.answer, r)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []test2{
		{data: []int{4, 5, 2, 8, 7, 1}, answer: []int{1, 2, 4, 5, 7, 8}},
		{data: []int{10, 5, 2, 20, 17, 50, 35}, answer: []int{2, 5, 10, 17, 20, 35, 50}},
		{data: []int{1, 2, 3, 4, 5, 6}, answer: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, v := range tests {
		r := selection_sort.SelectionSort(v.data)
		if !reflect.DeepEqual(r, v.answer) {
			t.Errorf("Esperado %d mas obteve %d", v.answer, r)
		}
	}
}
