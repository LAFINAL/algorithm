package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var idx = []int{0, 0, -1, 1}
var idy = []int{-1, 1, 0 ,0}
var farm [][]int
var worm int
var nextX int
var nextY int

func main() {
	// input

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cases, _ := strconv.Atoi(text)

	answers := make([]int, cases, cases)
	for i:=0; i<cases; i++{
		var column int
		var row int
		var cabbage_num int
		scanner.Scan()
		text = scanner.Text()
		textS := strings.Split(text, " ")
		column, _ = strconv.Atoi(textS[0])
		row, _ = strconv.Atoi(textS[1])
		cabbage_num, _ = strconv.Atoi(textS[2])

		farm = make([][]int, row, row)
		for k:=0; k<row; k++{
			farm[k] = make([]int, column, column)
		}

		for j:=0; j<cabbage_num; j++{
			var dy int
			var dx int
			scanner.Scan()
			text = scanner.Text()
			textS = strings.Split(text, " ")
			dy, _ = strconv.Atoi(textS[0])
			dx, _ = strconv.Atoi(textS[1])
			farm[dx][dy] = 1
		}

		worm = 0
		for i:=0; i<row; i++{
			for j:=0; j<column; j++{
				if farm[i][j] == 1 {
					worm = worm + 1
					dfs3(i, j, farm)
				}
			}
		}
		answers[i] = worm
	}
	for _, answer := range answers{
		fmt.Println(answer)
	}
}

func dfs3(startX int, startY int, farm [][]int) {
	farm[startX][startY] = 0

	for i:=0; i<4; i++{
		nextX = startX + idx[i]
		nextY = startY + idy[i]
		if nextX < 0 || nextY < 0 || nextX >= len(farm) || nextY >= len(farm[0]) {
			continue
		}

		if farm[nextX][nextY] == 1 {
			dfs3(nextX, nextY, farm)
		}
	}
}
