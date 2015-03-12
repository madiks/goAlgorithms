package sort

import (
	"testing"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Test_BubbleSort(t *testing.T) {
	array := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	wanted := []int{1, 3, 4, 8, 12, 22, 32, 34, 47, 55, 57, 77, 88, 96, 99, 100, 111, 121, 777, 888}
	BubbleSort(array)
	if !testEq(array, wanted) {
		t.Fail()
	}
}

func Test_BubbleModifySort(t *testing.T) {
	array := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	wanted := []int{1, 3, 4, 8, 12, 22, 32, 34, 47, 55, 57, 77, 88, 96, 99, 100, 111, 121, 777, 888}
	BubbleModifySort(array)
	if !testEq(array, wanted) {
		t.Fail()
	}
}

func Test_QuickSort(t *testing.T) {
	array := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	wanted := []int{1, 3, 4, 8, 12, 22, 32, 34, 47, 55, 57, 77, 88, 96, 99, 100, 111, 121, 777, 888}
	QuickSort(array, 0, 19)
	if !testEq(array, wanted) {
		t.Fail()
	}
}
