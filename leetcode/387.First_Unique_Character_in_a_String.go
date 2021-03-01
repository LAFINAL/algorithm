package leetcode

func firstUniqChar(s string) int {
	frequency := make([]int, 26, 26)
	for i:=0; i<len(s); i++ {
		frequency[s[i]-'a'] = frequency[s[i]-'a'] + 1
	}
	answer := -1
	for i:=0; i<len(s); i++ {
		if frequency[s[i]-'a'] == 1 {
			answer = i
			break
		}
	}
	return answer
}
