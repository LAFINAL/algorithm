package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//a := scanner.Text()
	a := scanner.Bytes()
	fmt.Println(a[0])

	// uint8 부호없는 정수.
}