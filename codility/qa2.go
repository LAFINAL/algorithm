package codility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TossReal2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	lottoNum := strings.Split(input, " ")

	nums := make([]int, len(lottoNum), len(lottoNum))

	for i:=0; i<len(lottoNum); i++ {
		b, _ := strconv.Atoi(lottoNum[i])
		if b < 1 || b > 45 {
			fmt.Print("false")
			return
		}
		nums[i] = b
	}

	checkDup := make(map[int]bool)
	for i:=0; i<len(nums); i++ {
		checkDup[nums[i]] = true
	}

	if len(checkDup) != 6 {
		fmt.Print("false")
		return
	}

	for i:=0; i<len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			fmt.Print("false")
			return
		}
	}

	fmt.Print("true")
}
