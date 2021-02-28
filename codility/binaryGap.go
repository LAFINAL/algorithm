package codility

func BianryGap(N int) int {
	binary := make([]int, 0, 0)
	for N != 0 {
		remainder := N % 2
		binary = append(binary, remainder)
		N = N / 2
	}

	loc := make([]int, 0, 0)
	for i:=0; i<len(binary); i++ {
		if binary[i] == 1 {
			loc = append(loc, i)
		}
	}

	if len(loc) == 1{
		return 0
	}

	max := 0
	for i:=0; i<len(loc)-1; i++ {
		if loc[i+1] - loc[i] > max {
			max = loc[i+1] - loc[i]
		}
	}
	return max-1
}
