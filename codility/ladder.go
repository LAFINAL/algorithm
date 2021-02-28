package codility

func Ladder(A []int, B []int) []int {
	answer := make([]int, len(A), len(A))

	normalMax := 0
	for i:=0; i<len(B); i++ {
		if B[i] > normalMax {
			normalMax = B[i]
		}
	}
	normalSlice := make([]int, normalMax+1, normalMax+1)
	normalSlice[0] = 1
	for i:=1; i<len(normalSlice); i++ {
		normalSlice[i] = normalSlice[i-1] * 2
	}

	fiboMax := 0
	for i:=0; i<len(A); i++{
		if A[i] > fiboMax {
			fiboMax = A[i]
		}
	}

	fiboSlice := make([]int, fiboMax, fiboMax)
	if fiboMax == 1{
		fiboSlice[0] = 1
	} else {
		fiboSlice[0] = 1
		fiboSlice[1] = 2
		for i:=2; i<len(fiboSlice); i++ {
			fiboSlice[i] = (fiboSlice[i-1] + fiboSlice[i-2])%normalSlice[len(normalSlice)-1]
		}
	}

	for i:=0; i<len(A); i++ {
		answer[i] = fiboSlice[A[i]-1] % normalSlice[B[i]]
	}

	return answer
}

//func fiboNum(a int) int {
//	memo := make([]int, a, a)
//	if a == 0 {
//		return 0
//	}
//	if a == 1 {
//		return 1
//	}
//	if a == 2 {
//		return 2
//	}
//
//	memo[0] = 1
//	memo[1] = 2
//	for i:=2; i<a; i++ {
//		memo[i] = memo[i-1] + memo[i-2]
//	}
//	return memo[a-1]
//}

//func normalNum(b int) int {
//	output := 1
//	for i:=0; i<b; i++{
//		output = output * 2
//	}
//	return output
//}