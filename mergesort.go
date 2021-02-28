package main

import "fmt"

func main() {
	// mergesort는 추가 공간이 필요함.
	a := []int{3,1,6,4,2,8,6,4,5,1}
	// 1,3,2,4,6 ... ,mid는 2잖아
	// 1,1,2,3,4,4,5,6,6,8
	b := make([]int, len(a))
	mergesort(a, b, 0, len(a)-1)
	fmt.Println(a)
}

func mergesort(source []int, temp []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergesort(source, temp, left, mid)
		mergesort(source, temp, mid+1, right)
		merge(source, temp, left, mid, right)
	}
}

func merge(source []int, temp []int, left int, mid int, right int) {

	k := left // temp array의 index
	low := left
	high := mid+1
	// 비교해서 채우기
	for low <= mid && high <= right{
		if source[low] > source[high] {
			temp[k] = source[high]
			high = high + 1
		} else {
			temp[k] = source[low]
			low = low + 1
		}
		k = k + 1
	}

	// 나머지 채우기
	if low > mid {
		// 오른쪽꺼 채우기
		for i:=high; i<=right; i++ {
			temp[k] = source[i]
			k = k + 1
		}

	} else {
		// 왼쪽꺼 채우기
		for i:=low; i<=mid; i++ {
			temp[k] = source[i]
			k = k + 1
		}
	}
	// 데이터 복사
	for i:=left; i<right+1; i++ {
		source[i] = temp[i]
	}
}
