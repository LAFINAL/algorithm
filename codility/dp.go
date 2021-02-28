package codility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dp [][]int

func Dp() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	totalCase, _ := strconv.Atoi(scanner.Text())
	for i:=0; i<totalCase; i++ {
		//scanner.Scan()
		//pageNum, _ := strconv.Atoi(scanner.Text())
		//scanner.Scan()
		//text := scanner.Text()
		//splitText := strings.Split(text, " ")
		//for j:=0; j<pageNum; j++ {
		//	splitText[j] ==
		//}
	}
	a := []int{100,150,200,210}
	// initialize slice
	dp = make([][]int, len(a), len(a))
	for i:=0; i<len(dp); i++{
		dp[i] = make([]int, len(a), len(a))
	}
	fmt.Println(dp)
	for i:=0; i<len(a); i++{
		dp[i][i] = a[i]
	}
	fmt.Println(dp)

	for y:=1; y<len(a); y++ {
		for x:=y-1; x>=0; x-- {
			dp[x][y] = min(dp[x][y-1], dp[x+1][y]) + a[y]
			fmt.Println("dp: ", dp)
		}
	}

	answer := 0
	for i:=1; i<len(dp[0]); i++ {
		answer = answer + dp[0][i]
	}
	fmt.Println("answer: ", answer)
}

func min(a int, b int) int {
	if a > b{
		return b
	} else {
		return a
	}
}
