package codility

var result = 0
func Real2(A []int) int {
	addFirstTwoAndRemoveFirst(A)
	return result
}

func addFirstTwoAndRemoveFirst(A []int) {
	if len(A) == 0 || len(A) == 1{
		return
	}

	for i:=0; i<len(A); i++ {
		for j:=i+1; j<len(A); j++ {
			if A[i] > A[j] {
				temp := A[i]
				A[i] = A[j]
				A[j] = temp
			}
		}
	}

	temp := A[0] + A[1]
	result = result + temp
	A[1] = temp
	A = A[1:]
	addFirstTwoAndRemoveFirst(A)
}