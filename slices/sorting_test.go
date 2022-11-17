package slices

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	expectedOrder := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice := []int{9, 7, 3, 8, 1, 6, 2, 0, 4, 5}
	SelectionSort(slice)

	if !reflect.DeepEqual(expectedOrder, slice) {
		t.Errorf("got %v, want %v", slice, expectedOrder)
	}
}
