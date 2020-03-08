package main

import "fmt"

func main(){

	for i:=0; i<5; i++{
		a := 5
		fmt.Println(&a)
	}

	// slice 특징.
	//a := make([]int, 3)
	//a[0] = 5
	//a[1] = 6
	//a[2] = 7
	//
	//b := make([]int, 2)
	//b[0] = 10
	//b[1] = 8
	//fmt.Println(&a[0])
	//a = b
	//fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))
	//
	//fmt.Println(&a[0])
}
