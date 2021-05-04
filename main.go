package main

import "fmt"

// annotation for test
func main(){
	a := 1.0
	n := 3.0

	for i:=0; i<5; i++ {
		b := a/n
		fmt.Println(b)
		b = b*10

	}
}
