package main

import "fmt"

type asdff struct {
	word string
	depth int
}
func main(){
	c := make([]string, 0, 0)
	c = append(c, "hot", "dot", "dog", "lot", "log", "cog")
	answer := solution("hit", "cog", c)
	fmt.Println(answer)
}

func solution(begin string, target string, words []string) int {
	answer := 0
	queue := make([]*asdff, 0, 0) // make and init slice
	a := new(asdff)
	a.word = begin
	a.depth = 0
	queue = append(queue, a)// queue.push
	flag := 0
	for len(queue) != 0 { // go does not have while loop
		fmt.Println(queue)
		fv := queue[0]// queue.pop
		queue = queue[1:] // slice slice

		for index, word := range words { // words 동안 돌아야 함. 조건 만족하는거
			if word == "" {
				continue
			}

			if condition(fv.word, word) {
				if word == target { // target이랑 같으면 answer +1하고 종료
					fmt.Println(word)
					fmt.Println(target)
					fmt.Println("test")
					answer = a.depth + 1
					flag = 1
					break
				}

				b := new(asdff)
				b.word = word
				b.depth = a.depth + 1
				queue = append(queue, b) // condition 만족하면 queue에 넣는다.
				words[index] = "" // 그리고 공백으로 치환
			}

		}
		if flag == 1{
			break
		}
	}
	return answer
}

func condition(a string, b string) bool {
	count := 0
	for i:=0; i<len(a); i++{
		if a[i] == b[i] {
			count = count + 1
		}
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}