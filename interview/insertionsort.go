package sorting

import "fmt"

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Insertion() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}
	//
    //// 앞에서부터 정렬임. 이미 정열되어있다면 시간복잡도 n
    //for i:=1; i<len(a); i++ {
    //	// 일단 처음 값을 저장해놔야함.
    //	keep := a[i]
    //	flag := -1
    //	for j:=i-1; j>=0; j-- {
	//		if keep < a[j] {
	//			a[j+1] = a[j]
	//			flag = j
	//		}
	//	}
	//	if flag != -1 {
	//		a[flag] = keep
	//	}
	//}
	//fmt.Println(a)

	// insertion sort는 bubble sort, selection sort 보다는 좋음. 비교 및 swap 회수가 적음.
	// 기존에 정렬이 이미 되어 있는 상태에서는 시간복잡도가 n이된다.

	for i:=1; i<len(a); i++ {
		key := a[i] // key를 미리 설정해 둔다.
		flag := -1 // index를 알 수 있는 flag 값도 있어야 한다.
		for j:=i-1; j>=0; j-- {
			if a[j] > key { // 앞에 수가 더 크면 뒤로 복사
				a[j+1] = a[j]
				flag = j
			}
		}

		if flag != -1 {
			a[flag] = key
		}
	}

	fmt.Println("Insertion sort: ", a)
}