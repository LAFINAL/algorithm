package codility

// more than은 초과라는 것을 명심하자.
// counting 하는 것은 map을 쓰면 굉장히 쉽다. 효율성도 만족한다.

func Dominator(A []int) int {
	half := len(A)/2
	tempMap := make(map[int]int)
	for i:=0; i<len(A); i++ {
		tempMap[A[i]] = tempMap[A[i]] + 1
	}

	a := 0
	flag := 0
	for key, value := range tempMap {
		if value > half {
			a = key
			flag = 1
			break
		}
	}
	if flag == 0 {
		return -1
	}

	for i:=0; i<len(A); i++ {
		if A[i] == a {
			return i
		}
	}

	return -1
}
