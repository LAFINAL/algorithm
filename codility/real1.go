package codility

func Real1(A [][]int) int {
	rowSum := make([]int, len(A), len(A))
	colSum := make([]int, len(A[0]), len(A[0]))
	for i:=0; i<len(A); i++{
		sum := 0
		for j:=0; j<len(A[0]); j++{
			sum = sum + A[i][j]
		}
		rowSum[i] = sum
	}

	for i:=0; i<len(A[0]); i++ {
		sum := 0
		for j:=0; j<len(A); j++ {
			sum = sum + A[j][i]
		}
		colSum[i] = sum
	}
	// rowSum이랑 colSum은 잘 나온다.

	rowSumRight := 0
	rowSumLeft := 0
	for i:=0; i<len(rowSum); i++ {
		rowSumRight = rowSumRight + rowSum[i]
	}

	rowCan := make([]int, 0, 0)
	for i:=0; i<len(rowSum); i++ {
		if i>0{
			rowSumLeft = rowSumLeft + rowSum[i-1]
		}
		rowSumRight = rowSumRight - rowSum[i]
		if rowSumLeft == rowSumRight {
			rowCan = append(rowCan, i)
		}
	}

	// col
	colSumRight := 0
	colSumLeft := 0
	for i:=0; i<len(colSum); i++ {
		colSumRight = colSumRight + colSum[i]
	}

	colCan := make([]int, 0, 0)
	for i:=0; i<len(colSum); i++ {
		if i > 0 {
			colSumLeft = colSumLeft + colSum[i-1]
		}
		colSumRight = colSumRight - colSum[i]
		if colSumLeft == colSumRight {
			colCan = append(colCan, i)
		}
	}
	return len(rowCan) * len(colCan)
}
