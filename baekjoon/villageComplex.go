package baekjoon

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main(){
	var row int
	_, err := fmt.Scan(&row)
	if err != nil {
		log.Fatal(err)
	}

	var inputArray [][]int
	inputArray = make([][]int, row, row)
	for i:=0; i<row; i++{
		var asdf string
		_, err = fmt.Scan(&asdf)
		if err != nil {
			log.Fatal(err)
		}

		for _, ch := range asdf {
			cch, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(cch)
			}
			inputArray[i] = append(inputArray[i], cch)
		}
	}

	var complexCount int
	var vcs []int
	var vc *int
	vc = new(int)
	for i:=0; i<len(inputArray); i++{
		for j:=0; j<len(inputArray); j++ {
			if inputArray[i][j] == 1 {
				complexCount = complexCount + 1
				dfs2(i, j, inputArray, vc)
				vcs = append(vcs, *vc)
				*vc = 0
			}
		}
	}

	sort.Slice(vcs, func(i int, j int) bool{
		return vcs[i] < vcs[j]
	})

	fmt.Println(complexCount)
	for i:=0; i<len(vcs); i++{
		fmt.Println(vcs[i])
	}
}

func dfs2(startX int, startY int, inputArray [][]int, vc *int){
	// 상, 하, 좌, 우
	var dx = [4]int{0, 0, -1, 1}
	var dy = [4]int{-1, 1, 0, 0}

	inputArray[startX][startY] = 0
	*vc = *vc+1

	for i:=0; i<4; i++{
		if startX+dx[i] < 0 || startY+dy[i] <0 || startX+dx[i] >= len(inputArray) || startY+dy[i] >= len(inputArray){
			continue
		}

		if inputArray[startX+dx[i]][startY+dy[i]] != 0 {
			dfs2(startX+dx[i], startY+dy[i], inputArray, vc)
		}
	}
}