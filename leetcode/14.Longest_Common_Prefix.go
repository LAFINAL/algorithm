package leetcode

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	answer := strs[0]
	for i:=0; i<len(answer); i++ {
		for j:=1; j<len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != answer[i] {
				return answer[0:i]
			}
		}
	}
	return answer
}
