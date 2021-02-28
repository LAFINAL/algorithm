package codility

func Distinct(A []int) int {
	distinctMap := make(map[int]bool)
	for i:=0; i<len(A); i++ {
		distinctMap[A[i]] = true
	}

	return len(distinctMap)
}
