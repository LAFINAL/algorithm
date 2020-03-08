package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Scan()
	text := scanner.Bytes()

	// capitalize and normalize
	gap := byte('a') - byte('A')
	normalize := byte('A')
	for i, ch := range text {
		if ch > byte('Z') {
			text[i] = ch - gap
		}
		text[i] = text[i] - normalize
	}

	alphabet := [26]int{}
	for _, ch := range text {
		alphabet[ch] = alphabet[ch] + 1
	}

	var max int
	var index int
	for i, al := range alphabet {
		if max < al {
			max = al
			index = i
		}
	}

	var count int
	for i:=0; i<len(alphabet); i++ {
		if alphabet[i] == max {
			count = count + 1
		}
	}

	if count > 1 {
		fmt.Println("?")
		return
	}

	var answer byte

	answer = byte(index) + normalize
	fmt.Print(string(answer))
}