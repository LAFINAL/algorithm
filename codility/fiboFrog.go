package codility

func FiboFrog(A []int) int {

	// fibo map을 만들자.
	fiboSlice := make([]int, 0, 0)
	fiboSlice = append(fiboSlice, 0)
	fiboSlice = append(fiboSlice, 1)
	n := 2
	for {
		sum := fiboSlice[n-1] + fiboSlice[n-2]
		if sum > 100000 {
			break
		}
		fiboSlice = append(fiboSlice, sum)
		n = n + 1
	}
	fiboMap := make(map[int]int)
	for i:=0; i<len(fiboSlice); i++ {
		fiboMap[fiboSlice[i]] = i
	}


	leaf := make([]int, 0, 0)
	for i:=0; i<len(A); i++ {
		if A[i] == 1 {
			leaf = append(leaf, i+1)
		}
	}
	leaf = append(leaf, len(A)+1)
	//visited := make([]int, len(leaf), len(leaf))
	//fiboFrogBFS(leaf, visited)
	return 0
}

func fiboFrogBFS(leaf []int, visited []int, fiboMap map[int]int) int {
	// bfs가 끝나는 조건은
	if visited[leaf[len(leaf)-1]] != 0 {
		// terminate recursive
	}

	// fibo에 맞는 애들 선출
	candidate := make([]int, 0, 0)
	for i:=0; i<len(leaf); i++ {
		_, exist := fiboMap[leaf[i]]
		if exist {
			candidate = append(candidate, i)
		}
	}
	//// candidate에는 대상자의 leaf의 index가 들어있음.
	//for i:=0; i<len(candidate); i++ {
	//	visited[candidate[i]] =
	//}


	return 0
}
