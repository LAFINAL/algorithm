package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Bytes()

	var alphabet []int
	alphabet = make([]int, 26, 26)

	for i:=0; i<len(alphabet); i++ {
		alphabet[i] = -1
	}

	stand := byte('a')
	for i, t := range(text) {
		if alphabet[t-stand] == -1 {
			alphabet[t-stand] = i
		}
	}

	for i:=0; i<len(alphabet); i++ {
		if i == len(alphabet){
			fmt.Println(alphabet[i])
		} else {
			fmt.Print(alphabet[i], " ")
		}
	}
}