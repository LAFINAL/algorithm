package leetcode

func romanToInt(s string) int {

	m := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var answer int
	for i:=0; i<len(s); i++ {
		if i != len(s)-1 {
			if m[s[i]] >= m[s[i+1]] {
				answer = answer + m[s[i]]
			} else {
				answer = answer - m[s[i]]
			}
		} else {
			answer = answer + m[s[i]]
		}
	}

	return answer
}
