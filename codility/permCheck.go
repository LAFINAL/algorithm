package codility

func PermCheck(A []int) int {
	max := 0
	for i:=0; i<len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}

	if max != len(A) {
		return 0
	}

	tempMap := make(map[int]int)
	for i:=0; i<len(A); i++ {
		tempMap[A[i]] = tempMap[A[i]] + 1
	}

	for i:=1; i<max+1; i++ {
		_, exist := tempMap[i]
		if !exist {
			return 0
		}
	}

	return 1
}
