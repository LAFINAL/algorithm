package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	info := strings.Split(text, " ")
	card_num, _ := strconv.Atoi(info[0])
	target, _ := strconv.Atoi(info[1])

	scanner.Scan()
	text = scanner.Text()
	text2 := strings.Split(text, " ")
	var card []int
	card = make([]int, len(text2), len(text2))
	for i:=0; i<len(card); i++{
		card[i], _ = strconv.Atoi(text2[i])
	}

	var score int
	var answer int
	for i:=0; i<card_num; i++{
		for j:=i+1; j<card_num; j++{
			for k:=j+1; k<card_num; k++{
				score = card[i] + card[j] + card[k]
				if score <= target && answer < score {
					answer = score
				}
			}
		}
	}
	fmt.Println(answer)
}
