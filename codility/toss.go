package codility

import (
	"fmt"
	"strconv"
)

func Toss1() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	value, _ := strconv.Atoi(input)
	if value % 2 == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}
