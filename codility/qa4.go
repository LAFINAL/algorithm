package codility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func TossReal4() {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	banks := strings.Split(input, " ")
	answer := make([]string, 0, 0)
	for i:=0; i<len(banks); i++ {
		bank := banks[i]
		flag := -1
		for j:=0; j<len(answer); j++ {
			if bank == answer[j] {
				flag = j
				break
			}
		}
		if flag == -1 {
			if len(answer) == 5 {
				answer = answer[1:]
			}
			answer = append(answer, bank)
			for k:=len(answer)-1; k>=0; k-- {
				if k != 0 {
					fmt.Print(answer[k] + " ")
				} else {
					fmt.Print(answer[k])
				}
			}
			fmt.Println()
		} else {
			answer = append(answer[:flag], answer[flag+1:]...)
			answer = append(answer, bank)
			for k:=len(answer)-1; k>=0; k-- {
				if k != 0 {
					fmt.Print(answer[k] + " ")
				} else {
					fmt.Print(answer[k])
				}
			}
			fmt.Println()
		}
	}
}
