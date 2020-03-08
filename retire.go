package main

import "fmt"

func main(){

	// prepare
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
	}

	timeTable := make([]int, n, n)
	payTable  := make([]int, n, n)
	totalPayTable := make([]int, n, n)

	for i:=0; i<n; i++ {
		_, err = fmt.Scanln(&timeTable[i], &payTable[i])
		if err != nil {
			fmt.Println(err)
		}

		totalPayTable[i] = payTable[i]
	}

	for i:=1; i<=n-1; i++ {
		for j:=0; j<i; j++ {
			if i-j >= timeTable[j] {
				totalPayTable[i] = Max(payTable[i] + totalPayTable[j], totalPayTable[i])
			}
		}
	}

	var answer int
	for i:=0; i<n-1; i++ {
		if i + timeTable[i] <= n {
			if answer < totalPayTable[i] {
				answer = totalPayTable[i]
			}
		}
	}
	fmt.Println(answer)
}

func Max(a, b int) (v int){
	if a > b {
		v = a
		return
	}
	v = b
	return
}