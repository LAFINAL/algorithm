package sorting

// selection sort는 sorting 중에 제일 안좋은 거.
// 시간복잡도는 전부 n2이다.
// min 값 찾아서 switch

func Selection() {
	a := []int{5,3,13,1,2,3,64,3,1,8,6}

	//for i:=0; i<len(a)-1; i++ {
	//	min := i
	//	for j:=i+1; j<len(a); j++ {
	//		if a[min] > a[j] {
	//			min = j
	//		}
	//	}
	//
	//	if min != i { // min이 i랑 같지 않으면 swap
	//		temp := a[i]
	//		a[i] = a[min]
	//		a[min] = temp
	//	}
	//}
	//fmt.Println(a)

	// selection sort는 bubble sort에 비해서 swap이 적다. 그래서 시간 복잡도는 똑같은데 더 효율적. min 값으로 앞에서부터.
	//for i:=0; i<len(a)-1; i++ {
	//	min := i
	//	for j:=i+1; j<len(a); j++ {
	//		if a[j] < a[min] {
	//			min = j
	//		}
	//	}
	//	if min != i {
	//		temp := a[min]
	//		a[min] = a[i]
	//		a[i] = temp
	//	}
	//}
	//fmt.Println("Selection sort: ", a)

	// selection sort는 min 값을 찾는 것. min 값을 찾아서 맨 앞에 쓰는 것.
	for i:=0; i<len(a)-1; i++ {
		min := i
		for j:=i+1; j<len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			temp := a[min]
			a[min] = a[i]
			a[i] = temp
		}
	}
}