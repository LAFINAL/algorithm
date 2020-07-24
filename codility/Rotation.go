package codility


func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	doubleA := make([]int, len(A)*2, len(A)*2)

	for i:=0; i<len(doubleA); i++ {
		doubleA[i] = A[i%len(A)]
	}

	K = K%len(A)
	if K==0 {
		return A
	} else {
		return doubleA[len(A)-K:2*len(A)-K]
	}

	return nil
}
