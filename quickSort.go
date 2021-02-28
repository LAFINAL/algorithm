package main

import "fmt"

func main() {
	a := []int{3,1,6,4,2,8,6,4,5}
	fmt.Println("before quicksort: ", a)
	quickSort(a, 0, len(a)-1)
	fmt.Println("after quicksort: ", a)
}

func quickSort(source []int, left int, right int) {

	if left < right {
		pivot := getPivot(source, left, right)

		quickSort(source, left, pivot-1)
		quickSort(source, pivot+1, right)
	}
}

func getPivot(source []int, left int, right int) int {
	pivot := left
	low := left
	high := right
	low = low + 1
	for low < high {
		for low < right && source[low] < source[pivot] {
			low = low + 1
		}

		for high > left && source[high] > source[pivot] {
			high = high - 1
		}

		if low < high {
			temp := source[low]
			source[low] = source[high]
			source[high] = temp
		}
	}

	if source[pivot] > source[high] {
		temp := source[high]
		source[high] = source[pivot]
		source[pivot] = temp
	}
	return high
}

