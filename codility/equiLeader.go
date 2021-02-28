package codility

func EquiLeader(A []int) int {
	tempMap := make(map[int]int)
	leader := make([]int, len(A), len(A))



	max := 0
	for i:=0; i<len(A); i++ {
		tempMap[A[i]] = tempMap[A[i]] + 1
		if tempMap[A[i]] > max {
			max = tempMap[A[i]]
		}
		if max > len(A)/2 {
			leader[i] = A[i]
		}
	}
	// leader slice가 완성되면
	for i:=0; i<len(leader); i++ {

	}
	return 0
}
