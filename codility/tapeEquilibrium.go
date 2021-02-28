package codility

// point
// 1. one pointer 문제. two pointer는 다소 어렵다.
// 2. max와 min의 초기값에 주의하자. max는 나올 수 있는 값 중 제일 작은 값으로, min은 제일 큰 값으로!
// 단순히 0으로 설정하면 안됨. 항상 음수를 생각하자.

func TapeEquilibrium(A []int) int {
	leftSideSum := 0
	rightSideSum := 0
	for i:=0; i<len(A); i++ {
		rightSideSum = rightSideSum + A[i]
	}
	// one pointer로 푸는 것. two pointer는 어렵다.
	canAnswer := make([]int, len(A)-1, len(A)-1)
	for i:=0; i<len(A)-1; i++ {
		leftSideSum = leftSideSum + A[i]
		rightSideSum = rightSideSum - A[i]
		distance := leftSideSum - rightSideSum
		if distance < 0 {
			distance = distance * -1
		}
		canAnswer[i] = distance
	}

	// max는 처음 최소값으로 잡고 min은 최대값으로 잡아야 실수가없다.
	min := 2000
	for i:=0; i<len(canAnswer); i++ {
		if canAnswer[i] < min {
			min = canAnswer[i]
		}
	}
	return min
}