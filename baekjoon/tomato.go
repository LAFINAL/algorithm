package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//var reader = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)

func main() {
	var column, row int
	fmt.Scanln(&column, &row)

	var WareHouse [][]int
	WareHouse = make([][]int, row, row)

	var Queue []int
	for i:=0; i<row; i++{
		sc.Scan()
		//text, _ := reader.ReadString('\n')
		text := sc.Text()
		//text = strings.TrimSuffix(text, "\n")
		temp_box := strings.Split(text, " ")
		WareHouse[i] = make([]int, column, column)
		for j:=0; j<column; j++{
			box, _ := strconv.Atoi(temp_box[j])
			if box == 1{
				Queue = append(Queue, i, j)
			}
			WareHouse[i][j] = box
		}
	}

	var idx = []int{0, 0, -1, 1}
	var idy = []int{-1, 1, 0, 0}
	var nextX int
	var nextY int
	for len(Queue) > 0 {
		dx := Queue[0]
		dy := Queue[1]
		Queue = Queue[2:]
		currentV := WareHouse[dx][dy]

		for i:=0; i<4; i++{
			nextX = dx+idx[i]
			nextY = dy+idy[i]
			if nextX < 0 || nextY < 0 || nextX >= row || nextY >= column{
				continue
			}
			if WareHouse[nextX][nextY] != 0 {
				continue
			}
			WareHouse[nextX][nextY] = currentV + 1
			Queue = append(Queue, nextX, nextY)
		}
	}

	var answer int
out:
	for i:=0; i<row; i++{
		for j:=0; j<column; j++{
			if WareHouse[i][j] == 0 {
				answer = 0
				break out
			}
			if WareHouse[i][j] > answer {
				answer = WareHouse[i][j]
			}
		}
	}
	fmt.Println(answer-1)
}

// 63060	160	Go	1