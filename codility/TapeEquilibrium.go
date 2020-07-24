package codility

func TapeEquilibrium(A []int) int {
	leftSideSum := 0
	rightSideSum := 0
	for i:=0; i<len(A); i++ {
		rightSideSum = rightSideSum + A[i]
	}
	// leftSideSum에는 0이, rightSideSum에는 전체 길이가 들어가 있음.

	distance := 0
	min := 100000000

	for i:=0; i<len(A); i++ {
		leftSideSum = leftSideSum + A[i]
		rightSideSum = rightSideSum - A[i]
		distance = rightSideSum - leftSideSum
		if distance < 0 {
			distance = distance * -1
		}
		if distance < min {
			min = distance
		}

	}
	return min
}