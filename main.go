package main

import "fmt"

func main(){
	a := 1.0
	n := 3.0
	// change annotation location

	// feature 3
	for i:=0; i<5; i++ {
		b := a/n
		fmt.Println(b)
		b = b*10

	}
}
// annotation for pr test