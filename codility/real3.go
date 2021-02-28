package codility

import (
	//"fmt"
	"strings"
)

func Real3(S string) string {
	if len(S) == 1 {
		return S
	}

	replacer := strings.NewReplacer("CC", "", "BB", "", "AA", "")
	tempAnswer := replacer.Replace(S)
	if tempAnswer == S {
		return tempAnswer
	} else {
		return Real3(tempAnswer)
	}

	return ""
}
