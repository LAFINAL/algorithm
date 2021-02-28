package leetcode

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var index int
	switch ruleKey {
	case "type":
		index = 0
	case "color":
		index = 1
	case "name":
		index = 2
	}

	count := 0
	for _, value := range items {
		if value[index] == ruleValue {
			count = count + 1
		}
	}

	return count
}