package main

import "fmt"

func main(){
	//a := "t"
	//fmt.Println(a)
	//fmt.Println(reflect.TypeOf(a))
	//b := 't'
	//fmt.Println(b)
	//fmt.Println(reflect.TypeOf(b))
	//var c byte
	//c = 't'
	//fmt.Println(c)
	//fmt.Println(reflect.TypeOf(c))
	//var d rune
	//d = 't'
	//fmt.Println(d)
	//fmt.Println(reflect.TypeOf(d))
	//var e int
	//d = 't'
	//fmt.Println(reflect.TypeOf(e))
	//a := "t"
	//for i, j := range a {
	//	fmt.Println(i, j)
	//	fmt.Println(reflect.TypeOf(j))
	//}

	// string은 불변이라 변경할 수 없다.

	a := [][]int{
		{1,2},
		{3,4},
		{4,5},
	}
	b := []int{10,11}

	fmt.Println(a)
	a = append(a, make([]int, 2, 2))
	fmt.Println(a)
	copy(a[2:], a[1:])
	fmt.Println(a)
	//copy(a[1],b)
	a[1] = b
	fmt.Println(a)

}