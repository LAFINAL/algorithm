package main

import "fmt"

func main(){

	answer := convert("AB", 1)

	if answer == "PAHNAPLSIIGYIR" {
		fmt.Println(answer)
	} else {
		fmt.Println(answer)
		fmt.Println("PAHNAPLSIIGYIR")
	}
}

func convert(s string, numRows int) string {
	a := make([]byte, 0, 0)
	if numRows == 1{
		return s
	}

	for i:=0; i<numRows; i++ {
		for j:=i; j<len(s); j=j+2*(numRows-1) {
			a = append(a, s[j])
			if j+(numRows-i-1)*2 < len(s) && i!=0 && i!=numRows-1 { // 여기 빼기 2, 4 이렇게
				//a = append(a, s[j-i*2])
				a = append(a, s[j+(numRows-i-1)*2])
			}
		}
	}

	return string(a)
}
