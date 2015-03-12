package sort

func BubbleSort(array []int) {
	arrLen := len(array)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func QuickSort(array []int, left int, right int) {
	if left > right {
		return
	}
	base := array[left]
	i := left
	j := right

	for i != j {

		for i < j && array[j] >= base {
			j--
		}

		for i < j && array[i] <= base {
			i++
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[left], array[i] = array[i], base

	QuickSort(array, left, i-1)
	QuickSort(array, i+1, right)
}

//I thought it base on the BubbleSort
func BubbleModifySort(array []int) {
	arrLen := len(array)
	for i := 0; i < arrLen; i++ {
		base := 0
		for j := 0; j < arrLen-1-i; j++ {
			if array[base] < array[j+1] {
				base = j + 1
			}
		}
		if base != arrLen-1-i {
			array[base], array[arrLen-1-i] = array[arrLen-1-i], array[base]
		}
	}
}
