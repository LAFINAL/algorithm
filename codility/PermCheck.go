package codility

func PermCheck(A []int) int {
	max := 0
	for i:=0; i<len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}

	tempSlice := make([]bool, max+1, max+1)
	tempSlice[0] = true

	for i:=0; i<len(A); i++ {
		tempSlice[A[i]] = true
	}

	for i:=0; i<len(tempSlice); i++ {
		if tempSlice[i] == false {
			return 0
		}
	}

	// 한 번 넣는거 확인은 Map으로 하면 된다.
	// map에 add하고 다시 확인해서 있으면 false 리턴.
	return 1
}
