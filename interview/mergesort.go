package sorting

import "fmt"

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Merge() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}
	mergeSort(a, 0, len(a)-1)
	fmt.Println("Merge sort: ", a)
}

//func mergeSort(a []int, left int, right int){
//	if left < right {
//		mid := (left + right) / 2
//		mergeSort(a, left, mid)
//		mergeSort(a, mid+1, right)
//		merge(a, left, right, mid)
//	}
//}
//
//func merge(a []int, left int, right int, mid int) {
//	// temp 배열 만들어야 함.
//	temp := make([]int, len(a))
//	one := left
//	two := mid+1
//	k := left
//
//	// 왼쪽 오른쪽 비교하면서 작은 값부터 넣기.
//	for one <= mid && two <= right {
//		if a[one] < a[two] {
//			temp[k] = a[one]
//			one = one + 1
//		} else {
//			temp[k] = a[two]
//			two = two + 1
//		}
//		k = k + 1
//	}
//	// 나머지 값 넣기.
//	if one > mid {
//		// two에서 전부 넣어야 함
//		for two <= right {
//			temp[k] = a[two]
//			k = k + 1
//			two = two + 1
//		}
//
//	} else {
//		for one <= mid {
//			temp[k] = a[one]
//			k = k + 1
//			one = one + 1
//		}
//	}
//	// temp에서 원래 a로 복사
//	for i:=left; i<right+1; i++ {
//		a[i] = temp[i]
//	}
//}

func mergeSort(source []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(source, left, mid)
		mergeSort(source, mid+1, right)
		merge(source, left, right)
	}
}

func merge(source []int, left int, right int) {
	mid := (left+right)/2
	low := left
	high := mid
	high = high + 1
	// temp 배열을 하나 생성. 추가 공간이 필요하다.
	temp := make([]int, len(source), len(source))
	k := left

	// left부터 mid, mid+1부터 right 까지 돈다.
	for low <= mid && high <= right {
		if source[low] < source[high] {
			temp[k] = source[low]
			low = low + 1
		} else {
			temp[k] = source[high]
			high = high + 1
		}
		k = k + 1
	}

	// 나머지 부분 채워줘야 한다.
	if low > mid {
		// 오른쪽 부분 채우기
		for high <= right {
			temp[k] = source[high]
			high = high + 1
			k = k + 1
		}

	} else {
		// 왼쪽 부분 채우기
		for low <= mid {
			temp[k] = source[low]
			low = low + 1
			k = k + 1
		}
	}

	for i:=left; i<=right; i++ {
		source[i] = temp[i]
	}
}