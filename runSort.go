package main

import (
	"fmt"
	"goAlgorithms/sort"
)

func main() {
	fmt.Println("Buble Sort:")
	arrayzor1 := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	fmt.Println("Unsorted array: ", arrayzor1)
	sort.BubbleSort(arrayzor1)
	fmt.Println("Sorted array: ", arrayzor1)

	fmt.Println("Buble Modify Sort:")
	arrayzor2 := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	fmt.Println("Unsorted array: ", arrayzor2)
	sort.BubbleModifySort(arrayzor2)
	fmt.Println("Sorted array: ", arrayzor2)

	fmt.Println("Quick Sort:")
	arrayzor3 := []int{12, 4, 8, 57, 34, 96, 22, 111, 99, 3, 77, 55, 88, 121, 1, 888, 100, 777, 32, 47}
	fmt.Println("Unsorted array: ", arrayzor3)
	sort.QuickSort(arrayzor3, 0, 19)
	fmt.Println("Sorted array: ", arrayzor3)
}
