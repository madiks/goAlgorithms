package sort

import (
	"testing"
)

func Benchmark_BubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		BubbleSort([]int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47})
	}
}

func Benchmark_BubbleModifySort(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		BubbleModifySort([]int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47})
	}
}

func Benchmark_QuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		QuickSort([]int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}, 0, 19)
	}
}
