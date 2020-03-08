package main

import (
	"fmt"
	"log"
	"sort"
)

func main(){

	// input
	var node, edge, start int
	_, err := fmt.Scanln(&node, &edge, &start)
	if err != nil{
		log.Fatal(err)
	}

	var adjacentlist [][]int
	adjacentlist = make([][]int, node, node)

	for i:=0; i<edge; i++ {
		var source, target int
		_, err := fmt.Scanln(&source, &target)

		// make adjacent list
		// caution => source index , real value
		adjacentlist[source-1] = append(adjacentlist[source-1], target)
		adjacentlist[target-1] = append(adjacentlist[target-1], source)
		if err != nil{
			log.Fatal(err)
		}
	}
	for _, adjacentslice := range adjacentlist{
		sort.Slice(adjacentslice, func(i, j int) bool {
			return adjacentslice[i] < adjacentslice[j]
		})
	}
	var visited []bool
	visited = make([]bool, node, node)

	// use queue
	var queue []int
	queue = make([]int, 0, node)
	queue = append(queue, start)
	visited[start-1] = true

	var answer []int
	answer = make([]int, 0, node)
	for len(queue) > 0 {
		vNode := queue[0]
		queue = queue[1:]
		answer = append(answer, vNode)
		for _, adjacentNode := range adjacentlist[vNode-1]{
			if visited[adjacentNode-1] != true {
				queue = append(queue, adjacentNode)
			}
			visited[adjacentNode-1] = true
		}
	}
	fmt.Println(answer)
	fmt.Println("end")

	// bfs

	//var queue []int
	//queue = make([]int, 0, 0)
	//
	//queue = append(queue, start)
	//visited[start-1] = true
	//
	//for len(queue) > 0 {
	//	visit := queue[0]
	//	fmt.Println("visit: ", visit)
	//	queue = queue[1:]
	//
	//	for _, v := range adjacentList[visit-1] {
	//		if visited[v-1] == false {
	//			visited[v-1] = true
	//			queue = append(queue, v)
	//		}
	//	}
	//}
}