package baekjoon

import (
	"fmt"
	"log"
	"sort"
)

func main(){
	// bfs 복습 겸 input, adjacent list 만들기
	var node, edge, start int
	_, err := fmt.Scanln(&node, &edge, &start)
	if err != nil {
		log.Fatal(err)
	}

	var adjacentList [][]int
	adjacentList = make([][]int, node, node)

	for i:=0; i<edge; i++ {
		var source, target int
		_, err = fmt.Scanln(&source, &target)
		if err != nil {
			log.Fatal(err)
		}

		adjacentList[source-1] = append(adjacentList[source-1], target)
		adjacentList[target-1] = append(adjacentList[target-1], source)

		for _, adjacentSlice := range adjacentList {
			sort.Slice(adjacentSlice, func (i, j int) bool {
				return adjacentSlice[i] < adjacentSlice[j]
			})
		}
	}
	fmt.Println(adjacentList)

	var visited []bool
	visited = make([]bool, node, node)

	visited[start-1] = true
	// push
	dfs(adjacentList, visited, start)

}

func dfs(adjacentList [][]int, visited []bool, start int) {
	fmt.Println(start)
	visited[start-1] = true

	for i:=0; i<len(adjacentList[start-1]); i++{
		if visited[adjacentList[start-1][i]-1] == false {
			dfs(adjacentList, visited, adjacentList[start-1][i])
		}
	}
	return
}