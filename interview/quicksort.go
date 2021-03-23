package sorting

import "fmt"

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Quick() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}
	quickSort(a, 0, len(a)-1)
	fmt.Println("Quick sort: ", a)
}

//func quickSort(a []int, left int, right int){
//	if left < right {
//		pivot := getPivot(a, left, right)
//		quickSort(a, left, pivot-1)
//		quickSort(a, pivot+1, right)
//	}
//}

//func getPivot(a []int, left int, right int) int {
//	low := left
//	high := right
//	pivot := left
//
//	for low < high {
//		for low < right && a[low] <= a[pivot] {
//			low = low + 1
//		}
//
//		for left < high && a[pivot] <= a[high] { // 조건
//			high = high - 1
//		}
//
//		if low < high {
//			temp := a[low]
//			a[low] = a[high]
//			a[high] = temp
//		}
//
//	}
//
//	if a[high] < a[pivot] {
//		temp := a[pivot]
//		a[pivot] = a[high]
//		a[high] = temp
//	}
//
//	return high
//}

//func quickSort(a []int, left int, right int) {
//	if left < right {
//		pivot := getPivot(a, left, right)
//		quickSort(a, left, pivot-1)
//		quickSort(a, pivot+1, right)
//	}
//}
//
//func getPivot(a []int, left int, right int) int {
//	low := left + 1
//	high := right
//	// pivot이 결국 left가 된다.
//	// for문 돌면서 low 값 설정
//	// for문 돌면서 high 값 설정
//
//	for low < high {
//		for low < right && a[low] <= a[left] {
//			low = low + 1
//		}
//
//		for left < high && a[left] <= a[high] {
//			high = high - 1
//		}
//
//		if low < high {
//			temp := a[low]
//			a[low] = a[high]
//			a[high] = temp
//		}
//	}
//
//	if a[left] > a[high] {
//		temp := a[left]
//		a[left] = a[high]
//		a[high] = temp
//	}
//
//	return high
//}

func quickSort(source []int, left int, right int) {
	if left < right {
		pivot := getPivot(source, left, right)
		quickSort(source, left, pivot-1)
		quickSort(source, pivot+1, right)
	}
}

func getPivot(source []int, left int, right int) int {
	high := right
	low := left
	low = low + 1

	for low < high {
		// 왼쪽
		for low < right && source[low] <= source[left] { // condition 조건 생각할 때 등호 주의
			low = low + 1
		}

		// 오른쪽
		for left < high && source[left] <= source[high] {
			high = high - 1
		}

		// low, high swap
		if low < high {
			temp := source[high]
			source[high] = source[low]
			source[low] = temp
		}
	}

	// high랑 pivot swap
	if source[left] > source[high] {
		temp := source[left]
		source[left] = source[high]
		source[high] = temp
	}

	return high
}