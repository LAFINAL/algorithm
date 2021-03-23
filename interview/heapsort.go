package sorting

import "fmt"

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Heap() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}

	// 일단 heap 자료 구조로 만들고나서
	// heap 자료 구조로 만드는 법은 절반으로 자른다음 child node 검사
	for i:=len(a)/2; i>=0; i-- {
		heapify(a, i, len(a))
	}
	// heap 자료 구조로 만든 다음에는 for문 돌아가면서 다시 heapify heapify가 자리 찾아가는거니까.
	for i:=len(a)-1; i>=1; i-- {
		// 제일 뒤로 swap
		temp := a[i]
		a[i] = a[0]
		a[0] = temp
		heapify(a, 0, i)
	}
	fmt.Println("Heap sort: ", a)
}

func heapify(a []int, i int, length int){
	// max child 찾아서 반복해야함.
	leftchild := i*2 + 1
	rightchild := i*2 + 2
	max := i

	if leftchild < length && a[leftchild] > a[max] {
		max = leftchild
	}
	if rightchild < length && a[rightchild] > a[max] {
		max = rightchild
	}

	if max != i {
		// max를 일단 swap
		temp := a[max]
		a[max]= a[i]
		a[i] = temp
		heapify(a, max, length)
	}
}